package main

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/vault/api"
	"go.uber.org/zap"
	vt "github.com/victor23d/vault-toolkit/api"
)


var (
	logger, _      = zap.NewProduction()
	log            = logger.Sugar()
)

func main() {
	//_ = os.Setenv("VAULT_ADDR", "http://127.0.0.1:8200")
	//_ = os.Setenv("PROJECT_ENV", "PROJECT/DEV")
	//_ = os.Setenv("VAULT_REDIRECT_ADDR", "")
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	deleteSecretEngine(client)
	enableSecretEngine(client)
	//addSecrets(client, Env)
	addEnvSecrets(client, []string{"DEV","QA","STAGE","PROD"})

	// test if exists
	_ = vt.GetAllSecret()
}

func addEnvSecrets(client *api.Client, envs []string ) {
	for _,v := range envs{
		addSecrets(client, v)
	}
}

func addSecrets(client *api.Client, env string ) {
	// kv secret engine /secret path
	// path := "secret/data/" + PROJECT + "/" + Env
	// /project as secret engine path
	path := vt.PROJECT + "/data/" + env
	//v1/project/data/dev/test

	tokens_test := map[string]interface{}{
		"aaa": "111",
		"bbb": "222222",
	}
	credential_test := map[string]interface{}{
		"type":           "service_account",
		"client_email":   "project_id@appspot.gserviceaccount.com",
		"private_key":    "-----BEGIN PRIVATE KEY-----\nMIIEloooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog\n-----END PRIVATE KEY-----\n",
		"private_key_id": "0000000000000000000000000000000000000000",
	}

	path1 := path + "/tokens-test"
	_, err := client.Logical().Write(path1, map[string]interface{}{"data": tokens_test})
	if err != nil {
		log.Fatal(err)
	}
	path2 := path + "/credential-test.json"
	_, err = client.Logical().Write(path2, map[string]interface{}{"data": credential_test})
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Add secrets to "+ env + " OK!")
}

func deleteSecretEngine(client *api.Client) {
	path := "sys/mounts/" + vt.PROJECT
	//path := "sys/mounts/" + vt.PROJECT_ENV
	_, err := client.Logical().Delete(path)
	if err != nil {
		log.Fatal(err)
	}
	//_, err = client.Logical().Delete(PROJECT_ENV + "/tokens-test")
	//_, err = client.Logical().Delete(PROJECT_ENV + "/credential-test.json")
	log.Info("Delete secret engine "+ path +" OK!")
}


func enableSecretEngine(client *api.Client) {
	path := "sys/mounts/" + vt.PROJECT
	//path := "sys/mounts/" + vt.PROJECT_ENV

	payLoad := []byte(`{"path":"PROJECT","type":"kv","config":{},"options":{"version":2},"generate_signing_key":true}`)
	var enablePROJECT map[string]interface{}
	if err := json.Unmarshal(payLoad, &enablePROJECT); err == nil {
		fmt.Errorf("%s", err)
	}

	_, err := client.Logical().Write(path, enablePROJECT)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Enable secret engine "+ path +" OK!")
}

