package main

import (
	"context"
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
	_, cancel := context.WithCancel(context.Background())

	// setup graceful shutdown handler
	gracefulShutdown := make(chan os.Signal, 1)
	shutdownSave := make(chan int, 1)
	defer os.Exit(0)

	// intercept kill signals
	signal.Notify(gracefulShutdown, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	go api.Run()

	// wait for shutdown signal
	<-gracefulShutdown
	close(gracefulShutdown)

	go func() {
		log.Info("Graceful shutdown...")
		shutdownSave <- 1
	}()

	// wait for save to complete
	<-shutdownSave
	close(shutdownSave)

	cancel()
}
