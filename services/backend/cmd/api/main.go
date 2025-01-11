package main

import (
	"context"
	"modactionlistener/backend/internal/app/twitch"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//fmt.Println(handler.LoadUsers())

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go twitch.Connect(ctx)

	<-stopChan
	cancel()
}
