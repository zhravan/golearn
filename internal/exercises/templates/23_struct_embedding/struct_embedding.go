package struct_embedding

// TODO:
// - Use struct embedding to compose types.
// - Base should hold common fields; User embeds Base and adds Name/Email.
// - Implement NewUser constructor returning a populated User.

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
    // TODO: return a properly initialized User embedding Base
    return User{}
}
