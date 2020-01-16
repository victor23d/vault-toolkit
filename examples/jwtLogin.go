package main

import (
	vt "github.com/victor23d/vault-toolkit/api"
	"go.uber.org/zap"
)

var (
	logger, _      = zap.NewProduction()
	log            = logger.Sugar()
)

func main() {
	token := vt.Login()
	log.Info(token)
}
