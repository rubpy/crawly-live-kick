package kick

import (
	"time"

	"github.com/rubpy/crawly"
)

//////////////////////////////////////////////////

type CrawlerSettings struct {
	crawly.CrawlerSettings

	MaximumChannelPlaybackURLAge time.Duration
	MinimumFetchChannelDelay     time.Duration
}

var DefaultSettings = CrawlerSettings{
	CrawlerSettings: crawly.DefaultCrawlerSettings,

	MaximumChannelPlaybackURLAge: 50 * time.Minute,
	MinimumFetchChannelDelay:     5 * time.Minute,
}

//////////////////////////////////////////////////

func (cr *Crawler) loadSettings() CrawlerSettings {
	return cr.settings.Load()
}

func (cr *Crawler) setSettings(settings CrawlerSettings) {
	cr.settings.Store(settings)
	crawly.SetCrawlerSettings(&cr.Crawler, settings.CrawlerSettings)
}

func (cr *Crawler) Settings() CrawlerSettings {
	return cr.loadSettings()
}

func (cr *Crawler) SetSettings(settings CrawlerSettings) {
	cr.setSettings(settings)
}
