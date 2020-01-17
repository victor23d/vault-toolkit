When I learn Vault, I made some toolkits for convenience by the way.

在学习Vault的时候顺手做了一些小工具方便自己使用。仅供参考！


Walk through:
```
export XX_PROJECT_ENV=XX_PROJECT/DEV
export VAULT_ADDR=http:// ... :8200
export VAULT_TOKEN= ...
```
1. Seed data into Vault

```
go run examples/seedData.go
```

2. Get all secret

```
go test -v ./api
```

You can see how to use the GetAllSecret() and WriteCredential() in the test file ./api/example_test.go

Use JWT Login to get token
```
export VAULT_TOKEN=""
export TOKEN_PATH="" # this is your jwt token path
```


TODO
Print all secrets
Use case: streaming/pipeline/search all secrets content like | grep

rsync download/upload/update/delete secrets reflect directory structure
Use case: secrets batch job

Dump/Restore to db
