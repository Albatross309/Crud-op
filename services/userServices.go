package userServices

import (
	"crud-app/database"
	"crud-app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	_, err := database.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", u.Name, u.Email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}

// Get all users
func GetUsers(c echo.Context) error {
	rows, err := database.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return err
		}
		users = append(users, u)
	}

	return c.JSON(http.StatusOK, users)
}

// Get user by ID
func GetUser(c echo.Context) error {
	id := c.Param("id")
	var u models.User
	err := database.DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

// Update user
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	_, err := database.DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", u.Name, u.Email, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

// Delete user
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
}
