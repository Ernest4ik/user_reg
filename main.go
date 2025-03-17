package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"user_reg/crud"
	sh "user_reg/schemas"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user sh.User
	age, _ := strconv.Atoi(r.FormValue("age"))

	user.Age = age
	user.Email = r.FormValue("email")
	user.Username = r.FormValue("username")
	err := crud.AddUser(user)
	if !err {
		w.WriteHeader(404)
		w.Write([]byte("ошибка запроса в бд"))
		return
	}
	data, _ := json.Marshal(user)
	w.WriteHeader(200)
	w.Write([]byte(data))
}

func GetForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "form.html")
}

func main() {
	http.HandleFunc("/", GetForm)
	http.HandleFunc("/user", RegisterUser)

	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Fatal(err)
	}
}
