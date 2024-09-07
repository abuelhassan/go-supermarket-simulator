package generator

import (
	"context"
	"sync"
	"time"
)

type Generator interface {
	Run(ctx context.Context, wg *sync.WaitGroup, onTick func())
}

func New(every time.Duration, after time.Duration) Generator {
	return &generator{
		every: every,
		after: after,
	}
}

type generator struct {
	every time.Duration
	after time.Duration
}

func (g generator) Run(ctx context.Context, wg *sync.WaitGroup, onTick func()) {
	defer wg.Done()

	tick, timer := time.NewTicker(g.every), time.NewTimer(g.after)
	defer tick.Stop()
	defer timer.Stop()

	for {
		select {
		case <-tick.C:
			onTick()
		case <-timer.C:
			return
		case <-ctx.Done():
			return
		}
	}
}
