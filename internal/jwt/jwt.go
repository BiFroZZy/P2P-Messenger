package jwt

import (
	"time"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

    _"messenger/internal/models"
)
var secretKey = []byte(os.Getenv("SECRET_KEY"))

func ParseJWT(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
}

func GenerateJWT(userID uuid.UUID) (string, error){
    // claims := models.JWT{
    //     UserID: userID,
    //     Exp: ,
    // }
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Second * 300).Unix(), 
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}