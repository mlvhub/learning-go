package root

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserService interface {
	Create(u *User) error
	GetByUsername(username string) (*User, error)
	Login(c Credentials) (error, User)
}
