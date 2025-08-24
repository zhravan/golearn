package struct_embedding

type Base struct {
	ID        int
	CreatedAt string
}

type User struct {
	Base
	Name  string
	Email string
}

func NewUser(id int, createdAt, name, email string) User {
	return User{
		Base:  Base{ID: id, CreatedAt: createdAt},
		Name:  name,
		Email: email,
	}
}
