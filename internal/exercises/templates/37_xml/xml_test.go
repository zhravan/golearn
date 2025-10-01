package xml

import (
	"strings"
	"testing"
)

func TestMarshalPerson(t *testing.T) {
	person := Person{
		Name:  "Alice",
		Age:   30,
		Email: "alice@example.com",
	}

	data, err := MarshalPerson(person)
	if err != nil {
		t.Fatalf("MarshalPerson() error = %v", err)
	}

	xmlStr := string(data)
	if !strings.Contains(xmlStr, "<name>Alice</name>") {
		t.Errorf("Expected XML to contain <name>Alice</name>, got %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<age>30</age>") {
		t.Errorf("Expected XML to contain <age>30</age>, got %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<email>alice@example.com</email>") {
		t.Errorf("Expected XML to contain <email>alice@example.com</email>, got %s", xmlStr)
	}
}

func TestUnmarshalPerson(t *testing.T) {
	xmlData := []byte(`<Person><name>Bob</name><age>25</age><email>bob@example.com</email></Person>`)

	person, err := UnmarshalPerson(xmlData)
	if err != nil {
		t.Fatalf("UnmarshalPerson() error = %v", err)
	}

	if person.Name != "Bob" {
		t.Errorf("Expected Name to be Bob, got %s", person.Name)
	}
	if person.Age != 25 {
		t.Errorf("Expected Age to be 25, got %d", person.Age)
	}
	if person.Email != "bob@example.com" {
		t.Errorf("Expected Email to be bob@example.com, got %s", person.Email)
	}
}

func TestMarshalBook(t *testing.T) {
	book := Book{
		Title:  "Go Programming",
		Author: "John Doe",
		Year:   2023,
	}

	data, err := MarshalBook(book)
	if err != nil {
		t.Fatalf("MarshalBook() error = %v", err)
	}

	xmlStr := string(data)
	if !strings.Contains(xmlStr, "<title>Go Programming</title>") {
		t.Errorf("Expected XML to contain <title>Go Programming</title>, got %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<author>John Doe</author>") {
		t.Errorf("Expected XML to contain <author>John Doe</author>, got %s", xmlStr)
	}
	if !strings.Contains(xmlStr, "<year>2023</year>") {
		t.Errorf("Expected XML to contain <year>2023</year>, got %s", xmlStr)
	}
}

func TestUnmarshalBook(t *testing.T) {
	xmlData := []byte(`<Book><title>Learning XML</title><author>Jane Smith</author><year>2024</year></Book>`)

	book, err := UnmarshalBook(xmlData)
	if err != nil {
		t.Fatalf("UnmarshalBook() error = %v", err)
	}

	if book.Title != "Learning XML" {
		t.Errorf("Expected Title to be Learning XML, got %s", book.Title)
	}
	if book.Author != "Jane Smith" {
		t.Errorf("Expected Author to be Jane Smith, got %s", book.Author)
	}
	if book.Year != 2024 {
		t.Errorf("Expected Year to be 2024, got %d", book.Year)
	}
}

func TestMarshalUnmarshalRoundTrip(t *testing.T) {
	originalPerson := Person{
		Name:  "Charlie",
		Age:   35,
		Email: "charlie@example.com",
	}

	// Marshal
	data, err := MarshalPerson(originalPerson)
	if err != nil {
		t.Fatalf("MarshalPerson() error = %v", err)
	}

	// Unmarshal
	parsedPerson, err := UnmarshalPerson(data)
	if err != nil {
		t.Fatalf("UnmarshalPerson() error = %v", err)
	}

	// Compare
	if parsedPerson.Name != originalPerson.Name {
		t.Errorf("Name mismatch: got %s, want %s", parsedPerson.Name, originalPerson.Name)
	}
	if parsedPerson.Age != originalPerson.Age {
		t.Errorf("Age mismatch: got %d, want %d", parsedPerson.Age, originalPerson.Age)
	}
	if parsedPerson.Email != originalPerson.Email {
		t.Errorf("Email mismatch: got %s, want %s", parsedPerson.Email, originalPerson.Email)
	}
}
