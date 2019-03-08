package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Println("Application started")
	http.HandleFunc("/", handler);
	http.HandleFunc("/users", users);
	//http.HandleFunc("/", handler);

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func users(w http.ResponseWriter, r *http.Request) {

	var users []User

	user1 := User{"1", "nick", "best friend"}
	user2 := User{"2", "bob", "best coorker"}
	user3 := User{"3", "lois", "best partizanen"}
	users = append(users, user1)
	users = append(users, user2)
	users = append(users, user3)

	js, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Server", "A Go Web Server")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081");
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET");
	w.Header().Set("Access-Control-Allow-Headers", "*");
	w.Header().Set("Access-Control-Allow-Credentials", "true");
	w.WriteHeader(200)
	//json.NewEncoder(w).Encode(user)
	w.Write(js)
}
