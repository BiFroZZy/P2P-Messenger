package database

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	l "messenger/internal/logger"
	"messenger/internal/models"
	"messenger/internal/config"
)
var (
	logger = l.InitLogger()
	ctx = context.Background()
)

type UserRepository struct{
	conn *pgx.Conn
	ctx context.Context
}

type Result struct{
	UserID 			uuid.UUID
	UserLogin 		string
	UserPassword 	string
	UserEmail 		string
}

func (r *UserRepository) Close() error{
	if r.conn != nil{
		return r.conn.Close(r.ctx)
	}
	return nil
}

func NewConnection() (*UserRepository, error){
	cfg, err := config.Load()
	if err != nil{
		logger.Error().Err(err).Msg("Failed to get configs while creating connection to database")
		return nil, err
	}
	conn, err := pgx.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		logger.Error().Err(err).Msg("Error occured while connecting to database")
		return nil, err
	}
	return &UserRepository{
		conn: conn,
		ctx: ctx,
	}, nil
}

func (r *UserRepository) GetUser(getUserLogin string)(*Result, error){
	var user models.User
	err := r.conn.QueryRow(ctx, os.Getenv("GET_USER"), getUserLogin).Scan(&user.ID, &user.Login, &user.Password, &user.Email)
	if err != nil{
		if err == pgx.ErrNoRows{
			logger.Error().Err(err).Msg("No such user")
		}
		logger.Error().Err(err).Msg("Error occured while querying user")
	}
	resultStruct := Result{
		UserID: user.ID,
		UserLogin: user.Login,
		UserPassword: user.Password,
		UserEmail: user.Email,
	}
	return &resultStruct, nil
}

func (r *UserRepository) SetUser(username, login, password, email string) error{
	id := uuid.New()
	_, err := r.conn.Exec(ctx, os.Getenv("SET_USER"), id, username, login, password, email)
	if err != nil{
		logger.Error().Err(err).Msg("Can't insert user data in database")
		return err
	}
	return nil
}