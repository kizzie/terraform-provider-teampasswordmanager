build:
	go build -o terraform-provider-tpm

terraform_init:
	terraform init ./terraform

plan: build terraform_init
	TF_LOG=INFO TPM_AUTH_TOKEN=a2F0OnBhc3N3b3Jk terraform plan ./terraform

apply: build terraform_init
	TF_LOG=INFO TPM_AUTH_TOKEN=a2F0OnBhc3N3b3Jk terraform apply ./terraform

update-dependencies:
	go get -u
