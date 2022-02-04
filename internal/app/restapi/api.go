package restapi

import (
	"log"
	"net/http"
	"time"

	"github.com/minipkg/log/accesslog"

	"github.com/go-ozzo/ozzo-routing/v2/content"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/cors"
	"github.com/go-ozzo/ozzo-routing/v2/file"
	"github.com/go-ozzo/ozzo-routing/v2/slash"

	"smtp2api/internal/pkg/config"

	"github.com/minipkg/ozzo_routing"
	"github.com/minipkg/ozzo_routing/errorshandler"

	commonApp "smtp2api/internal/app"
	"smtp2api/internal/app/restapi/controller"
)

// Version of API
const Version = "1.0.0"

// App is the application for API
type App struct {
	*commonApp.App
	Server *http.Server
}

// New func is a constructor for the ApiApp
func New(commonApp *commonApp.App, cfg config.Configuration) *App {
	app := &App{
		App:    commonApp,
		Server: nil,
	}

	// build HTTP server
	server := &http.Server{
		Addr:    cfg.Server.HTTPListen,
		Handler: app.buildHandler(),
	}
	app.Server = server

	return app
}

func (app *App) buildHandler() *routing.Router {
	router := routing.New()

	router.Use(
		accesslog.Handler(app.Infra.Logger),
		slash.Remover(http.StatusMovedPermanently),
		errorshandler.Handler(app.Infra.Logger),
		cors.Handler(cors.AllowAll),
	)
	//router.NotFound(file.Content("website/index.html"))

	// serve index file
	router.Get("/", file.Content("website/index.html"))
	router.Get("/a/*", file.Content("website/index.html"))
	router.Get("/u/*", file.Content("website/index.html"))

	router.Get("/favicon.ico", file.Content("website/favicon.ico"))
	router.Get("/manifest.json", file.Content("website/manifest.json"))
	// serve files under the "static" subdirectory
	router.Get("/static/*", file.Server(file.PathMap{
		"/static/": "/website/static/",
	}))

	api := router.Group("/api")
	api.Use(
		content.TypeNegotiator(content.JSON),
		ozzo_routing.SetHeader("Content-Type", "application/json; charset=UTF-8"),
	)

	authMiddleware := routing.Handler(func(context *routing.Context) error {
		return nil
	})
	//authMiddleware := auth.Middleware(app.Infra.Logger, app.Auth.Service)
	//
	//auth.RegisterHandlers(api.Group(""),
	//	app.Auth.Service,
	//	app.Infra.Logger,
	//)

	app.RegisterHandlers(api, authMiddleware)

	return router
}

// Run is func to run the ApiApp
func (app *App) Run() error {
	go func() {
		defer func() {
			if err := app.Stop(); err != nil {
				app.Infra.Logger.Error(err)
			}

			err := app.Infra.Logger.Sync()
			if err != nil {
				log.Println(err.Error())
			}
		}()
		// start the HTTP server with graceful shutdown
		routing.GracefulShutdown(app.Server, 10*time.Second, app.Infra.Logger.Infof)
	}()
	app.Infra.Logger.Infof("server %v is running at %v", Version, app.Server.Addr)
	if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// RegisterHandlers sets up the routing of the HTTP handlers.
func (app *App) RegisterHandlers(rg *routing.RouteGroup, authMiddleware routing.Handler) {
	//	Example
	rgTest := rg.Group("/test")
	controller.RegisterReferenceTestHandlers(rgTest, app.Infra.YaruzRepository, app.Infra.Logger, authMiddleware)
	controller.RegisterDataTestHandlers(rgTest, app.Infra.YaruzRepository, app.Domain.User, app.Domain.Advertiser, app.Domain.AdvertisingCampaign, app.Domain.Offer, app.Infra.Logger, authMiddleware)
	controller.RegisterAccountHandlers(rg, app.Domain.User, app.Infra.Logger, authMiddleware)
}
