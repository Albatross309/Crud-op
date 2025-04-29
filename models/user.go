package models

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// func CreateUser(c echo.Context) error {
// 	var user User
// 	if err := c.Bind(&user); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
// 	}
// if err:=database.DB.Create(&user).Error;err!=nil {
// 	fmt.Println("DB error:",err)
// 	return c.JSON(http.StatusInternalServerError,map[string]string{"error":"Failed to create user"})
// }
// 	return c.JSON(http.StatusCreated, user)
// }
