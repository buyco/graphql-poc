package core

import (
	"github.com/buyco/keel/pkg/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type App struct {
	Router *Router
}

// App constructor
func NewApp() *App {
	loadEnv()
	loggerLevelFromEnv()
	app := &App{
		Router: NewRouter(),
	}

	return app
}

func (app *App) HandleRoute() {
	app.Router.SetRoutes()
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.Router.ServeHTTP(w, r)
}

// Load environment file
func loadEnv() {
	if env := os.Getenv("ENV"); env != "" {
		err := utils.LoadEnvFile(".env", env)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Set logger level from environement
func loggerLevelFromEnv() {
	switch os.Getenv("ENV") {
	case "production", "demo":
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}
