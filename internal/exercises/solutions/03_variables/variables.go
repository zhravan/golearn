package variables

// MakePerson returns two variables (name and age) initialized in function scope
func MakePerson() (string, int) {
	// Use short variable declarations
	name := "Gopher"
	age := 10
	return name, age
}
