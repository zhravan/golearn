package xml

import "encoding/xml"

// Person represents a person with basic information
type Person struct {
	Name  string `xml:"name"`
	Age   int    `xml:"age"`
	Email string `xml:"email"`
}

// Book represents a book with title, author, and year
type Book struct {
	Title  string `xml:"title"`
	Author string `xml:"author"`
	Year   int    `xml:"year"`
}

// MarshalPerson converts a Person struct to XML bytes
func MarshalPerson(p Person) ([]byte, error) {
	return xml.Marshal(p)
}

// UnmarshalPerson converts XML bytes to a Person struct
func UnmarshalPerson(data []byte) (Person, error) {
	var p Person
	err := xml.Unmarshal(data, &p)
	return p, err
}

// MarshalBook converts a Book struct to XML bytes
func MarshalBook(b Book) ([]byte, error) {
	return xml.Marshal(b)
}

// UnmarshalBook converts XML bytes to a Book struct
func UnmarshalBook(data []byte) (Book, error) {
	var b Book
	err := xml.Unmarshal(data, &b)
	return b, err
}
