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

	// asyncq config
	// TODO: complete me

	// toolchains config
	ToolchainsConfig string `env:"TOOLCHAINS_CONFIG"     envDefault:"./toolchains/config.yaml"`
	QUEUES           string `env:"QUEUES"     envDefault:"default"`
	QUEUES_TYPE      string `env:"S3_USERNAME" envDefault:"official"`

	// lakefs config
	LakefsS3Domain string `env:"LAKEFS_S3_DOMAIN" envDefault:"s3.lakefs.example.com"`
	LakefsHost     string `env:"LAKEFS_HOST"      envDefault:""`
	LakefsPort     int    `env:"LAKEFS_PORT"      envDefault:"34766"`
	LakefsUsername string `env:"LAKEFS_USERNAME"  envDefault:"lakefs"`
	LakefsPassword string `env:"LAKEFS_PASSWORD"  envDefault:"lakefs"`

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
