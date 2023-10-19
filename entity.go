package kick

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/rubpy/crawly"
	"github.com/rubpy/crawly-live-kick/api"
	"github.com/rubpy/crawly/clog"
)

//////////////////////////////////////////////////

type EntityData struct {
	Live bool `json:"live"`

	Channel     *api.Channel `json:"channel"`
	PlaybackURL string       `json:"playback_url"`
	FetchedAt   time.Time    `json:"fetched_at"`
}

func (cr *Crawler) entityHandler(ctx context.Context, entity *crawly.Entity, result *crawly.TrackingResult) error {
	handle, ok := entity.Handle.(Handle)
	if !ok || !handle.Valid() {
		return crawly.InvalidHandle
	}

	data, _ := entity.Data.(EntityData)
	defer func() {
		entity.Data = data
	}()

	if handle.Type != HandleChannelSlug {
		return crawly.InvalidHandle
	}

	settings := cr.loadSettings()
	max := func(min time.Duration, v time.Duration) time.Duration {
		if v < min {
			return min
		}

		return v
	}

	maximumChannelPlaybackURLAge := max(1*time.Minute, settings.MaximumChannelPlaybackURLAge)
	minimumFetchChannelDelay := max(1*time.Minute, settings.MinimumFetchChannelDelay)

	now := time.Now()
	if data.PlaybackURL == "" || now.Sub(data.FetchedAt) > maximumChannelPlaybackURLAge {
		data.PlaybackURL = ""
		data.Live = false

		if now.Sub(data.FetchedAt) >= minimumFetchChannelDelay {
			lp := clog.Params{
				Message: "fetchChannel",
				Level:   slog.LevelDebug,

				Values: clog.ParamGroup{
					"channelSlug": handle.Value,
				},
			}

			channel, err := cr.FetchChannel(ctx, handle.Value)
			if err == nil {
				data.FetchedAt = now

				data.Channel = channel
				data.PlaybackURL = channel.PlaybackURL
			} else {
				err = fmt.Errorf("FetchChannel: %w", err)
			}

			lp.Err = err
			cr.Log(ctx, lp)

			if err != nil {
				return err
			}
		}
	}

	if data.PlaybackURL != "" {
		data.Live = false

		lp := clog.Params{
			Message: "checkIVS",
			Level:   slog.LevelDebug,

			Values: clog.ParamGroup{
				"playbackURL": data.PlaybackURL,
			},
		}

		available, tokenValid, err := cr.CheckIVS(ctx, data.PlaybackURL)
		if err == nil {
			if tokenValid {
				data.Live = available
			} else {
				data.PlaybackURL = ""
			}

			lp.Set("available", available)
			lp.Set("tokenValid", tokenValid)
		} else {
			err = fmt.Errorf("CheckIVS: %w", err)
		}

		lp.Err = err
		cr.Log(ctx, lp)

		if err != nil {
			return err
		}
	}

	return nil
}
