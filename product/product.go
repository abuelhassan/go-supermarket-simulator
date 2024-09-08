package product

type Product struct {
	id        int
	name      string
	buyPrice  float32
	sellPrice float32
	display   Display
	license   License
}

var products []Product

func (p Product) ID() int {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) BuyPrice() float32 {
	return p.buyPrice
}

func (p Product) SellPrice() float32 {
	return p.sellPrice
}

func (p Product) Display() Display {
	return p.display
}

func (p Product) License() License {
	return p.license
}

func init() {
	products = []Product{
		{id: 1, name: "Cereal", buyPrice: 2.50, sellPrice: 5.00, display: DisplayShelf, license: LicenseBasic},
		{id: 2, name: "Bread", buyPrice: 1.15, sellPrice: 5.00, display: DisplayShelf, license: LicenseBasic},
		{id: 3, name: "Flour", buyPrice: 1.00, sellPrice: 3.75, display: DisplayShelf, license: LicenseBasic},
		{id: 4, name: "Oil", buyPrice: 1.75, sellPrice: 4.50, display: DisplayShelf, license: LicenseBasic},
		{id: 5, name: "Pasta", buyPrice: 1.00, sellPrice: 3.00, display: DisplayShelf, license: LicenseBasic},
		{id: 6, name: "Sugar", buyPrice: 1.00, sellPrice: 4.65, display: DisplayShelf, license: LicenseBasic},

		{id: 7, name: "Water", buyPrice: 0.75, sellPrice: 2.00, display: DisplayFridge, license: LicenseDrinks},
		{id: 8, name: "Ice Coffee", buyPrice: 2.00, sellPrice: 5.50, display: DisplayFridge, license: LicenseDrinks},
		{id: 9, name: "Limeade Soda", buyPrice: 0.65, sellPrice: 2.75, display: DisplayFridge, license: LicenseDrinks},
		{id: 10, name: "Cola Soda", buyPrice: 0.65, sellPrice: 2.75, display: DisplayFridge, license: LicenseDrinks},
		{id: 11, name: "Apple Juice", buyPrice: 0.85, sellPrice: 2.80, display: DisplayFridge, license: LicenseDrinks},
		{id: 12, name: "Orange Juice", buyPrice: 0.85, sellPrice: 2.80, display: DisplayFridge, license: LicenseDrinks},
	}
}
