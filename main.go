package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/abuelhassan/go-supermarket-simulator/product"
	"github.com/abuelhassan/go-supermarket-simulator/simulator"
	"github.com/abuelhassan/go-supermarket-simulator/store"
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

	st := store.New()
	initializeStore(st)

	var wg sync.WaitGroup

	wg.Add(1)
	simulator.RunDay(ctx, &wg, st)

	wg.Wait()
	summary := st.Summary()
	fmt.Println(&summary)
}

func initializeStore(st store.Store) {
	st.PurchaseLicense(product.LicenseBasic)
	st.PurchaseDisplay(product.DisplayShelf)

	prods := st.AllowedProducts()
	for _, v := range prods {
		st.BuyProduct(v, 15)
	}
}
