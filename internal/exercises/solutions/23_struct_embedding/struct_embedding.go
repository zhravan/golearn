package struct_embedding

// Base represents common fields shared across entity types.
type Base struct {
	ID        int
	CreatedAt string
}

// User embeds Base and adds user-specific fields.
type User struct {
	Base
	Name  string
	Email string
}

// NewUser creates a new User with all fields populated.
func NewUser(id int, createdAt, name, email string) User {
	return User{
		Base: Base{
			ID:        id,
			CreatedAt: createdAt,
		},
		Name:  name,
		Email: email,
	}
}
