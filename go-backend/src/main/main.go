package main

import (
	"./models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Env struct {
	db *sql.DB
}

func main() {
	log.Println("Application started")

	db, err := models.NewDB("root:root@tcp(localhost:3306)/godb")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}

	http.HandleFunc("/", handler);
	http.Handle("/users", getUsers(env));
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

func getUsers(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		users, err := models.AllUsers(env.db)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

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

		w.Write(js)
	})
}
