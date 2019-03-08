package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handler);
	http.HandleFunc("/users", users);
	//http.HandleFunc("/", handler);

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func users(w http.ResponseWriter, r *http.Request) {

	const listOfPerson = "[" +
		"{'id': '1'; 'name': 'Nick'; 'description': 'Best friend'}," +
		"{'id': '2'; 'name': 'Bob'; 'description': 'Some guy'}," +
		"{'id': '3'; 'name': 'Sam'; 'description': 'Coworker'}" +
		"]"

	w.Header().Set("Server", "A Go Web Server")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081");
	w.Header().Set("Access-Control-Allow-Methods", "POST,GET");
	w.Header().Set("Access-Control-Allow-Headers", "*");
	w.Header().Set("Access-Control-Allow-Credentials", "true");
	w.WriteHeader(200)
	w.Write([]byte(listOfPerson))
}
