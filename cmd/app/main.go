package main

import (
	"github.com/joho/godotenv"

	"messenger/cmd/server"
	l "messenger/internal/logger"
	"messenger/internal/config"
)

func main(){
	log := l.InitLogger()
	cfg, err := config.Load()
	if err!=nil{
		log.Error().Err(err).Msg("Can't load configs")
	}
	if err := godotenv.Load(); err != nil{
		log.Err(err).Msg("Error occured while getting .env file")
	}
	server := server.New(cfg)
	server.Start()

	log.Info().Msg("App is started!")
}
