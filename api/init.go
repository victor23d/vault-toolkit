package api

import (
	"go.uber.org/zap"
	"os"
	"strings"
)

var (
	logger, _      = zap.NewProduction()
	log            = logger.Sugar()
	XX_PROJECT_ENV = os.Getenv("XX_PROJECT_ENV")
	XX_PROJECT     string
	XX_ENV         string
	vaultAddr      = os.Getenv("VAULT_ADDR")
	vaultToken     string
	tokenURL       string
	//vaultToken     = login()
	//vaultRedirectAddr = os.Getenv("VAULT_REDIRECT_ADDR")
)

// set environment variables
func init() {
	if XX_PROJECT_ENV == "" {
		log.Fatalw("xx_vault", "XX_PROJECT_ENV", "null")
	}
	if vaultAddr == "" {
		log.Fatalw("xx_vault", "VAULT_ADDR", "null")
	}
	log.Infow("xx_vault", "XX_PROJECT_ENV", XX_PROJECT_ENV)
	log.Infow("xx_vault", "VAULT_ADDR", vaultAddr)
	XX_PROJECT = strings.Split(XX_PROJECT_ENV, "/")[0]
	XX_ENV = strings.Split(XX_PROJECT_ENV, "/")[1]
	vaultToken = os.Getenv("VAULT_TOKEN")
	if vaultToken == ""{
		tokenURL = os.Getenv("TOKEN_URL")
		log.Infow("xx_vault", "TOKEN_URL", tokenURL)
	}else {
		log.Infow("xx_vault", "VAULT_TOKEN", vaultToken[:6]+"******")
	}
}
