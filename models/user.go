package models

// Defining User struct
type User struct {
	ID        int
	firstName string
	lastName  string
}

// Defining variable users and nextId
var (
	// Data will looks like this
	// {{ID:1, firstName: "Indiana", lastName: "Johnes"},{ID:2, firstName: "Amelia", lastName: "Ethart"}}
	users  []*User
	nextId = 1
)

// Function that will get all users
func getUsers() []*User {
	return users
}

// Function that will add user to the users list
func addUser(u User) (User, error) {
	u.ID = nextId
	nextId++
	// Why we need to use &u instead of *u or u?
	users = append(users, &u)
	return u, nil
}
