package api

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/vault/api"
	"io/ioutil"
	"os"
	"os/user"
	"runtime"
	"strings"
)

// GetAllSecret returns the secrets in the PATH "PROJECT/ENV/*"
func GetAllSecret() map[string]interface{} {
	allSecrets := make(map[string]interface{})
	config := api.Config{Address: vaultAddr}
	client, err := api.NewClient(&config)
	if err != nil {
		log.Error(err)
	}
	client.SetToken(vaultToken)

	xxProject := strings.Split(XX_PROJECT_ENV, "/")[0]
	xxENV := strings.Split(XX_PROJECT_ENV, "/")[1]
	path := xxProject + "/metadata/" + xxENV

	//log.Infow("xx_vault", "path", path)

	secrets, err := client.Logical().List(path)
	if err != nil {
		log.Error(err)
	}

	if secrets == nil {
		log.Errorw("xx_vault", "List secrets", secrets)
	}
	if secrets.Data == nil {
		log.Errorw("xx_vault", "List secrets", secrets)
		log.Errorw("xx_vault", "Can not retrieve secret list, secrets.Data", secrets.Data)
	}
	fmt.Printf("\nkeys list: %s\n", secrets.Data["keys"])

	// list keys loop
	for _, v := range secrets.Data["keys"].([]interface{}) {
		path = xxProject + "/data/" + xxENV + "/" + v.(string)
		//log.Infow("xx_vault", "path", path)
		secret, err := client.Logical().Read(path)
		if err != nil {
			log.Error(err)
		}
		if secrets == nil {
			log.Errorw("xx_vault", "secrets is nil", secret)
		}
		if secrets.Data == nil {
			log.Errorw("xx_vault", "secrets.Data is nil", secrets.Data)
		}
		//log.Info(secret.Data)
		// combine all json
		allSecrets[v.(string)] = secret.Data["data"]

		//NoPath separate all json, ignore the secret path
		//for k2,v2 := range secret.Data{
		//	nvSecret[k2] = v2.(string)
		//}
	}
	logAllSecret(allSecrets)
	return allSecrets
}

func logAllSecret(secretMap map[string]interface{}) {
	hideMap := make(map[string]map[string]string)
	for k, v := range secretMap {
		hideMap[k] = make(map[string]string)
		for k2, vv2 := range v.(map[string]interface{}) {
			switch v2 := vv2.(type) {
			case string:
				if len(v2) == 0 {
					log.Errorw("xx_vault", "secret is nil", "key", k, "value", v)
					continue
				}
				if len(v2) > 40 {
					hideMap[k][k2] = v2[:20] + "***** *****"
				} else if len(v2)%2 == 0 {
					hideMap[k][k2] = v2[:len(v2)/2] + "******"
				} else {
					hideMap[k][k2] = v2[:len(v2)-1/2] + "******"
				}
			}
		}
	}
	hideJson, err := json.MarshalIndent(hideMap, "", "    ")
	if err != nil {
		log.Error(err)
	}
	fmt.Println(string(hideJson))
}

// WriteCredential write credential which name end with ".json" to the returned path
func WriteCredential(secretMap map[string]interface{}) string {
	var path string
	for k, v := range secretMap {
		if strings.Contains(k, ".json") {
			dirName := k[0 : len(k)-5]
			content, err := json.Marshal(v)
			if err != nil {
				log.Error(err)
			}
			thisUser, _ := user.Current()
			if thisUser.Uid == "0" || runtime.GOOS == "linux" {
				// /secrets/pubsub-admin/pubsub-admin.json
				dir := "/secrets/" + dirName + "/"
				path = dir + k
				if err := os.MkdirAll(dir, 0755); err != nil {
					log.Error(err)
				}
				err = ioutil.WriteFile(path, content, 0644)
				if err != nil {
					log.Error(err)
				}
			} else {
				// /secrets/XX_PROJECT/DEV/pubsub-admin.json
				dir := "/secrets/" + XX_PROJECT_ENV + "/"
				path = dir + k
				if err := os.MkdirAll(dir, 0755); err != nil {
					log.Error(err)
				}
				err = ioutil.WriteFile(path, content, 0644)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}
	return path
}

// deprecated for v1
func getAllSecretNoPath() map[string]string {
	nvSecret := make(map[string]string)
	config := api.Config{Address: vaultAddr}
	client, err := api.NewClient(&config)
	if err != nil {
		log.Error(err)
	}
	client.SetToken(vaultToken)
	secrets, err := client.Logical().List(XX_PROJECT_ENV)
	if err != nil {
		log.Error(err)
	}
	if secrets == nil {
		log.Errorw("xx_vault", "List secrets", secrets)
		log.Errorw("xx_vault", "Can not retrieve secret list", client)
	}
	//log.Info(secrets.Data)

	// list keys loop
	for _, v := range secrets.Data["keys"].([]interface{}) {
		secret, err := client.Logical().Read(XX_PROJECT_ENV + "/" + v.(string))
		if err != nil {
			log.Error(err)
		}
		if secret == nil {
			log.Errorw("xx_vault", "secret is nil", secret)
		}
		//log.Info(secret.Data)
		for k2, v2 := range secret.Data {
			nvSecret[k2] = v2.(string)
		}
	}
	logSecretJsonNoPath(nvSecret)
	return nvSecret
}

// deprecated for v1
func logSecretJsonNoPath(secretMap map[string]string) {
	hideMap := make(map[string]string)
	for k, v := range secretMap {
		if len(v) == 0 {
			log.Errorw("xx_vault", "secret is nil", "key", k, "value", v)
			continue
		}
		if len(v) > 40 {
			hideMap[k] = v[:20] + "******"
		} else if len(v)%2 == 0 {
			hideMap[k] = v[:len(v)/2] + "******"
		} else {
			hideMap[k] = v[:len(v)-1/2] + "******"
		}
	}
	hideJson, err := json.MarshalIndent(hideMap, "", "    ")
	if err != nil {
		log.Error(err)
	}
	fmt.Println(string(hideJson))
}
