package twitch

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v4"
)

// RegisterBanHandler handles bans and timeouts
func RegisterBanHandler(client *twitch.Client) {
	client.OnClearChatMessage(func(message twitch.ClearChatMessage) {
		if message.TargetUsername != "" {
			if message.BanDuration == 0 {
				fmt.Printf(
					"User %s has been permanently banned in channel %s\n",
					message.TargetUsername,
					message.Channel,
				)
			} else {
				fmt.Printf("User %s has been timed out for %d seconds in channel %s\n", message.TargetUsername, message.BanDuration, message.Channel)
			}
		} else {
			fmt.Printf("Chat was cleared in channel %s\n", message.Channel)
		}
	})
}