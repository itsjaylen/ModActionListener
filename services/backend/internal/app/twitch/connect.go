package twitch

import (
	"context"
	"fmt"

	"github.com/gempir/go-twitch-irc/v4"
)

func Connect(ctx context.Context) error {
	client := twitch.NewAnonymousClient()

	// Register twitch handler
	RegisterBanHandler(client)
	RegisterMessageHandler(client)


	client.Join("agent00")

	fmt.Printf("Connected to Twitch IRC as anonymous user\n")

	err := client.Connect()
	if err != nil {
		panic(err)
	}

	return nil
}
