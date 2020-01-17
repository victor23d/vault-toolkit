path "XX_PROJECT/*" {
  capabilities = ["read", "list"]
}

path "XX_PROJECT/data/DEV/*" {
  capabilities = ["read"]
}

path "XX_PROJECT/metadata/DEV/" {
  capabilities = ["list"]
}

path "XX_PROJECT/data/QA/*" {
  capabilities = ["read"]
}

path "XX_PROJECT/metadata/QA/" {
  capabilities = ["list"]
}

path "XX_PROJECT/data/STAGE/*" {
  capabilities = ["read"]
}

path "XX_PROJECT/metadata/STAGE/" {
  capabilities = ["list"]
}

path "XX_PROJECT/data/PROD/*" {
  capabilities = ["read"]
}

path "XX_PROJECT/metadata/PROD/" {
  capabilities = ["list"]
}

