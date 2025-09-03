package interfaces

// TODO:
// - Define an interface with a Greet() string method.
// - Implement the interface on a Person type.
// - Implement SayHello to accept a Greeter and return its greeting.
// - main demonstrates usage; do not change signatures used in tests.

type Greeter interface {
    Greet() string
}

type Person struct {
    Name string
}

func (p Person) Greet() string {
    // TODO: return a greeting that uses p.Name
    return ""
}

func SayHello(g Greeter) string {
    // TODO: forward the greeting from the Greeter
    return ""
}
// (no main in skeleton)
