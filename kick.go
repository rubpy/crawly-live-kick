package kick

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/rubpy/crawly-live-kick/api"
)

//////////////////////////////////////////////////

var (
	InvalidChannelSlug    = errors.New("invalid channel slug")
	InvalidPlaybackURL    = errors.New("invalid playback URL")
	UnexpectedIVSResponse = errors.New("unexpected IVS response")
)

func (cr *Crawler) CheckIVS(ctx context.Context, playbackURL string) (available bool, tokenValid bool, err error) {
	if !isValidURL(playbackURL) {
		err = InvalidPlaybackURL
		return
	}

	if cr.client == nil {
		err = NilClient
		return
	}

	if ctx == nil {
		ctx = context.Background()
	} else {
		if err = ctx.Err(); err != nil {
			return
		}
	}

	resp, err := cr.client.Request(ctx, "GET", playbackURL, nil, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		available = true
		tokenValid = true
		err = nil
	} else if resp.StatusCode == http.StatusNotFound {
		available = false
		tokenValid = true
		err = nil
	} else {
		available = false
		tokenValid = false
		err = UnexpectedIVSResponse
	}

	return
}

func (cr *Crawler) FetchChannel(ctx context.Context, channelSlug string) (channel *api.Channel, err error) {
	if !IsValidChannelSlug(channelSlug) {
		err = InvalidChannelSlug
		return
	}

	if ctx == nil {
		ctx = context.Background()
	} else {
		if err = ctx.Err(); err != nil {
			return
		}
	}

	uri, err := url.JoinPath("/v1/channels/", strings.ToLower(channelSlug))
	if err != nil {
		return
	}

	resp, err := cr.apiClient.Request(ctx, "GET", uri, nil, nil, nil)
	if err != nil {
		return
	}
	defer resp.Close()

	channel = &api.Channel{}
	if err = resp.Decode(channel); err != nil {
		return
	}

	return
}

func isValidURL(s string) bool {
	if s == "" {
		return false
	}

	u, err := url.Parse(s)
	if err != nil || u.Scheme == "" {
		return false
	}

	return true
}

func IsValidChannelSlug(channelSlug string) bool {
	for _, r := range channelSlug {
		if !((r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r == '-' || r == '_')) {
			return false
		}
	}

	return true
}

func ExtractChannelSlug(channelURL string) (slug string, ok bool) {
	if channelURL == "" {
		return
	}

	u, err := url.Parse(channelURL)
	if err != nil {
		return
	}

	h := u.Hostname()
	if !strings.Contains(h, "kick.") {
		return
	}

	path := strings.Trim(u.Path, "/ ")
	if path == "" {
		return
	}

	if strings.ContainsRune(path, '/') {
		parts := strings.SplitN(path, "/", 2)
		if len(parts) == 0 {
			return
		}

		path = parts[0]
	}

	if !IsValidChannelSlug(path) {
		return
	}

	return path, true
}
