package smtp

import (
	"bytes"
	"log"
	"net"
	"net/mail"

	commonApp "smtp2api/internal/app"
	"smtp2api/internal/pkg/config"

	"github.com/mhale/smtpd"
)

// Version of API
const Version = "1.0.0"

// App is the application for API
type App struct {
	*commonApp.App
	Server *smtp2ApiServer
}

type smtp2ApiServer struct {
}

// New func is a constructor for the ApiApp
func New(commonApp *commonApp.App, cfg config.Configuration) *App {
	app := &App{
		App:    commonApp,
		Server: nil,
	}

	// build HTTP server
	server := &smtp2ApiServer{}
	app.Server = server

	return app
}

func (app *App) Run() {
	app.Server.Run()
}

func (s *smtp2ApiServer) Run() {
	smtpd.ListenAndServe("127.0.0.1:2525", s.mailHandler, "MyServerApp", "")
}

func (s *smtp2ApiServer) mailHandler(origin net.Addr, from string, to []string, data []byte) error {
	msg, _ := mail.ReadMessage(bytes.NewReader(data))
	subject := msg.Header.Get("Subject")
	log.Printf("Received mail from %s for %s with subject %s", from, to[0], subject)
	return nil
}
