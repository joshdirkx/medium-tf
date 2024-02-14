package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"medium_post":        resourceMediumPost(),
			"medium_profile":     resourceMediumProfile(),
			"medium_publication": resourceMediumPublication(),
		},
	}
}
