package product

type License struct {
	id    license
	price float32
}

type license int

const (
	lBasic license = iota + 1
	lDrinks
)

var (
	LicenseBasic  = License{id: lBasic, price: 0.0}
	LicenseDrinks = License{id: lDrinks, price: 200.0}
)

func (l License) String() string {
	switch l.id {
	case lBasic:
		return "basic"
	case lDrinks:
		return "drinks"
	}
	return "invalid"
}

func (l License) Price() float32 {
	return l.price
}

func (l License) Products() []Product {
	res := make([]Product, 0)
	for _, p := range products {
		if p.license.id == l.id {
			res = append(res, p)
		}
	}
	return res
}
