package twitch

import (
	"encoding/json"
	"fmt"
	"log"
	"modactionlistener/backend/internal/app/entity"

	"github.com/gempir/go-twitch-irc/v4"
)

func RegisterMessageHandler(client *twitch.Client) {
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		timestamp := message.Time.UTC().Format("2006-01-02 15:04:05")

		msg := entity.Message{
			Timestamp: timestamp,
			Channel:   message.Channel,
			ChannelID: message.RoomID,
			User:      message.User.DisplayName,
			UserID:    message.User.ID,
			Message:   message.Message,
		}

		// Marshal the Message struct into a JSON string
		messageBody, err := json.Marshal(msg)
		if err != nil {
			log.Printf("Failed to marshal message: %v", err)
			return
		}

		fmt.Printf("Message received: %s\n", string(messageBody))

		// Setup the message to be sent to the database 
		fmt.Printf("Mods?: %v\n", message.User.IsMod)

	})
}