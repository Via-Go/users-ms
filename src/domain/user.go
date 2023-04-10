package domain

type Role int

const (
	Administrator Role = iota
	Moderator
	Common
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     int    `json:"role"`
	Id       string `json:"id"`
}
