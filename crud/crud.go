package crud

import (
	db "user_reg/database"
	sh "user_reg/schemas"
)

func AddUser(user sh.User) bool {
	_, err := db.Conn.Exec("insert into users (username, age, email, role) values ($1, $2, $3, $4)", user.Username, user.Age, user.Email, user.Role)
	if err != nil {
		return false
	}

	return true
}
