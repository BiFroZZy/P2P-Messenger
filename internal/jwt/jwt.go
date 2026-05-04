package jwt

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
    
    "messenger/internal/config"
	_ "messenger/internal/models"
)
var secretKey = []byte(os.Getenv("SECRET_KEY"))

func ParseJWT(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
}

func GenerateJWT(userID uuid.UUID, username string) (string, error){
    cfg, err := config.Load()
    if err != nil{
        return "Failed to loaf configs", err
    }
    // models.Validation(cfg)
    // models.Validation(user)
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     cfg.JWTexp, 
        "iss":     username,
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}