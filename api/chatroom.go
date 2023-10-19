package api

//////////////////////////////////////////////////

type Chatroom struct {
  ID        int64 `json:"id"`
  ChannelID int64 `json:"channel_id"`

  ChatableType string `json:"chatable_type"`
  ChatableID   int64  `json:"chatable_id"`

  ChatModeOld     string `json:"chat_mode_old"`
  ChatMode        string `json:"chat_mode"`
  SlowMode        bool   `json:"slow_mode"`
  FollowersMode   bool   `json:"followers_mode"`
  SubscribersMode bool   `json:"subscribers_mode"`
  EmotesMode      bool   `json:"emotes_mode"`

  MessageInterval      int `json:"message_interval"`
  FollowingMinDuration int `json:"following_min_duration"`

  CreatedAt string `json:"created_at"`
  UpdatedAt string `json:"updated_at"`
}
