package product

type Display struct {
	id    display
	price float32
}

type display int

const (
	dShelf display = iota + 1
	dFridge
)

var (
	DisplayShelf  = Display{id: dShelf, price: 200.0}
	DisplayFridge = Display{id: dFridge, price: 300.0}
)

func (d Display) String() string {
	switch d.id {
	case dShelf:
		return "shelf"
	case dFridge:
		return "fridge"
	}
	return "invalid"
}

func (d Display) Price() float32 {
	return d.price
}
