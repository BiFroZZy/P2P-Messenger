package handlers

import (
	"net/http"

	repo "messenger/internal/repository"

	"github.com/labstack/echo/v4"
)

func ShowAuth(c echo.Context) error{
	return c.Render(http.StatusOK, "auth.html", map[string]interface{}{
		"Title": "Authorization",
	})
}
func ShowReg(c echo.Context) error {
	return c.Render(http.StatusOK, "reg.html", map[string]interface{}{
		"Title": "Registration",
	})
}
func ShowMessages(c echo.Context) error{
	return c.Render(http.StatusOK, "messages.html", map[string]interface{}{
		"Title":"Messages",
		"UserID": c.Param("id"),
	})
}
func ShowProfile(c echo.Context) error{
	userID := c.Param("id")
	repo, err := repo.NewConnection()
	if err != nil{
		return err
	}
	defer repo.Close()
	
	resultStruct, err := repo.GetUserByID(userID)
	if err != nil{
		return err
	}
	return c.Render(http.StatusOK, "profile.html", map[string]interface{}{
		"Title":"Proflle",
		"UserID": userID,
		"Nickname": resultStruct.UserName,
      	"Email": resultStruct.UserLogin,
		"CreatedAt": resultStruct.CreatedAt.Format("2006-01-02 15:04:05"), // userCreatedAt.Format("2006-01-02 15:04:05")
	})
}
