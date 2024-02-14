package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Data source for listing Medium publications
func resourceMediumPublication() *schema.Resource {
	return &schema.Resource{
		Create: resourceMediumPublicationCreate,
		Read:   resourceMediumPublicationRead,
		Update: resourceMediumPublicationUpdate,
		Delete: resourceMediumPublicationDelete,

		Schema: map[string]*schema.Schema{
			"access_token": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"publications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"image_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceMediumPublicationCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMediumPublicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	accessToken := d.Get("access_token").(string)
	userId := d.Get("user_id").(string)

	publications, err := ListPublications(accessToken, userId)
	if err != nil {
		return diag.FromErr(err)
	}

	// Convert publications to a format that Terraform can work with
	var tfPublications []interface{}
	for _, publication := range publications {
		pub := make(map[string]interface{})
		pub["id"] = publication.Id
		pub["name"] = publication.Name
		pub["description"] = publication.Description
		pub["url"] = publication.Url
		pub["image_url"] = publication.ImageUrl
		tfPublications = append(tfPublications, pub)
	}

	if err := d.Set("publications", tfPublications); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%s-publications", userId))

	return nil
}

func resourceMediumPublicationUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMediumPublicationDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
