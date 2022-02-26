package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/thatDAMNbobby/farmercookbook/appcontext"
	"github.com/thatDAMNbobby/farmercookbook/config"
	"github.com/thatDAMNbobby/farmercookbook/servers"
	"github.com/thatDAMNbobby/farmercookbook/servers/runnable"
	"github.com/thatDAMNbobby/farmercookbook/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	configImpl := config.Loader()
	log.SetReportCaller(configImpl.LogLevel == log.TraceLevel)
	log.SetLevel(configImpl.LogLevel)
	utils.PrintDebugJSON("config", configImpl)
	appcontextImpl := appcontext.New(configImpl)

	serversImpl := servers.New(
		&servers.Deps{Handlers: appcontextImpl.Handlers},
		&configImpl.Server,
	)

	waitForInterrupts(configImpl.Name, serversImpl, appcontextImpl)
}

func waitForInterrupts(name string, servers runnable.Runnable, appcontext runnable.Runnable) {
	wait := time.Second * 15
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	defer os.Exit(0)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go start(ctx, name, servers, appcontext)

	sig := <-signalChannel // Block until we receive an interrupt signal
	log.WithField("signal", sig).Info("Received signal, shutting down")

	// This sleep facilitates a cleaner shutdown of the k8s node the app runs in
	if sig == syscall.SIGTERM {
		sleepDuration := time.Second * 28
		log.WithField("duration", sleepDuration).Info("Sleeping")
		time.Sleep(sleepDuration)
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), wait)
	defer shutdownCancel()

	servers.Stop(shutdownCtx)
	appcontext.Stop(shutdownCtx)

	log.Info("Done shutting down")
}

func start(ctx context.Context, name string, servers runnable.Runnable, appcontext runnable.Runnable) {
	log.WithField("name", name).Info("App starting")
	appcontext.Start(ctx)
	servers.Start(ctx)
}
