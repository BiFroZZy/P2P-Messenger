package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/go-playground/validator/v10"

	l "messenger/internal/logger"
)
var (
	logger = l.InitLogger()
)



type Result struct{
	UserID 			uuid.UUID
	UserLogin 		string
	UserPassword 	string
	UserEmail 		string
}

type Message struct{
	ID 			uuid.UUID 	`validate:"uuid,required"`
	CreatedAt 	time.Time 	`validate:"required"`
	Content 	string		`validate:"alphanum"`
	UserName 	string 		`validate:"min=2,max=50"`
}

type User struct {
	ID 			uuid.UUID 	`validate:"uuid,required"`
	Name 		string		`validate:"min=2,max=20,required"`
	Login 		string		`validate:"min=4,max=10,required"`
	Password 	string		`validate:"min=6,max=10,required"`
	Email 		string		`validate:"email,required"`
	CreatedAt 	time.Time	`validate:"required"`
}

func Validation(data interface{}) error{
	validate := validator.New()
	if err := validate.Struct(data);err != nil{
		logger.Error().Err(err).Msg("Error occured while validating struct")
		return err
	}
	return nil
}