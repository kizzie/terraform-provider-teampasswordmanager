package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/kizzie/go-teampasswordmanager/teampasswordmanager"

	"errors"
	"log"
	"strconv"
)

// CustomField stores the label and data for each of the customfields in the
// password resource
type CustomField struct {
	Label string
	Data  string
}

func dataSourcePassword() *schema.Resource {
	return &schema.Resource{
		Read: resourcePasswordRead,

		Schema: map[string]*schema.Schema{
			"password_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_fields": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 10,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"data": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourcePasswordRead(d *schema.ResourceData, meta interface{}) error {
	id := d.Get("password_id").(string)
	name := d.Get("name").(string)
	project := d.Get("project").(string)

	if id != "" {
		log.Print("Getting password by ID: " + id)
		return getPasswordByID(id, d, meta)
	} else if name != "" && project != "" {
		log.Print("Getting password by Name and Project: " + name + "/" + project)
		return getPasswordByName(name, project, d, meta)
	} else {
		return errors.New("Did not find either ID or both Name and Project")
	}
}

func getPasswordByName(name string, project string, d *schema.ResourceData, meta interface{}) error {
	client := meta.(*teampasswordmanager.Client)
	obj, err := client.GetPasswordByName(name, project)
	if err != nil {
		log.Print(err)
		return err
	}

	return setupObject(strconv.Itoa(obj.ID), obj, d)
}

func getPasswordByID(id string, d *schema.ResourceData, meta interface{}) error {
	client := meta.(*teampasswordmanager.Client)
	idAsInt, _ := strconv.Atoi(id)
	obj, err := client.GetPassword(idAsInt)
	if err != nil {
		return err
	}
	return setupObject(id, obj, d)
}

func setupObject(id string, obj teampasswordmanager.Password, d *schema.ResourceData) error {
	d.SetId(id)
	d.Set("password_id", obj.ID)
	d.Set("name", obj.Name)
	d.Set("project", obj.Project.Name)
	d.Set("username", obj.Username)
	d.Set("password", obj.Password)

	objCustomFields := obj.CustomFields()
	customFields := []interface{}{}

	for i := 0; i < 10; i++ {
		field := map[string]interface{}{
			"label": objCustomFields[i].Label,
			"data":  objCustomFields[i].Data,
		}
		customFields = append(customFields, field)
	}

	// log.Print(custom_fields)
	d.Set("custom_fields", customFields)
	log.Print(d.Get("custom_fields"))

	// return nil if there are no issues
	return nil
}
