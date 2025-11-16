package structs

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}
