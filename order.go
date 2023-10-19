package kick

import (
	"context"

	"github.com/rubpy/crawly"
)

//////////////////////////////////////////////////

type OrderData struct{}

func (cr *Crawler) orderHandler(ctx context.Context, order *crawly.Order, result *crawly.TrackingResult) error {
	handle, ok := order.Handle.(Handle)
	if !ok || !handle.Valid() {
		return crawly.InvalidHandle
	}

	data, _ := order.Data.(OrderData)
	defer func() {
		order.Data = data
	}()

	switch handle.Type {
	case HandleChannelSlug:
		if !IsValidChannelSlug(handle.Value) {
			return crawly.InvalidHandle
		}

	default:
		return crawly.InvalidHandle
	}

	return nil
}
