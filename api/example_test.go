package api

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/vault/api"
	"io/ioutil"
	"strings"
)

func init() {
	//_ = os.Setenv("VAULT_ADDR", "http://127.0.0.1:8200")
	//_ = os.Setenv("XX_PROJECT_ENV", "XX_PROJECT/DEV")
	//_ = os.Setenv("VAULT_REDIRECT_ADDR", "")
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	deleteSecrets(client)
	enableSecretEngine(client)
	addSecrets(client)
}

func addSecrets(client *api.Client) {
	nvProject := strings.Split(XX_PROJECT_ENV, "/")[0]
	nvEnv := strings.Split(XX_PROJECT_ENV, "/")[1]
	// kv secret engine /secret path
	// path := "secret/data/" + nvProject + "/" + nvEnv
	// /project as secret engine path
	path := nvProject + "/data/" + nvEnv
	//v1/zond/data/dev/test

	tokens_test := map[string]interface{}{
		"aaa": "111",
		"bbb": "222222",
	}
	pubsub_admin_test := map[string]interface{}{
		"type":           "service_account",
		"client_email":   "project_id@appspot.gserviceaccount.com",
		"private_key":    "-----BEGIN PRIVATE KEY-----\nMIIEloooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog\n-----END PRIVATE KEY-----\n",
		"private_key_id": "0000000000000000000000000000000000000000",
	}

	path1 := path + "/tokens"
	_, err := client.Logical().Write(path1, map[string]interface{}{"data": tokens_test})
	if err != nil {
		log.Fatal(err)
	}
	path2 := path + "/pubsub-admin.json"
	_, err = client.Logical().Write(path2, map[string]interface{}{"data": pubsub_admin_test})
	if err != nil {
		log.Fatal(err)
	}
}
func deleteSecrets(client *api.Client) {

	path := "sys/mounts/XX_PROJECT"
	_, err := client.Logical().Delete(path)
	if err != nil {
		log.Fatal(err)
	}
	//_, err = client.Logical().Delete(XX_PROJECT_ENV + "/token")
	//_, err = client.Logical().Delete(XX_PROJECT_ENV + "/pubsub-admin.json")
}

func ExampleGetAllSecret() {

	_ = GetAllSecret()
	// Output:
	// keys list: [pubsub-admin.json tokens]
	// {
	//     "pubsub-admin.json": {
	//         "client_email": "project_id@appspot.******",
	//         "private_key": "-----BEGIN PRIVATE K***** *****",
	//         "private_key_id": "00000000000000000000******",
	//         "type": "service_account******"
	//     },
	//     "tokens": {
	//         "aaa": "111******",
	//         "bbb": "222******"
	//     }
	// }

}

func enableSecretEngine(client *api.Client) {
	path := "sys/mounts/XX_PROJECT"

	payLoad := []byte(`{"path":"XX_PROJECT","type":"kv","config":{},"options":{"version":2},"generate_signing_key":true}`)
	var enableXX_PROJECT map[string]interface{}
	if err := json.Unmarshal(payLoad, &enableXX_PROJECT); err == nil {
		fmt.Errorf("%s", err)

	}

	_, err := client.Logical().Write(path, enableXX_PROJECT)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleWriteCredential(){
	secrets := GetAllSecret()
	path := WriteCredential(secrets)
	content,err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
	// Output:
	// keys list: [pubsub-admin.json tokens]
	// {
	//     "pubsub-admin.json": {
	//         "client_email": "project_id@appspot.******",
	//         "private_key": "-----BEGIN PRIVATE K***** *****",
	//         "private_key_id": "00000000000000000000******",
	//         "type": "service_account******"
	//     },
	//     "tokens": {
	//         "aaa": "111******",
	//         "bbb": "222******"
	//     }
	// }
	// {"client_email":"project_id@appspot.gserviceaccount.com","private_key":"-----BEGIN PRIVATE KEY-----\nMIIEloooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooog\n-----END PRIVATE KEY-----\n","private_key_id":"0000000000000000000000000000000000000000","type":"service_account"}
}

// deprecated for v1
func ExampleGetAllSecretNoPath() {
	// _ = getAllSecretNoPath()
	// Output:
	// {
	//     "aaa": "111******",
	//     "bbb": "222******",
	//     "client_email": "project_id@appspot.******",
	//     "private_key": "-----BEGIN PRIVATE K******",
	//     "private_key_id": "00000000000000000000******",
	//     "type": "service_account******"
	// }
}
