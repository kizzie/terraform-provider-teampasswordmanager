provider "tpm" {
  url = "http://localhost/teampasswordmanager"
}

data "tpm_password" "password" {
  password_id = "1"
}

data "tpm_password" "password_by_name" {
  name    = "postgres"
  project = "stage.devops__foo--bar"
}

resource "local_file" "file" {
  content = <<EOS
Password by ID:
Name:     ${data.tpm_password.password.name},
Project:  ${data.tpm_password.password.project},
Password: ${data.tpm_password.password.password},
Custom Fields:
${data.tpm_password.password.custom_fields[0].label}: ${data.tpm_password.password.custom_fields[0].data}
${data.tpm_password.password.custom_fields[1].label}: ${data.tpm_password.password.custom_fields[1].data}

Password by Name:
Name:     ${data.tpm_password.password_by_name.name},
Project:  ${data.tpm_password.password_by_name.project},
Password: ${data.tpm_password.password_by_name.password},
Custom Fields:
${data.tpm_password.password_by_name.custom_fields[0].label}: ${data.tpm_password.password_by_name.custom_fields[0].data}
${data.tpm_password.password_by_name.custom_fields[1].label}: ${data.tpm_password.password_by_name.custom_fields[1].data}
EOS

  filename = "file.log"
}
