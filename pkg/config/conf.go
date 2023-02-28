package config

import (
	"github.com/caarlos0/env/v7"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type conf struct {
	// server config
	Prod    bool `env:"PROD"                envDefault:"false"`
	Debug   bool `env:"DEBUG"               envDefault:"false"`
	Workers int  `env:"WORKERS"             envDefault:"1"`

	// horse config
	HorseUsername string `env:"HORSE_USERNAME" envDefault:""`
	HorsePassword string `env:"HORSE_PASSWORD" envDefault:""`

	// redis config
	RedisHost     string `env:"REDIS_HOST"     envDefault:"localhost"`
	RedisPort     int    `env:"REDIS_PORT"     envDefault:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisDbIndex  int    `env:"REDIS_DB_INDEX" envDefault:"0"`

	// toolchains config
	ToolchainsConfig string `env:"TOOLCHAINS_CONFIG" envDefault:"./toolchains/config.yaml"`
	Queues           string `env:"QUEUES"            envDefault:"default"`
	QueuesType       string `env:"QUEUES_TYPE"       envDefault:"official"`

	// rollbar
	RollbarAccessToken string `env:"ROLLBAR_ACCESS_TOKEN" envDefault:""`
}

var Conf *conf

func init() {
	cfg := conf{}
	if err := env.Parse(&cfg); err != nil {
		logrus.Fatalf("failed to parse config: %+v", err)
	}
	logrus.Infof("config object: %+v", cfg)
	Conf = &cfg
}
