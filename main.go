package main

import (
	"fmt"

	_ "github.com/joint-online-judge/go-tiger/pkg/logger"

	"github.com/hibiken/asynq"
	"github.com/joint-online-judge/go-horse/pkg/task"
	"github.com/joint-online-judge/go-tiger/pkg/config"
	"github.com/sirupsen/logrus"
)

func main() {
	conf := config.Conf
	connOpt := asynq.RedisClientOpt{
		Addr:     fmt.Sprintf("%s:%d", conf.RedisHost, conf.RedisPort),
		Password: conf.RedisPassword,
		DB:       conf.RedisDbIndex,
	}
	srv := asynq.NewServer(connOpt, asynq.Config{})
	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeSubmitRecord, SubmitRecordTask)
	if err := srv.Run(mux); err != nil {
		logrus.Fatalf("could not run server: %v", err)
	}
}
