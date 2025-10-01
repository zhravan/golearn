package xml

// Person represents a person with basic information
type Person struct {
	// TODO: Add XML struct tags to define how this struct maps to XML
	// Use `xml:"name"` for the Name field
	// Use `xml:"age"` for the Age field
	// Use `xml:"email"` for the Email field
	Name  string
	Age   int
	Email string
}

// Book represents a book with title, author, and year
type Book struct {
	// TODO: Add XML struct tags
	// Use `xml:"title"` for Title
	// Use `xml:"author"` for Author
	// Use `xml:"year"` for Year
	Title  string
	Author string
	Year   int
}

// MarshalPerson converts a Person struct to XML bytes
// TODO: Implement this function using xml.Marshal
func MarshalPerson(p Person) ([]byte, error) {
	// TODO: Use xml.Marshal to convert the Person struct to XML
	// Import "encoding/xml" package
	return nil, nil
}

// UnmarshalPerson converts XML bytes to a Person struct
// TODO: Implement this function using xml.Unmarshal
func UnmarshalPerson(data []byte) (Person, error) {
	// TODO: Use xml.Unmarshal to parse XML data into a Person struct
	// Import "encoding/xml" package
	var p Person
	return p, nil
}

// MarshalBook converts a Book struct to XML bytes
// TODO: Implement this function using xml.Marshal
func MarshalBook(b Book) ([]byte, error) {
	// TODO: Use xml.Marshal to convert the Book struct to XML
	// Import "encoding/xml" package
	return nil, nil
}

// UnmarshalBook converts XML bytes to a Book struct
// TODO: Implement this function using xml.Unmarshal
func UnmarshalBook(data []byte) (Book, error) {
	// TODO: Use xml.Unmarshal to parse XML data into a Book struct
	// Import "encoding/xml" package
	var b Book
	return b, nil
}
