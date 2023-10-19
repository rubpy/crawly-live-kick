package kick

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/rubpy/crawly"
	"github.com/rubpy/crawly-live-kick/api"
	"github.com/rubpy/crawly/cclient"
)

//////////////////////////////////////////////////

type config struct {
	logger    *slog.Logger
	client    cclient.Client
	apiClient cclient.APIClient

	settings struct {
		v  CrawlerSettings
		ok bool
	}
}

var (
	NilConfig = errors.New("config is nil")
	NilClient = errors.New("client is nil")
)

func validateConfig(cfg *config) error {
	if cfg == nil {
		return NilConfig
	}

	return nil
}

func buildCrawlerFromConfig(cfg *config) (cr *Crawler, err error) {
	if cfg == nil {
		err = NilConfig
		return
	}

	cl := cfg.client
	if cl == nil {
		logger := cfg.logger
		if logger != nil {
			logger = logger.WithGroup("client")
		}

		cl, err = cclient.NewClient(cclient.WithLogger(logger))
		if err != nil {
			return nil, fmt.Errorf("cclient.NewClient: %w", err)
		}
	}

	acl := cfg.apiClient
	if acl == nil {
		logger := cfg.logger
		if logger != nil {
			logger = logger.WithGroup("apiClient")
		}

		acl, err = cclient.NewAPIClient(logger, cl, api.DefaultBaseURL, api.DefaultHeader)
		if err != nil {
			return nil, fmt.Errorf("cclient.NewAPIClient: %w", err)
		}
	}

	cr = &Crawler{
		client:    cl,
		apiClient: acl,
	}

	cr.Crawler.SetLogger(cfg.logger)
	crawly.SetCrawlerHandlers(&cr.Crawler, crawly.CrawlerHandlers{
		Order:  cr.orderHandler,
		Entity: cr.entityHandler,
	})

	if cfg.settings.ok {
		cr.SetSettings(cfg.settings.v)
	} else {
		cr.SetSettings(DefaultSettings)
	}

	return cr, nil
}

type ConfigOption func(cfg *config)

//////////////////////////////////////////////////

func WithLogger(logger *slog.Logger) ConfigOption {
	return func(cfg *config) {
		cfg.logger = logger
	}
}

func WithClient(client cclient.Client) ConfigOption {
	return func(cfg *config) {
		cfg.client = client
	}
}

func WithAPIClient(apiClient cclient.APIClient) ConfigOption {
	return func(cfg *config) {
		cfg.apiClient = apiClient
	}
}

func WithSettings(settings CrawlerSettings) ConfigOption {
	return func(cfg *config) {
		cfg.settings.v = settings
		cfg.settings.ok = true
	}
}
