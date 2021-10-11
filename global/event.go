package global

import (
	"context"
	"fmt"
	"github.com/xtpv/xtp-gin/pkg/events"
)

var Event *events.Event

const Sync = "sync"

func StartEvent(ctx context.Context) {
	Event = events.New(ctx)
	Event.Start()

	Event.Bind(Sync, func(arg interface{}) error {

		fmt.Println(arg)
		return nil
	})
}
