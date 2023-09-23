# Managing Secrets From Azure With External Secrets Operator (ESO)

TODO: Intro

## Setup

```bash
export RESOURCE_GROUP=$(\
    yq ".production.azure.resourceGroup" settings.yaml)

export KEYVAULT=cncf-demo-$(date +%Y%m%d%H%M%S)

yq --inplace ".production.azure.keyvault = \"$KEYVAULT\"" \
    settings.yaml

az keyvault create --name $KEYVAULT \
    --resource-group $RESOURCE_GROUP --location eastus \
    | tee azure-keyvault.json

az keyvault secret set --name production-postgresql \
    --vault-name $KEYVAULT \
    --value '{"password": "YouWillNeverFindOut"}'

# Navigate to the newly established Azure Key-Vault.
# Then select Role Assignment (IAM).
# Select the member aks-external-secrets-example-agentpool and
#   assign the role Key Vault Secret Officer to the member.
# You have now authorized access to the azure key-vault for the
#   managed identities in azure.

yq --inplace \
    ".spec.provider.azurekv.vaultUrl = \"https://$KEYVAULT.vault.azure.net/\"" \
    eso/secret-store-azure.yaml

cat eso/secret-store-azure.yaml

cp eso/secret-store-azure.yaml infra/.

git add .

git commit -m "Secret Store"

git push

kubectl --namespace external-secrets get clustersecretstores

# Wait until the ClusterSecretStore appears
```

## How Did You Define Your App?

* [Helm](helm.md)
* [Kustomize](kustomize.md)
* [Carvel ytt](carvel.md)
* [cdk8s](cdk8s.md)