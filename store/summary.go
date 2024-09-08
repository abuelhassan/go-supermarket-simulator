package store

import (
	"bytes"
	"fmt"

	"github.com/abuelhassan/go-supermarket-simulator/product"
)

type Summary struct {
	Revenue float32
	Bought  map[product.Product]int
	Missed  map[product.Product]int
}

func newSummary() Summary {
	return Summary{
		Revenue: 0,
		Bought:  make(map[product.Product]int),
		Missed:  make(map[product.Product]int),
	}
}

func (s *Summary) missed(p product.Product, cnt int) {
	s.Missed[p] += cnt
}

func (s *Summary) bought(p product.Product, cnt int) {
	s.Bought[p] += cnt
	s.Revenue += float32(cnt) * (p.SellPrice() - p.BuyPrice())
}

func (s *Summary) String() string {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Revenue: %f\n", s.Revenue))
	for k, v := range s.Bought {
		b.WriteString(fmt.Sprintf("\tSold %d of %s\n", v, k.Name()))
	}
	for k, v := range s.Missed {
		b.WriteString(fmt.Sprintf("\tShort on %s %d times\n", k.Name(), v))
	}
	return b.String()
}
