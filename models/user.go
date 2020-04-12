package models

// User represents a user struct
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID int = 1
)

// GetUsers returns users
func GetUsers() []*User {
	return users
}

// AddUser adds a user
func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
