package entity

// Define the nested structs that match the JSON structure
type UserResponse struct {
	Banned        bool          `json:"banned"`
	DisplayName   string        `json:"displayName"`
	Login         string        `json:"login"`
	ID            string        `json:"id"`
	Bio           string        `json:"bio"`
	Followers     int           `json:"followers"`
	ChatColor     string        `json:"chatColor"`
	Logo          string        `json:"logo"`
	Banner        string        `json:"banner"`
	CreatedAt     string        `json:"createdAt"`
	UpdatedAt     string        `json:"updatedAt"`
	EmotePrefix   string        `json:"emotePrefix"`
	Roles         Roles         `json:"roles"`
	Badges        []Badge       `json:"badges"`
	ChatterCount  int           `json:"chatterCount"`
	ChatSettings  ChatSettings  `json:"chatSettings"`
	LastBroadcast LastBroadcast `json:"lastBroadcast"`
	Panels        []Panel       `json:"panels"`
}

type Roles struct {
	IsAffiliate bool `json:"isAffiliate"`
	IsPartner   bool `json:"isPartner"`
	IsStaff     bool `json:"isStaff"`
}

type Badge struct {
	SetID       string `json:"setID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type ChatSettings struct {
	ChatDelayMs                  int      `json:"chatDelayMs"`
	FollowersOnlyDurationMinutes int      `json:"followersOnlyDurationMinutes"`
	SlowModeDurationSeconds      int      `json:"slowModeDurationSeconds"`
	BlockLinks                   bool     `json:"blockLinks"`
	IsSubscribersOnlyModeEnabled bool     `json:"isSubscribersOnlyModeEnabled"`
	IsEmoteOnlyModeEnabled       bool     `json:"isEmoteOnlyModeEnabled"`
	IsFastSubsModeEnabled        bool     `json:"isFastSubsModeEnabled"`
	IsUniqueChatModeEnabled      bool     `json:"isUniqueChatModeEnabled"`
	RequireVerifiedAccount       bool     `json:"requireVerifiedAccount"`
	Rules                        []string `json:"rules"`
}

type LastBroadcast struct {
	StartedAt string `json:"startedAt"`
	Title     string `json:"title"`
}

type Panel struct {
	ID string `json:"id"`
}