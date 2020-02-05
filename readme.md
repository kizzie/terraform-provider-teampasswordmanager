# Team Password Provider for terraform

Provides a data type for passwords from team password manager.

# To use:

See the example in the terraform folder.

The provider can either be fully configured or can be partially configured
with the rest of the config passed in from the command line:

```
provider "tpm" {
  url = "http://localhost/teampasswordmanager"
}
```

or
```
provider "tpm" {
  url = "http://localhost/teampasswordmanager"
  auth_token = "abcde"
}
```
The auth token is the base 64 encoded username and password. If team password
manager is configured to accept other auth tokens then these can be passed in instead.

To use then either give it the ID of the password or the project and name of the password to get:

```
data "tpm_password" "password" {
  password_id = "1"
}

data "tpm_password" "password_by_name" {
  name    = "postgres"
  project = "stage.devops__foo--bar"
}

```

These can be used in other parts of terraform as data resources:

```
Name:     ${data.tpm_password.password.name},
Project:  ${data.tpm_password.password.project},
Password: ${data.tpm_password.password.password},
Custom Fields:
${data.tpm_password.password.custom_fields[0].label}: ${data.tpm_password.password.custom_fields[0].data}
${data.tpm_password.password.custom_fields[1].label}: ${data.tpm_password.password.custom_fields[1].data}

```

# Troubleshooting

Run terraform with either debug or info logs showing:

```
TF_LOG=INFO terraform apply ./terraform
```

This will show the logs from the underlying library as well as the current provider
