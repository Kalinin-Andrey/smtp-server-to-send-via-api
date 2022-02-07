package app

import (
	"context"
	golog "log"
	"smtp2api/pkg/email_provider"
	"smtp2api/pkg/email_provider/unisender"

	"github.com/minipkg/log"
	"smtp2api/internal/pkg/config"
)

// App struct is the common part of all applications
type App struct {
	Cfg   config.Configuration
	Infra *Infrastructure
}

type Infrastructure struct {
	Logger        log.ILogger
	EmailProvider email_provider.EmailProvider
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

	return app
}

func NewInfra(ctx context.Context, logger log.ILogger, cfg config.Configuration) (*Infrastructure, error) {
	emailProvider := unisender.New()

	return &Infrastructure{
		Logger:        logger,
		EmailProvider: emailProvider,
	}, nil
}

// Run is func to run the App
func (app *App) Run() error {
	return nil
}

func (app *App) Stop() error {
	return nil
}
