package json

// Task:
// Implement JSON encoding and decoding using Go's standard library.
//
// 1. Define a struct named Person with fields Name and Email.
// 2. Implement MarshalPerson to convert a Person struct into JSON.
// 3. Implement UnmarshalPerson to convert a JSON string into a Person struct.
// 4. Handle and return errors properly in both functions.

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// MarshalPerson should convert a Person struct to a JSON string.
func MarshalPerson(p Person) (string, error) {
	return "", nil
}

// UnmarshalPerson should convert a JSON string to a Person struct.
func UnmarshalPerson(jsonStr string) (Person, error) {
	return Person{}, nil
}
