package api

//////////////////////////////////////////////////

type User struct {
  ID       int64  `json:"id"`
  Username string `json:"username"`

  Bio            string `json:"bio"`
  ProfilePicture string `json:"profile_pic"`

  Country string `json:"country"`
  State   string `json:"state"`
  City    string `json:"city"`

  Instagram string `json:"instagram"`
  Twitter   string `json:"twitter"`
  YouTube   string `json:"youtube"`
  Discord   string `json:"discord"`
  TikTok    string `json:"tiktok"`
  Facebook  string `json:"facebook"`

  AgreedToTerms   bool   `json:"agreed_to_terms"`
  EmailVerifiedAt string `json:"email_verified_at"`
}
