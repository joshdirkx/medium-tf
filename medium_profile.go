package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Data source for getting a Medium user profile
func resourceMediumProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceMediumProfileCreate,
		Read:   resourceMediumProfileRead,
		Update: resourceMediumProfileUpdate,
		Delete: resourceMediumProfileDelete,

		Schema: map[string]*schema.Schema{
			"access_token": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
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
	}
}

func resourceMediumProfileCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMediumProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	accessToken := d.Get("access_token").(string)

	userProfile, err := GetUserProfile(accessToken)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(userProfile.Data.Id)
	d.Set("username", userProfile.Data.Username)
	d.Set("name", userProfile.Data.Name)
	d.Set("url", userProfile.Data.Url)
	d.Set("image_url", userProfile.Data.ImageUrl)

	return nil
}

func resourceMediumProfileUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMediumProfileDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
