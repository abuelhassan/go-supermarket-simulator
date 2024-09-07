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

type Customer struct {
	Speed time.Duration
	Order Order
}

type Order map[int]int

func New(opts []int) Customer {
	items, o := 1+rand.Intn(maxItems), make(Order)
	for i := 0; i < items; i++ {
		o[rand.Intn(len(opts))]++
	}
	speed := speedMinMillis + rand.Intn(speedMaxMillis-speedMinMillis)
	return Customer{
		Speed: time.Duration(speed) * time.Millisecond,
		Order: o,
	}
}
