package api

import (
	"go.uber.org/zap"
	"os"
	"strings"
)

var (
	logger, _      = zap.NewProduction()
	log            = logger.Sugar()
	PROJECT_ENV = os.Getenv("PROJECT_ENV")
	PROJECT     string
	ENV         string
	vaultAddr      = os.Getenv("VAULT_ADDR")
	vaultToken     string
	tokenURL       string
	//vaultToken     = login()
	//vaultRedirectAddr = os.Getenv("VAULT_REDIRECT_ADDR")
)

// set environment variables
func init() {
	if PROJECT_ENV == "" {
		log.Fatal( "PROJECT_ENV: ", "null")
	}
	if vaultAddr == "" {
		log.Fatal( "VAULT_ADDR: ", "null")
	}
	log.Info( "PROJECT_ENV: ", PROJECT_ENV)
	log.Info( "VAULT_ADDR: ", vaultAddr)
	PROJECT = strings.Split(PROJECT_ENV, "/")[0]
	ENV = strings.Split(PROJECT_ENV, "/")[1]
	vaultToken = os.Getenv("VAULT_TOKEN")
	if vaultToken == ""{
		tokenURL = os.Getenv("TOKEN_URL")
		log.Info( "TOKEN_URL: ", tokenURL)
	}else {
		log.Info( "VAULT_TOKEN: ", vaultToken[:6]+"******")
	}
}
