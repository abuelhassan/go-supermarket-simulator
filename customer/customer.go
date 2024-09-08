package customer

import (
	"math/rand"
	"time"
)

const (
	maxItems = 7

	speedMinMillis = 500
	speedMaxMillis = 800
)

type Customer[T comparable] struct {
	Speed time.Duration
	Order Order[T]
}

type Order[T comparable] map[T]int

func New[T comparable](opts []T) Customer[T] {
	items, o := 1+rand.Intn(maxItems), make(Order[T])
	for i := 0; i < items; i++ {
		o[opts[rand.Intn(len(opts))]]++
	}
	speed := speedMinMillis + rand.Intn(speedMaxMillis-speedMinMillis)
	return Customer[T]{
		Speed: time.Duration(speed) * time.Millisecond,
		Order: o,
	}
}
