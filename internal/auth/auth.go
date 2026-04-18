package auth

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	j "messenger/internal/jwt"
	l "messenger/internal/logger"
	repository"messenger/internal/repository"
)
var (
	logger = l.InitLogger()
	ctx = context.Background()
)

func HashPassword(password string) ([]byte, error){
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error().Err(err).Msg("Error occured while hashing the password")
		return nil, err
	}
	return hashedPass,nil
}

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
		cookie, err := c.Cookie("token")
		if err != nil{
			return c.Render(http.StatusOK, "auth.html", map[string]interface{}{
				"Title": "Authorization",
				"Error": "Your session is over, authorize again to continue",
			})
		}
		token, err := j.ParseJWT(cookie.Value)
		if err != nil || !token.Valid {
			logger.Error().Err(err).Msg("Error in parsing cookie value!")
			return c.Redirect(http.StatusUnauthorized, "/public/auth")
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
		}
		return next(c)
    }
}

func AuthUser(c echo.Context) error{
	authLogin := c.FormValue("auth_login")
	authPassword := c.FormValue("auth_password")

	repo, err := repository.NewConnection()
	if err != nil{
		return err
	}
	defer repo.Close()

	resultStruct, err := repo.GetUser(authLogin)
	if err != nil{
		return err
	}

	token, err := j.GenerateJWT(resultStruct.UserID)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, "Error occured while generating JWT")
	}
	c.SetCookie(&http.Cookie{
		Name: "token",
		Value: token,
		Path: "/",
		Secure: false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		MaxAge: 300,
	})
	ok := bcrypt.CompareHashAndPassword([]byte(resultStruct.UserPassword), []byte(authPassword))
	if authLogin == resultStruct.UserLogin && ok == nil{
		return c.Redirect(http.StatusSeeOther, "/users/messages")
	}else{
		return c.Render(http.StatusUnauthorized, "auth.html", map[string]interface{}{
			"Title": "Auth",
			"Error": "Wrong login or password",
		})
	}
}