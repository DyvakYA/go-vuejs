package models

import "database/sql"

// entity User
type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// get all users from db
func AllUsers(db *sql.DB) ([]*User, error) {
	rows, err := db.Query("SELECT * FROM users")

	// catch exceptions
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// make users from rows
	users := make([]*User, 0)
	for rows.Next() {
		user := new(User)
		err := rows.Scan(&user.Id, &user.Name, &user.Description)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// catch exceptions
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// return result or null
	return users, nil
}
