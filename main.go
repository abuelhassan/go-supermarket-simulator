package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/abuelhassan/go-supermarket-simulator/conn"
	"github.com/abuelhassan/go-supermarket-simulator/customer"
	"github.com/abuelhassan/go-supermarket-simulator/generator"
)

const (
	customerTick = 300 * time.Millisecond
	dayDur       = 20 * time.Second

	storeCap = 2
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-c
		cancel()
		fmt.Println("Shutting Down...")
		time.Sleep(3 * time.Second)
	}()

	cm := conn.NewManager[int](storeCap)
	gen := generator.New(customerTick, dayDur)

	var wg sync.WaitGroup

	wg.Add(1)
	go gen.Run(ctx, &wg, func() {
		wg.Add(1)
		go customerManager(ctx, &wg, cm)
	})

	wg.Wait()
}

func customerManager(ctx context.Context, wg *sync.WaitGroup, cm conn.Manager[int]) {
	defer wg.Done()

	cn := cm.CreateConn(ctx, wg)
	if cn == nil {
		fmt.Println("Customer dismissed")
		return
	}

	go func() {
		for m := range cn.C() {
			fmt.Printf("Customer %d, Picked up %d item(s).\n", cn.ID(), m)
		}
	}()

	cus := customer.New([]int{0, 1, 2, 3, 4, 5})
	for _, v := range cus.Order {
		time.Sleep(cus.Speed)
		cn.C() <- v
	}
	cn.Done()
}
