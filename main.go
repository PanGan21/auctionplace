package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PanGan21/auctionplace/cmd/command"
)

const CANCEL_TIMEOUT = 3

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)
	defer func() {
		fmt.Println("Closing the app..")
		signal.Stop(signals)
		cancel()
	}()

	go func() {
		<-signals
		fmt.Println("Sending cancel signal..")
		time.Sleep(CANCEL_TIMEOUT * time.Second)
		os.Exit(1)
	}()

	cmd := command.NewRootCommand(ctx)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
