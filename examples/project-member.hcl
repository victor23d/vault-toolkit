path "PROJECT/*" {
  capabilities = ["read", "list"]
}

path "PROJECT/data/DEV/*" {
  capabilities = ["create", "update", "read", "delete"]
}

path "PROJECT/metadata/DEV/" {
  capabilities = ["list"]
}

path "PROJECT/data/QA/*" {
  capabilities = ["create", "update", "read", "delete"]
}

path "PROJECT/metadata/QA/" {
  capabilities = ["list"]
}

path "PROJECT/data/STAGE/*" {
  capabilities = ["read"]
}

path "PROJECT/metadata/STAGE/" {
  capabilities = ["list"]
}

path "PROJECT/data/PROD/*" {
  capabilities = []
}

path "PROJECT/metadata/PROD/" {
  capabilities = ["list"]
}

