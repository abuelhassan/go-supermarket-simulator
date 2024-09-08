package store

import (
	"sync"

	"github.com/abuelhassan/go-supermarket-simulator/product"
)

type Store interface {
	Balance() float32
	PurchaseLicense(l product.License) bool
	PurchaseDisplay(d product.Display) bool
	BuyProduct(p product.Product, cnt int) bool
	UpdateBill(billID int, p product.Product, cnt int) bool
	CollectBill(billID int)
	AllowedProducts() []product.Product
	Summary() Summary
}

func New() Store {
	return &store{
		balance:  500,
		licenses: make(map[product.License]struct{}),
		displays: make(map[product.Display]struct{}),
		products: make(map[product.Product]int),
		bills:    make(map[int]float32),
		summary:  newSummary(),
	}
}

type store struct {
	mu       sync.Mutex
	balance  float32
	licenses map[product.License]struct{}
	displays map[product.Display]struct{}
	products map[product.Product]int
	bills    map[int]float32
	summary  Summary
}

func (s *store) Balance() float32 {
	return s.balance
}

func (s *store) PurchaseLicense(l product.License) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.licenses[l]; ok {
		return true
	}
	if l.Price() > s.balance {
		return false
	}
	s.balance -= l.Price()
	s.licenses[l] = struct{}{}
	for _, p := range l.Products() {
		s.products[p] = 0
	}
	return true
}

func (s *store) PurchaseDisplay(d product.Display) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.displays[d]; ok {
		return true
	}
	if d.Price() > s.balance {
		return false
	}
	s.balance -= d.Price()
	s.displays[d] = struct{}{}
	return true
}

func (s *store) BuyProduct(p product.Product, cnt int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.licenses[p.License()]; !ok {
		return false
	}
	if _, ok := s.displays[p.Display()]; !ok {
		return false
	}
	if p.BuyPrice()*float32(cnt) > s.balance {
		return false
	}

	s.balance -= p.BuyPrice() * float32(cnt)
	s.products[p] += cnt

	return true
}

func (s *store) UpdateBill(billID int, p product.Product, cnt int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.products[p] < cnt {
		s.summary.missed(p, cnt)
		return false
	}

	s.products[p] -= cnt
	s.bills[billID] += p.SellPrice() * float32(cnt)
	s.summary.bought(p, cnt)

	return true
}

func (s *store) CollectBill(billID int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.balance += s.bills[billID]
	delete(s.bills, billID)
}

func (s *store) AllowedProducts() []product.Product {
	res := make([]product.Product, 0, len(s.products))
	for p := range s.products {
		res = append(res, p)
	}
	return res
}

func (s *store) Summary() Summary {
	return s.summary
}
