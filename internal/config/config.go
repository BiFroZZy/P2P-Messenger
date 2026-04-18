package config

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	l "messenger/internal/logger"
)

var logger = l.InitLogger()

type Config struct{
	ServerPort 		string 				`envconfig:"SERVER_PORT" required:"true"`
	ShutdownTimeout	time.Duration		`envcongig:"SHUTDOWN_TIMEOUT"`
	DatabaseURL 	string 				`envconfig:"DB_URL" required:"true"`
	JWTexp 			int64 				`envconfig:"JWT_EXP" required:"true"`
	// JWTiss 			string 				`envconfig:"JWT_ISS" required:"true"`
	JWTiat 			int64 				`envconfig:"JWT_IAT" required:"true"`
}

func Load() (*Config, error){
	var cfg Config
	if err := godotenv.Load(); err != nil{
		logger.Error().Err(err).Msg("Can't load env file")
	}
	if err := envconfig.Process("", &cfg); err != nil{
		logger.Error().Err(err).Msg("Error occured while getting configs")
		return nil, err
	}
	return &cfg, nil
}