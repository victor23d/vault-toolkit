package api

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/vault/api"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strings"
)

// GetAllSecret returns the secrets in the PATH "PROJECT/ENV/*"
func GetAllSecret() map[string]interface{} {
	path := PROJECT + "/metadata/" + ENV
	return GetAllSecretFromPath(path)
}

func GetAllSecretFromPath(path string) map[string]interface{} {

	//Get dynamic token if VAULT_TOKEN not set
	vaultToken = getVaultToken(tokenURL)

	allSecrets := make(map[string]interface{})
	config := api.Config{Address: vaultAddr}
	client, err := api.NewClient(&config)
	if err != nil {
		log.Error(err)
	}
	client.SetToken(vaultToken)

	secrets, err := client.Logical().List(path)
	if err != nil {
		log.Error(err)
	}

	if secrets == nil {
		log.Error( "List secrets: ", secrets)
	}
	if secrets.Data == nil {
		log.Error( "List secrets: ", secrets)
		log.Error( "Can not retrieve secret list, secrets.Data: ", secrets.Data)
	}
	fmt.Printf("\nkeys list: %s\n", secrets.Data["keys"])

	// list keys loop
	for _, v := range secrets.Data["keys"].([]interface{}) {
		path = PROJECT + "/data/" + ENV + "/" + v.(string)
		secret, err := client.Logical().Read(path)
		if err != nil {
			log.Error(err)
		}
		if secrets == nil {
			log.Error( "secrets is nil: ", secret)
		}
		if secrets.Data == nil {
			log.Error( "secrets.Data is nil: ", secrets.Data)
		}
		//log.Info(secret.Data)
		// combine all json
		allSecrets[v.(string)] = secret.Data["data"]

		//NoPath separate all json, ignore the secret path
		//for k2,v2 := range secret.Data{
		//	secretMap[k2] = v2.(string)
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
					log.Error( "secret is nil", "key", k, "value", v)
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
				// /secrets/credential-test/credential-test.json
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
				// /secrets/PROJECT/DEV/credential-test.json
				dir := "/secrets/" + PROJECT_ENV + "/"
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


// GetAllSecretFromParent path can not contain '/'
// Example: secrets := GetAllSecretFromParent("PROJECT")
// secrets will have directory end with "/", otherwise you won't know if it is a json secret or a directory.
func GetAllSecretFromParent(path string) map[string]interface{} {
	if strings.Contains(path, "/") {
		// GetAllSecretFromParent path doesn't contain '/'
		log.Error( "GetAllSecretFromParent path can not contain '/', path=", path)

	}
	config := api.Config{Address: vaultAddr}
	client, err := api.NewClient(&config)
	if err != nil {
		log.Error(err)
	}
	client.SetToken(vaultToken)

	path = path + "/metadata/"

	//log.Info("path: ", path)

	secrets, err := client.Logical().List(path)
	if err != nil {
		log.Error(err)
	}

	if secrets == nil {
		log.Error( "List secrets: ", secrets)
	}
	if secrets.Data == nil {
		log.Error( "List secrets: ", secrets)
		log.Error( "Can not retrieve secret list, secrets.Data: ", secrets.Data)
	}
	//fmt.Printf("\nkeys list: %s\n", secrets.Data["keys"])

	parentSecretMap := make(map[string]interface{})
	for _,v := range secrets.Data["keys"].([]interface{}){
		secretMap := GetAllSecretFromPath(path + v.(string))
		parentSecretMap[v.(string)] = secretMap
	}
	parentSecretJson, err := json.MarshalIndent(parentSecretMap, "", "    ")
	if err != nil {
		log.Error(err)
	}
	fmt.Println(string(parentSecretJson))

	return parentSecretMap
}


// deprecated for v1
func getAllSecretNoPath() map[string]string {
	secretMap := make(map[string]string)
	config := api.Config{Address: vaultAddr}
	client, err := api.NewClient(&config)
	if err != nil {
		log.Error(err)
	}
	client.SetToken(vaultToken)
	secrets, err := client.Logical().List(PROJECT_ENV)
	if err != nil {
		log.Error(err)
	}
	if secrets == nil {
		log.Error( "List secrets: ", secrets)
		log.Error( "Can not retrieve secret list: ", client)
	}
	//log.Info(secrets.Data)

	// list keys loop
	for _, v := range secrets.Data["keys"].([]interface{}) {
		secret, err := client.Logical().Read(PROJECT_ENV + "/" + v.(string))
		if err != nil {
			log.Error(err)
		}
		if secret == nil {
			log.Error( "secret is nil", secret)
		}
		//log.Info(secret.Data)
		for k2, v2 := range secret.Data {
			secretMap[k2] = v2.(string)
		}
	}
	logSecretJsonNoPath(secretMap)
	return secretMap
}

// deprecated for v1
func logSecretJsonNoPath(secretMap map[string]string) {
	hideMap := make(map[string]string)
	for k, v := range secretMap {
		if len(v) == 0 {
			log.Error( "secret is nil", "key: ", k, "value: ", v)
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


// getVaultToken returns vaultToken if VAULT_TOKEN env set, or get token from tokenURL from http or disk
func getVaultToken(tokenURL string) string {
	vaultTokenEnv := os.Getenv("VAULT_TOKEN")
	if vaultTokenEnv != ""{
		return vaultTokenEnv
	}else {
		if strings.Contains("http", tokenURL) {
			resp, err := http.Get(tokenURL)
			if err != nil {
				log.Error( "Get token error: ", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				bodyBytes, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Error(err)
				}
				vaultToken = string(bodyBytes)
			} else {
				log.Error( "Can not retrieve token: ", resp.Status)
			}
			return vaultToken
		} else {
			_, err := os.Stat(tokenURL)
			if err != nil {
				log.Error( "vault token not found: ", err)
			}
			b, err := ioutil.ReadFile(tokenURL)
			vaultToken = string(b)
			return vaultToken
		}
	}
}
