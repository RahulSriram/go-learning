package main

import "fmt"

// This `person` struct type has `name` and `age` fields.
type person struct {
	name string
	age  int
}

func (p *person) GetName string {
	return p.name
}

func (p *person) GetAge int {
	return p.age
}

func (p person) ReturnPerson person {
	return p
}

func main() {
	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})
	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})
	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})
	// An `&` prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})
	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	p := s
	fmt.Println(p.GetName)
	s.name = "qwert"
	fmt.Println(s.GetName)
	fmt.Println(p.GetName)
	// You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.GetAge)
	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.GetAge)
}
