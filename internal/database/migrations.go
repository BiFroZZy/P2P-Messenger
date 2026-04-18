package database

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"

	"messenger/internal/config"
	l "messenger/internal/logger"
)

var (
	logger = l.InitLogger()
	cfg config.Config
)
func NewMigrate(migrationsPath string) error{
	m, err := migrate.New("file://"+migrationsPath, cfg.DatabaseURL)
	if err != nil{
		logger.Error().Err(err).Msg("Error occured while creatting migration")
		return err
	}
	if err := m.Up(); err != nil && errors.Is(err, migrate.ErrNoChange){
		return err
	}
	srcErr, dbErr := m.Close()
	if srcErr!= nil{
		return err
	}
	if dbErr!= nil{
		return err
	}
	logger.Info().Msg("Migrations apllied succeessfully")
	return nil
}