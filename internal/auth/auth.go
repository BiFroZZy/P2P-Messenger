package auth

import (
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	j "messenger/internal/jwt"
	"messenger/internal/logger"
	"messenger/internal/models"
	"messenger/internal/repository"
)
var (
	log = logger.InitLogger()
	ctx = context.Background()
)
func HashPassword(password string) ([]byte, error){
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("Error occured while hashing the password")
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
			log.Error().Err(err).Msg("Error in parsing cookie value!")
			return c.Redirect(http.StatusUnauthorized, "/public/auth")
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
		}
		return next(c)
    }
}

func Auth(c echo.Context) error{
	conn, err := database.InitConnection()
	if err != nil{
		log.Err(err).Msg("Error occured while connecting to DB")
	}
	authUsername := c.FormValue("username")
	authPassword := c.FormValue("password")
	userID, _, _, _ := database.GetUser(authUsername)
	token, err := j.GenerateJWT(userID)
	if err != nil{
		return c.JSON(http.StatusInternalServerError, "Error occured while generating JWT")
	}
	c.SetCookie(&http.Cookie{
		Name: "token",
		Value: token,
		Path: "/",
		MaxAge: 300,
	})
	user := models.User{}
	err = conn.QueryRow(ctx, os.Getenv("AUTH_QUERY")).Scan(&user.Login, &user.Password)
	if err != nil{
		if err == pgx.ErrNoRows{
			log.Error().Err(err).Msg("No rows found!")
		}
		log.Error().Err(err).Msg("Error occured while querying user in auth")
	}
	hashPass, err := HashPassword(user.Password)
	ok := bcrypt.CompareHashAndPassword(hashPass, []byte(authPassword))
	if authUsername == user.Login && ok == nil{
		return c.Redirect(http.StatusSeeOther, "/messages")
	}
	database.SetUser(authUsername, authPassword)
	return c.Render(http.StatusUnauthorized, "auth.html", map[string]interface{}{
		"Title": "Auth",
		"Error": "Wromng login or password",
	})
}