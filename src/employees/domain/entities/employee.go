package entities

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	
}

func NewEmployee(name string) *Employee {
	return &Employee{ID: 1, Name: name}
}

func (e *Employee) GetName() string {
	return e.Name
}

func (e *Employee) SetName(name string) {
	e.Name = name
}
