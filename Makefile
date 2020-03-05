encrypt-secrets:
	gcloud kms encrypt --location global --keyring siyan-io --key auth-server --plaintext-file secrets.json --ciphertext-file secrets.json.encrypted

decrypt-secrets:
	gcloud kms decrypt --location global --keyring siyan-io --key auth-server --plaintext-file secrets.json --ciphertext-file secrets.json.encrypted
