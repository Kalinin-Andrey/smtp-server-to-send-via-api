package app

import (
	"context"
	golog "log"

	"smtp2api/internal/domain/advertiser"
	"smtp2api/internal/domain/advertising_campaign"
	"smtp2api/internal/domain/offer"

	"smtp2api/internal/infrastructure/repository/yaruzplatform"

	"smtp2api/pkg/yarus_platform"

	"github.com/pkg/errors"

	minipkg_gorm "github.com/minipkg/db/gorm"
	"github.com/minipkg/db/redis"
	"github.com/minipkg/db/redis/cache"
	"github.com/minipkg/log"
	"smtp2api/internal/pkg/apperror"
	"smtp2api/internal/pkg/config"

	"smtp2api/internal/domain/email"
)

// App struct is the common part of all applications
type App struct {
	Cfg   config.Configuration
	Infra *Infrastructure
}

type Infrastructure struct {
	Logger      log.ILogger
	APIProvider APIProvider
}

type APIProvider interface {
	Send() error
}

// New func is a constructor for the App
func New(ctx context.Context, cfg config.Configuration) *App {
	logger, err := log.New(cfg.Log)
	if err != nil {
		golog.Fatal(err)
	}

	infra, err := NewInfra(ctx, logger, cfg)
	if err != nil {
		golog.Fatal(err)
	}

	app := &App{
		Cfg:   cfg,
		Infra: infra,
	}

	err = app.Init()
	if err != nil {
		golog.Fatal(err)
	}

	return app
}

func NewInfra(ctx context.Context, logger log.ILogger, cfg config.Configuration) (*Infrastructure, error) {
	IdentityDB, err := minipkg_gorm.New(logger, cfg.DB.Identity)
	if err != nil {
		return nil, err
	}

	rDB, err := redis.New(cfg.DB.Redis)
	if err != nil {
		return nil, err
	}

	yaruzRepository, err := yarus_platform.NewPlatform(ctx, cfg.YaruzConfig())
	if err != nil {
		return nil, err
	}

	return &Infrastructure{
		Logger:          logger,
		IdentityDB:      IdentityDB,
		Redis:           rDB,
		YaruzRepository: yaruzRepository,
	}, nil
}

func (app *App) Init() (err error) {
	if err := app.SetupRepositories(); err != nil {
		return err
	}
	app.SetupServices()
	return nil
}

func (app *App) SetupRepositories() (err error) {
	var ok bool

	userRepo, err := yaruzplatform.GetRepository(app.Infra.Logger, app.Infra.YaruzRepository, email.EntityType)
	if err != nil {
		golog.Fatalf("Can not get yaruz repository for entity %q, error happened: %v", email.EntityType, err)
	}

	app.Domain.userRepository, ok = userRepo.(email.Repository)
	if !ok {
		return errors.Errorf("Can not cast yaruz repository for entity %q to %vRepository. Repo: %v", email.EntityType, email.EntityType, userRepo)
	}

	advertiserRepo, err := yaruzplatform.GetRepository(app.Infra.Logger, app.Infra.YaruzRepository, advertiser.EntityType)
	if err != nil {
		golog.Fatalf("Can not get yaruz repository for entity %q, error happened: %v", advertiser.EntityType, err)
	}

	app.Domain.advertiserRepository, ok = advertiserRepo.(advertiser.Repository)
	if !ok {
		return errors.Errorf("Can not cast yaruz repository for entity %q to %vRepository. Repo: %v", advertiser.EntityType, advertiser.EntityType, advertiserRepo)
	}

	advertisingCampaignRepo, err := yaruzplatform.GetRepository(app.Infra.Logger, app.Infra.YaruzRepository, advertising_campaign.EntityType)
	if err != nil {
		golog.Fatalf("Can not get yaruz repository for entity %q, error happened: %v", advertising_campaign.EntityType, err)
	}

	app.Domain.advertisingCampaignRepository, ok = advertisingCampaignRepo.(advertising_campaign.Repository)
	if !ok {
		return errors.Errorf("Can not cast yaruz repository for entity %q to %vRepository. Repo: %v", advertising_campaign.EntityType, advertising_campaign.EntityType, advertisingCampaignRepo)
	}

	offerRepo, err := yaruzplatform.GetRepository(app.Infra.Logger, app.Infra.YaruzRepository, offer.EntityType)
	if err != nil {
		golog.Fatalf("Can not get yaruz repository for entity %q, error happened: %v", offer.EntityType, err)
	}

	app.Domain.offerRepository, ok = offerRepo.(offer.Repository)
	if !ok {
		return errors.Errorf("Can not cast yaruz repository for entity %q to %vRepository. Repo: %v", offer.EntityType, offer.EntityType, offerRepo)
	}

	//if app.Auth.SessionRepository, err = redisrep.NewSessionRepository(app.Infra.Redis, app.Cfg.SessionLifeTime, app.Domain.User.Repository); err != nil {
	//	return errors.Errorf("Can not get new SessionRepository err: %v", err)
	//}
	//app.Auth.TokenRepository = jwt.NewRepository()

	app.Infra.Cache = cache.NewService(app.Infra.Redis, app.Cfg.CacheLifeTime)

	return nil
}

func (app *App) SetupServices() {
	app.Domain.User = email.NewService(app.Infra.Logger, app.Domain.userRepository)
	app.Domain.Advertiser = advertiser.NewService(app.Infra.Logger, app.Domain.advertiserRepository)
	app.Domain.AdvertisingCampaign = advertising_campaign.NewService(app.Infra.Logger, app.Domain.advertisingCampaignRepository)
	app.Domain.Offer = offer.NewService(app.Infra.Logger, app.Domain.offerRepository)
	//app.Auth.Service = auth.NewService(app.Cfg.JWTSigningKey, app.Cfg.JWTExpiration, app.Domain.User.Service, app.Infra.Logger, app.Auth.SessionRepository, app.Auth.TokenRepository)
}

// Run is func to run the App
func (app *App) Run() error {
	return nil
}

func (app *App) Stop() error {
	errRedis := app.Infra.Redis.Close()
	errDB01 := app.Infra.IdentityDB.Close()
	errDB02 := app.Infra.YaruzRepository.Stop()

	switch {
	case errDB01 != nil:
		return errors.Wrapf(apperror.ErrInternal, "db close error: %v", errDB01)
	case errDB02 != nil:
		return errors.Wrapf(apperror.ErrInternal, "yarus repository close error: %v", errDB02)
	case errRedis != nil:
		return errors.Wrapf(apperror.ErrInternal, "redis close error: %v", errRedis)
	}

	return nil
}
