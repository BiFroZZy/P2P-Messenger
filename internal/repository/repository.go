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
	conn 		*pgx.Conn
	ctx 		context.Context
	user  		models.User
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
	var user models.User
	return &UserRepository{
		conn: conn,
		ctx: ctx,
		user: user,
	}, nil
}

func (r *UserRepository) GetUser(getUserLogin string)(*models.Result, error){
	err := r.conn.QueryRow(r.ctx, os.Getenv("GET_USER"), getUserLogin).Scan(&r.user.ID, &r.user.Name, &r.user.Login, &r.user.Password)
	resultStruct := models.Result{
		UserID: 		r.user.ID,
		UserName:		r.user.Name,	
		UserLogin: 		r.user.Login,
		UserPassword: 	r.user.Password,
	}
	if err != nil{
		if err == pgx.ErrNoRows{
			logger.Error().Err(err).Msg("No such user")
			return &resultStruct, nil
		}
		logger.Error().Err(err).Msg("Error occured while querying user")
		return nil, err
	}
	return &resultStruct, nil
}

func (r *UserRepository) GetUserByID(getUserID string)(*models.Result, error){
	err := r.conn.QueryRow(r.ctx, os.Getenv("GET_USER_BY_ID"), getUserID).Scan(&r.user.Name, &r.user.Login, &r.user.Password)
	resultStruct := models.Result{
		UserName:		r.user.Name,	
		UserLogin: 		r.user.Login,
		UserPassword: 	r.user.Password,
	}
	if err != nil{
		if err == pgx.ErrNoRows{
			logger.Error().Err(err).Msg("No such user")
			return &resultStruct, nil
		}
		logger.Error().Err(err).Msg("Error occured while querying user")
		return nil, err
	}
	return &resultStruct, nil
}

func (r *UserRepository) SetUser(username, login, password string) error{
	id := uuid.New()
	_, err := r.conn.Exec(r.ctx, os.Getenv("SET_USER"), id, username, login, password)
	if err != nil{
		logger.Error().Err(err).Msg("Can't insert user data in database")
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(userID uuid.UUID) error{
	_, err := r.conn.Exec(r.ctx, os.Getenv("DELETE_USER"), userID)
	if err != nil{
		logger.Error().Err(err).Msg("Failed to delete user from database")
		return err
	}
	return nil
}