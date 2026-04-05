package database

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"messenger/internal/logger"
	"messenger/internal/models"
	"messenger/internal/config"
)
var (
	log = logger.InitLogger()
	ctx = context.Background()
)

func InitConnection() (*pgx.Conn, error){
	cfg := config.Config{}
	conn, err := pgx.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Error().Err(err).Msg("Error occured while connecting to database")
		return nil, err
	}
	return conn, nil
}

func GetUser(getUserLogin string)(uuid.UUID, string, string, string){
	conn, err := InitConnection()
	if err != nil{
		log.Error().Err(err).Msg("Error occired while getting user info")
	}
	defer conn.Close(ctx)

	user := models.User{}
	err = conn.QueryRow(ctx, os.Getenv("GET_USER"), getUserLogin).Scan(&user.ID, &user.Login, &user.Password, &user.Email)
	if err != nil{
		if err == pgx.ErrNoRows{
			log.Error().Err(err).Msg("There is no such user")
		}
		log.Error().Err(err).Msg("Error occured while querying user")
	}

	return user.ID, user.Login, user.Password, user.Email
}
func SetUser(getUserLogin, getUserPassword string){
	conn, err := InitConnection()
	if err != nil{
		log.Error().Err(err).Msg("Can't connect to DB")
	}
	defer conn.Close(ctx)

	userUUID := uuid.New()
	_, err = conn.Exec(ctx, os.Getenv("SET_USER"), getUserLogin, getUserPassword, userUUID)
	if err != nil{
		log.Error().Err(err).Msg("Can't set user in DB")
	}
}