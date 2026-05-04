package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	repository"messenger/internal/repository"
)

func RegUser(c echo.Context) error{
	regUsername := c.FormValue("reg_username")
	regLogin := c.FormValue("reg_login")
	regPassword := c.FormValue("reg_password")
	
	repo, err := repository.NewConnection()
	if err != nil{
		return err
	}
	defer repo.Close()

	hashedPass, err := HashPassword(regPassword)
	if err != nil{
		logger.Error().Err(err).Msg("Failed to hash password")
		return c.Render(http.StatusUnauthorized, "reg.html", map[string]interface{}{
			"Title": "Registration",
			"Error": "Internal server error",
		})
	}
	resultStruct, err := repo.GetUser(regLogin)
	if err!=nil{
		return c.Render(http.StatusUnauthorized, "reg.html", map[string]interface{}{
			"Title": "Registration",
			"Error": "There is no such user",
		})
	}
	ok := bcrypt.CompareHashAndPassword([]byte(resultStruct.UserPassword), []byte(regPassword))
	if regLogin == resultStruct.UserLogin && ok == nil{
		return c.Render(http.StatusUnauthorized, "reg.html", map[string]interface{}{
			"Title": "Registration",
			"Error": "Password or login already exists",
		})
	}
	if err := repo.SetUser(regUsername, regLogin, string(hashedPass)); err != nil{
		logger.Error().Err(err).Msg("Failed to set user in register page!")
		return err
	}
	
	return c.Redirect(http.StatusSeeOther, "/users/messages")
}