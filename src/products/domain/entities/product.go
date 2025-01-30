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
