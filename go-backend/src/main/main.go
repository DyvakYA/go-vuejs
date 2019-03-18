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

// main method for starting application
func main() {
	log.Println("Application started")

	// connect to db
	db, err := models.NewDB("root:root@tcp(localhost:3306)/godb")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}

	// request handlers
	http.HandleFunc("/", handler);
	http.Handle("/users", getUsers(env));

	// request address port
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}

// get user from db
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

		// make json from users
		js, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// make answer
		w.Header().Set("Server", "A Go Web Server")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081");
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET");
		w.Header().Set("Access-Control-Allow-Headers", "*");
		w.Header().Set("Access-Control-Allow-Credentials", "true");
		w.WriteHeader(200)

		// answer
		w.Write(js)
	})
}
