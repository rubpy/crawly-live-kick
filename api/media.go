package api

//////////////////////////////////////////////////

type Image struct {
  Responsive string `json:"responsive,omitempty"`
  URL        string `json:"url,omitempty"`
  Source     string `json:"src,omitempty"`
  SourceSet  string `json:"srcset,omitempty"`
}

type Video struct {
  ID           int64  `json:"id"`
  LivestreamID int64  `json:"live_stream_id"`
  Slug         string `json:"slug"`
  UUID         string `json:"uuid"`

  Views             int    `json:"views"`
  Thumb             string `json:"thumb"`
  S3                string `json:"s3"`
  TradingPlatformID int64  `json:"trading_platform_id"`

  CreatedAt string `json:"created_at"`
  UpdatedAt string `json:"updated_at"`
  DeletedAt string `json:"deleted_at"`
}
