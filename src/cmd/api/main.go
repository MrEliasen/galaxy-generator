package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/mreliasen/ihniwiad/pkg/api"
	"github.com/mreliasen/ihniwiad/pkg/logger"
)

func main() {
	log := logger.New(log.WarnLevel)

	godotenv.Load()

	// setup graceful shutdown handler
	gracefulShutdown := make(chan os.Signal, 1)
	defer os.Exit(0)

	// intercept kill signals
	signal.Notify(gracefulShutdown, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	go api.Run()

	// wait for shutdown signal
	<-gracefulShutdown
	close(gracefulShutdown)
	log.Info("Shutdown signal received.")
}
