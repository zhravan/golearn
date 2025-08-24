package closures

import "fmt"

func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	next := closure()
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	anotherNext := closure()
	fmt.Println(anotherNext())
	fmt.Println(anotherNext())
}
