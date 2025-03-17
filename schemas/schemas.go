package schemas

type User struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Role     bool   `json:"role"`
}
