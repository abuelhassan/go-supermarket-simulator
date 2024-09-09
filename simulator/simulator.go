package simulator

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/abuelhassan/go-supermarket-simulator/conn"
	"github.com/abuelhassan/go-supermarket-simulator/customer"
	"github.com/abuelhassan/go-supermarket-simulator/generator"
	"github.com/abuelhassan/go-supermarket-simulator/product"
	"github.com/abuelhassan/go-supermarket-simulator/store"
)

const (
	customerTick = 300 * time.Millisecond
	dayDur       = 20 * time.Second

	storeCap = 2
)

type picked struct {
	p   product.Product
	cnt int
}

func RunDay(ctx context.Context, st store.Store) {
	cm := conn.NewManager[picked](storeCap)
	gen := generator.New(customerTick, dayDur)

	var wg sync.WaitGroup

	gen.Run(ctx, func() {
		wg.Add(1)
		go run(ctx, &wg, cm, st)
	})

	wg.Wait()
}

func run(ctx context.Context, wg *sync.WaitGroup, cm conn.Manager[picked], st store.Store) {
	defer wg.Done()

	cn := cm.CreateConn(ctx)
	if cn == nil {
		fmt.Println("Customer dismissed")
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for m := range cn.C() {
			if ok := st.UpdateBill(cn.ID(), m.p, m.cnt); ok {
				fmt.Printf("Customer %d, Picked up %d item(s) of %s.\n", cn.ID(), m.cnt, m.p.Name())
			} else {
				fmt.Printf("Customer %d didn't find %s\n", cn.ID(), m.p.Name())
			}
		}
	}()

	cus := customer.New(st.AllowedProducts())
	for k, v := range cus.Order {
		time.Sleep(cus.Speed)
		cn.C() <- picked{p: k, cnt: v}
	}
	cn.Done()
}
