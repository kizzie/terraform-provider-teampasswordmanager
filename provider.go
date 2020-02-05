package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/kizzie/go-teampasswordmanager/teampasswordmanager"
)

// Provider for tpm
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TPM_URL", nil),
				Description: "URL of the team password manager instance - for example: http://localhost/teampasswordmanager",
			},
			"authtoken": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TPM_AUTH_TOKEN", nil),
				Description: "Base 64 encoded auth token for accessing the tpm server",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"tpm_password": dataSourcePassword(),
		},

		ConfigureFunc: tpmClientConfigure,
	}
}

func tpmClientConfigure(d *schema.ResourceData) (interface{}, error) {
	url, _ := d.Get("url").(string)
	authtoken, _ := d.Get("authtoken").(string)

	config := teampasswordmanager.ClientConfig{
		BaseURL:   url,
		AuthToken: authtoken,
	}

	client, _ := teampasswordmanager.NewClient(&config)

	return &client, nil
}
