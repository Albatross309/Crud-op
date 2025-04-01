package userServices

import (
	"crud-app/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if handleGenericError(w, "Failed to read request body!", err) {
		return
	}
	var newUser User
	if err = json.Unmarshal(body, &newUser); handleGenericError(w, "Failed to unmarshal request body!", err) {
		return
	}
	db, err := database.DbConnectionn()
	if handleGenericError(w, "Failed to connect to the database!", err) {
		return
	}
	defer db.Close()
	statement, err := db.Prepare("INSERT INTO users(name,email)VALUES(?,?)")
	if handleGenericError(w, "Failed to create statement!", err) {
		return
	}
	defer statement.Close()
	result, err := statement.Exec(newUser.Name, newUser.Email)
	if handleGenericError(w, "Failed to execute statement!", err) {
		return
	}
	createdID, err := result.LastInsertId()
	if handleGenericError(w, "Filed to retrive last insert ID!", err) {
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Successfully created user %d", createdID)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.DbConnectionn()
	if handleGenericError(w, "Failed to connect to the database!", err) {
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")
	if handleGenericError(w, "Failed to retrieve users!", err) {
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); handleGenericError(w, "Failed to scan users!", err) {
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); handleGenericError(w, "Failed to convert users to JSON!", err) {
		return
	}
}

// GetUserByID retrieves a user from the database by ID.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if handleGenericError(w, "Failed to convert parameter to integer", err) {
		return
	}

	db, err := database.DbConnectionn()
	if handleGenericError(w, "Failed to connect to the database!", err) {
		return
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM users WHERE id = ?", ID)
	if handleGenericError(w, "Failed to retrieve user "+strconv.FormatUint(ID, 10), err) {
		return
	}
	defer row.Close()

	var user User
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); handleGenericError(w, "Failed to scan users!", err) {
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); handleGenericError(w, "Failed to convert user to JSON!", err) {
		return
	}
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 10, 20)
	if handleGenericError(w, "Failed to convert parameter to integer", err) {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if handleGenericError(w, "Failed to read request body!", err) {
		return
	}
	var updateUser User
	if err = json.Unmarshal(body, &updateUser); handleGenericError(w, "Failed to unmarshal request body!", err) {
		return
	}
	db, err := database.DbConnectionn()
	if handleGenericError(w, "Failed to coonect to the database!", err) {
		return
	}
	defer db.Close()
	satatement, err := db.Prepare("UPDATE users SET name=?,email=?WHERE id=?")
	if handleGenericError(w, "Failed to create statement!", err) {
		return
	}
	defer satatement.Close()
	if _, err := satatement.Exec(updateUser.Name, updateUser.Email, ID); handleGenericError(w, "Failed to execute statement!", err) {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseInt(params["id"], 10, 20)
	if handleGenericError(w, "Failed to convert parameter to integer", err) {
		return
	}
	db, err := database.DbConnectionn()
	if handleGenericError(w, "Failed to connect to the database!", err) {
		return
	}
	defer db.Close()
	statement, err := db.Prepare("DELETE FROM users WHERE id=?")
	if handleGenericError(w, "Failed to create statement!", err) {
		return
	}
	defer statement.Close()
	if _, err := statement.Exec(ID); handleGenericError(w, "Failed to execute statement!", err) {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func handleGenericError(w http.ResponseWriter, errorMessage string, err error) bool {
	if err != nil {
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return true
	}
	return false
}
