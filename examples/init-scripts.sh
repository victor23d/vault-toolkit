ROLE=xx_project-app



set -ex
vault policy write ${ROLE} ${ROLE}.hcl
vault kv put secret/myapp/config username='appuser'  password='suP3rsec(et!'  ttl='30s'

vault auth enable userpass || true
vault write auth/userpass/users/test-user  password=123  policies=${ROLE}

# vault login -method=userpass username=test-user password=training
vault kv get secret/myapp/config


VAULT_SA_NAME=$(kubectl get sa vault-auth -o jsonpath="{.secrets[*]['name']}")
SA_JWT_TOKEN=$(kubectl get secret $VAULT_SA_NAME -o jsonpath="{.data.token}" | base64 --decode; echo)
SA_CA_CRT=$(kubectl get secret $VAULT_SA_NAME -o jsonpath="{.data['ca\.crt']}" | base64 --decode; echo)

APISERVER=$(kubectl config view --minify | grep server | cut -f 2- -d ":" | tr -d " ")
K8S_HOST=$(echo $APISERVER | cut -f 3- -d "/")

vault auth enable kubernetes || true
vault write auth/kubernetes/config  token_reviewer_jwt="$SA_JWT_TOKEN"  kubernetes_host="https://$K8S_HOST"  kubernetes_ca_cert="$SA_CA_CRT"
vault write auth/kubernetes/role/${ROLE} bound_service_account_names=vault-auth bound_service_account_namespaces=default policies=${ROLE} ttl=24h


# ref: https://learn.hashicorp.com/vault/developer/vault-agent-k8s
vault auth tune -default-lease-ttl=1m -max-lease-ttl=1m kubernetes/
