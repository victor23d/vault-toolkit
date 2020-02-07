package api

import "testing"

func TestGetAllSecretFromParent(t *testing.T){

	secrets := GetAllSecretFromParent(PROJECT)
	if secrets[ENV+"/"] == nil{
		t.Error(secrets[ENV+"/"])
		t.Fatal(secrets)
	}
}
