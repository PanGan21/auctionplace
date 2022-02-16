package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PanGan21/auctionplace/cmd/command"
)

const FALLBACK_TIMEOUT = 3

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)
	defer func() {
		log.Println("closing app")
		signal.Stop(signals)
		cancel()
	}()

	go func() {
		<-signals
		log.Println("sending cancel signal")
		cancel()
		time.Sleep(FALLBACK_TIMEOUT * time.Second)
		os.Exit(1)
	}()

	cmd := command.NewRootCommand(ctx)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	cmd.Execute()
}
