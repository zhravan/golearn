package interfaces

import "fmt"

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func SayHello(g Greeter) string {
	return g.Greet()
}

func main() {
	p := Person{Name: "Alice"}
	fmt.Println(SayHello(p))
}
