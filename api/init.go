package api

import (
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"strings"
)

var (
	logger, _      = zap.NewProduction()
	log            = logger.Sugar()
	XX_PROJECT_ENV = os.Getenv("XX_PROJECT_ENV")
	XX_PROJECT	string
	XX_ENV	string
	vaultAddr      = os.Getenv("VAULT_ADDR")
	vaultToken     string
	//vaultToken     = login()
	//vaultRedirectAddr = os.Getenv("VAULT_REDIRECT_ADDR")
)

func init() {
	if XX_PROJECT_ENV == "" {
		log.Fatalw("xx_vault", "XX_PROJECT_ENV", "null")
	}
	XX_PROJECT = strings.Split(XX_PROJECT_ENV, "/")[0]
	XX_ENV = strings.Split(XX_PROJECT_ENV, "/")[1]
	log.Infow("xx_vault", "XX_PROJECT_ENV", XX_PROJECT_ENV)
	log.Infow("xx_vault", "VAULT_ADDR", vaultAddr)
	vaultToken = os.Getenv("VAULT_TOKEN")
	if vaultToken == "" {
		tokenPath := "/etc/vault/token"
		_, err := os.Stat(tokenPath)
		if err != nil {
			log.Fatal("xx_vault", "vault token not found", err)
		}
		b, err := ioutil.ReadFile(tokenPath)
		vaultToken = string(b)
	}
	log.Infow("xx_vault", "VAULT_TOKEN", vaultToken[:6]+"******")
}
