package main

import (
	"context"
	"github.com/defgenx/go-graphql-poc/core"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var doneChan = make(chan struct{}, 0)

func main() {
	// Init main context
	mainContext, cancelFunc := context.WithCancel(context.Background())
	// Init app
	app := initApp()
	// Run webserver ansyc
	go runWebServer(mainContext, app)
	// Listen system signal to stop goroutines by context
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	defer signal.Stop(c)
	select {
	case <-c:
		log.Infoln("Signal caught canceling contexts")
		// Cancel contexts and children
		cancelFunc()
		<-doneChan
		log.Infoln("Stopping program...")
	}
	os.Exit(0)
}

func runWebServer(ctx context.Context, app *core.App) {
	// Init server
	srv := &http.Server{
		Addr:         os.Getenv("SERVER_ADDR"),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.LoggingHandler(os.Stdout, app.Router),
	}

	log.Infof("Server started on: %s", srv.Addr)

	go closeWebServer(ctx, srv)

	castPeriod, err := initRetryServerPeriod()
	if err != nil {
		log.Fatal(err)
	}
	// Automatically restart server
	for {
		if err := srv.ListenAndServe(); err != nil {
			log.Errorf("Server fail to start: %s", err.Error())
		}
		log.Errorf("Server will retry in %f seconds", castPeriod.Seconds())
		time.Sleep(castPeriod * time.Second)
	}
}

func closeWebServer(ctx context.Context, srv *http.Server) {
	select {
	case <-ctx.Done():
		log.Info("Closing webserver...")
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Fatal(err)
		}
		doneChan <- struct{}{}
	}
}

func initApp() *core.App {
	// Initialize Application
	var MyApp = core.NewApp()
	// Execute route handler
	MyApp.HandleRoute()

	return MyApp
}

func initRetryServerPeriod() (castPeriod time.Duration, err error) {
	retryPeriodEnv := os.Getenv("SERVER_RETRY_PERIOD")
	castPeriod, err = time.ParseDuration(retryPeriodEnv)
	if err != nil {
		return 0, err
	}

	return castPeriod, nil
}
