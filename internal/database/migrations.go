package database

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"messenger/internal/config"
	l "messenger/internal/logger"
)

var (
	logger = l.InitLogger()
	cfg config.Config
)
func NewMigrate(path, dbURL string) error{
	m, err := migrate.New("file://"+path, dbURL)
	if err != nil{
		logger.Error().Err(err).Msg("Error occured while creating migration")
		return err
	}
	if err := m.Up(); err != nil && errors.Is(err, migrate.ErrNoChange){
		return err
	}
	srcErr, dbErr := m.Close()
	if srcErr!= nil{
		logger.Error().Err(err).Msg("Failed to close source")
		return err
	}
	if dbErr!= nil{
		logger.Error().Err(err).Msg("Failed to close database")
		return err
	}
	logger.Info().Msg("Migrations applied succeessfully")
	return nil
}