package entity

// Message represents the structure of the message data
type Message struct {
	Timestamp string `json:"timestamp"`
	Channel   string `json:"channel"`
	ChannelID string `json:"channel_id"`
	User      string `json:"user"`
	UserID    string `json:"user_id"`
	Message   string `json:"message"`
}
