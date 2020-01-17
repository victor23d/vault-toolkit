package api

import (
	"fmt"
	"io/ioutil"
)

func ExampleGetAllSecret() {

	_ = GetAllSecret()
	// Output:
	// keys list: [credential-test.json tokens-test]
	// {
	//     "credential-test.json": {
	//         "client_email": "project_id@appspot.******",
	//         "private_key": "-----BEGIN PRIVATE K***** *****",
	//         "private_key_id": "00000000000000000000******",
	//         "type": "service_account******"
	//     },
	//     "tokens-test": {
	//         "aaa": "111******",
	//         "bbb": "222******"
	//     }
	// }

}

func ExampleGetAllSecretFromParent() {

	GetAllSecretFromParent("XX_PROJECT")
	// Output:
	// keys list: [credential-test.json tokens-test]
	// {
	//     "credential-test.json": {
	//         "client_email": "project_id@appspot.******",
	//         "private_key": "-----BEGIN PRIVATE K***** *****",
	//         "private_key_id": "00000000000000000000******",
	//         "type": "service_account******"
	//     },
	//     "tokens-test": {
	//         "aaa": "111******",
	//         "bbb": "222******"
	//     }
	// }

}

func ExampleWriteCredential() {
	secrets := GetAllSecret()
	path := WriteCredential(secrets)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
	// Output:
	// keys list: [credential-test.json tokens-test]
	// {
	//     "credential-test.json": {
	//         "client_email": "project_id@appspot.******",
	//         "private_key": "-----BEGIN PRIVATE K***** *****",
	//         "private_key_id": "00000000000000000000******",
	//         "type": "service_account******"
	//     },
	//     "tokens-test": {
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
