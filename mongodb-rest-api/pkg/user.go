package root

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserService interface {
	CreateUser(u *User) error
	GetByUsername(username string) (*User, error)
}
