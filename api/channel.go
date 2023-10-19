package api

//////////////////////////////////////////////////

type ChannelLivestreamCategory struct {
  ID         int64 `json:"id"`
  CategoryID int64 `json:"category_id"`

  Slug        string `json:"slug"`
  Name        string `json:"name"`
  Description string `json:"description"`
  Viewers     int    `json:"viewers"`

  Tags     []string `json:"tags,omitempty"`
  Category struct {
    ID   int64  `json:"id"`
    Slug string `json:"slug"`
    Name string `json:"name"`
    Icon string `json:"icon"`
  } `json:"category"`
  DeletedAt string `json:"deleted_at"`
}

type ChannelPreviousLivestream struct {
  ID        int64  `json:"id"`
  ChannelID int64  `json:"channel_id"`
  Slug      string `json:"slug"`

  SessionTitle string `json:"session_title"`
  Duration     int64  `json:"duration"`
  Thumbnail    *Image `json:"thumbnail,omitempty"`
  Video        *Video `json:"video,omitempty"`

  ViewerCount   int    `json:"viewer_count"`
  Views         int    `json:"views"`
  Language      string `json:"language"`
  Source        string `json:"source"`
  TwitchChannel string `json:"twitch_channel"`

  IsLive      bool `json:"is_live"`
  IsMature    bool `json:"is_mature"`
  RiskLevelID int  `json:"risk_level_id"`

  CreatedAt string `json:"created_at"`
}

type ChannelLivestream struct {
  ID        int64  `json:"id"`
  ChannelID int64  `json:"channel_id"`
  Slug      string `json:"slug"`
  IsLive    bool   `json:"is_live"`

  SessionTitle string `json:"session_title"`
  Thumbnail    *Image `json:"thumbnail,omitempty"`

  ViewerCount int                         `json:"viewer_count"`
  Viewers     int                         `json:"viewers"`
  IsMature    bool                        `json:"is_mature"`
  Categories  []ChannelLivestreamCategory `json:"categories,omitempty"`

  CreatedAt string `json:"created_at"`
}

type ChannelMedia struct {
  ID   int64  `json:"id"`
  UUID string `json:"uuid"`

  ModelID        int64  `json:"model_id"`
  ModelType      string `json:"model_type"`
  CollectionName string `json:"collection_name"`

  Name            string `json:"name"`
  Filename        string `json:"file_name"`
  MimeType        string `json:"mime_type"`
  Disk            string `json:"disk"`
  Size            int64  `json:"size"`
  ConversionsDisk string `json:"conversions_disk"`

  CreatedAt string `json:"created_at"`
  UpdatedAt string `json:"updated_at"`
}

type Channel struct {
  ID     int64  `json:"id"`
  UserID int64  `json:"user_id"`
  Slug   string `json:"slug"`

  PlaybackURL string `json:"playback_url"`

  Livestream *ChannelLivestream `json:"livestream,omitempty"`
  User       *User              `json:"user,omitempty"`
  Chatroom   *Chatroom          `json:"chatroom,omitempty"`

  BannerImage         *Image                      `json:"banner_image,omitempty"`
  PreviousLivestreams []ChannelPreviousLivestream `json:"previous_livestreams,omitempty"`
  Media               []ChannelMedia              `json:"media,omitempty"`

  FollowersCount      int  `json:"followersCount"`
  IsBanned            bool `json:"is_banned"`
  VODEnabled          bool `json:"vod_enabled"`
  SubscriptionEnabled bool `json:"subscription_enabled"`
  Muted               bool `json:"muted"`
}
