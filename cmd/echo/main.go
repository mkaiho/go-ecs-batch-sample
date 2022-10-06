package main

import (
	"github.com/mkaiho/go-ecs-batch-sample/logging"
)

func init() {
	logging.InitLoggerWithZap()
}

func main() {
	logging.GetLogger().Info("not implemented")
}
