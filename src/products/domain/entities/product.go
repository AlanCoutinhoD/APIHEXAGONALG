package entities

type Product struct {
	ID    int
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{ID: 1, Name: name, Price: price}
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string) {
	p.Name = name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) SetPrice(price float64) {
	p.Price = price
}

func (p *Product) GetID() int {
	return p.ID
}

func (p *Product) SetID(id int) {
	p.ID = id
}	

func (p *Product) Delete() {
	p.ID = 0
	p.Name = ""
	p.Price = 0
}




