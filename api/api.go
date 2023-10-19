package api

import "net/http"

//////////////////////////////////////////////////

var DefaultBaseURL = "https://kick.com/api/"

var DefaultHeader = http.Header{
	"Content-Type":      {"application/json"},
	"X-Requested-Wwith": {"XMLHttpRequest"},
	"Referer":           {"https://kick.com/"},
	"Origin":            {"https://kick.com"},
}
