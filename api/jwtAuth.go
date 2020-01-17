package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type jwtAuth struct {
	Jwt  string `json:"jwt"`
	Role string `json:"role"`
}

type auth struct {
	Auth map[string]interface{} `json:"auth"`
	//Client_token map[string]string `json:"client_token"`
}

// login() use jwt token to login Vault and return token. Set "TOKEN_PATH" env to use
func Login() string {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr, Timeout: time.Second * 10}

	requestURL := vaultAddr + "/v1/auth/" + "kubernetes" + "/login"
	log.Info(requestURL)

	if tokenURL == "" {
		tokenURL = "/var/run/secrets/kubernetes.io/serviceaccount/token"
	}
	jwtToken, err := ioutil.ReadFile(tokenURL)
	if err != nil {
		log.Error("set environment TOKEN_PATH")
		log.Error(err)
	}

	jwtAuth, err := json.Marshal(jwtAuth{
		Jwt:  string(jwtToken),
		Role: XX_PROJECT +  "-app",
	})
	log.Info(string(jwtAuth))
	req, err := http.NewRequest("PUT", requestURL, bytes.NewBuffer(jwtAuth))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Error(resp.Status)
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
	}

	//var loginAuth map[string]auth
	var loginAuth auth

	if err := json.Unmarshal(s, &loginAuth); err != nil {
		panic(err)
	}
	return loginAuth.Auth["client_token"].(string)
}
