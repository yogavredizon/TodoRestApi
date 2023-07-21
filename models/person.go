package models

type Person struct {
	Name string
}

func (p *Person) GetName() string {
	return p.Name
}
