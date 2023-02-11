package main

import (
	_ "github.com/joint-online-judge/go-tiger/pkg/logger"

	"github.com/joint-online-judge/go-tiger/pkg/configs"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Hi, %s!", configs.Conf.HorseUsername)
}
