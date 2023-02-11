package main

import (
	_ "github.com/joint-online-judge/go-tiger/pkg/logger"

	"github.com/joint-online-judge/go-tiger/pkg/config"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Hi, %s!", config.Conf.HorseUsername)
}
