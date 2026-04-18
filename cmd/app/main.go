package main

import (
	"github.com/joho/godotenv"

	"messenger/cmd/server"
	l "messenger/internal/logger"
	"messenger/internal/config"
	"messenger/internal/database"
)

func main(){
	logger := l.InitLogger()
	cfg, err := config.Load()
	if err!=nil{
		logger.Error().Err(err).Msg("Can't load configs")
	}
	if err := godotenv.Load(); err != nil{
		logger.Err(err).Msg("Error occured while getting .env file")
	}
	if err := database.NewMigrate("./migrations", cfg.DatabaseURL); err != nil{
		logger.Error().Err(err).Msg("Failed to run migration")	
	}
	server := server.New(cfg)
	server.Start()

	logger.Info().Msg("App is started!")
}
