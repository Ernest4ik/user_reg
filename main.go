package main

import (
	"encoding/json"
	"log"
	"net/http"
	"user_reg/crud"
	sh "user_reg/schemas"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user sh.User
	errjs := json.NewDecoder(r.Body).Decode(&user)

	if err := crud.AddUser(user); !err || errjs != nil {
		w.WriteHeader(400)
		return
	}
	data, _ := json.Marshal(user)
	w.WriteHeader(200)
	w.Write([]byte(data))
}

func main() {
	http.HandleFunc("/user", RegisterUser)

	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Fatal(err)
	}
}
