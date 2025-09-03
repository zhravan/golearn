package structs

// TODO:
// - Define a Person struct with Name and Age fields.
// - Implement a constructor-like function NewPerson(name, age) that returns a Person.

type Person struct {
    Name string
    Age  int
}

func NewPerson(name string, age int) Person {
	// TODO: return a properly initialized Person
	return Person{}
}
