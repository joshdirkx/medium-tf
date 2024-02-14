package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform/diag"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"net/http"
)

func resourceMediumPost() *schema.Resource {
	return &schema.Resource{
		Create: resourceMediumPostCreate,
		Read:   resourceMediumPostRead,
		Update: resourceMediumPostUpdate,
		Delete: resourceMediumPostDelete,

		Schema: map[string]*schema.Schema{
			"access_token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The access token for API authentication.",
			},
			"user_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The user ID of the Medium account.",
			},
			"title": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The title of the Medium post.",
			},
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The content of the Medium post.",
			},
			"content_format": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The format of the content (html or markdown).",
			},
			"publish_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "draft",
				Description: "The publish status of the Medium post (draft, public, or unlisted).",
			},
			// Add other fields like tags, canonical_url as needed
		},
	}
}

// Medium API Reference
// https://github.com/Medium/medium-api-docs
func resourceMediumPostCreate(d *schema.ResourceData, m interface{}) error {
	// Extract configuration from the schema
	accessToken := d.Get("access_token").(string)
	userId := d.Get("user_id").(string)
	title := d.Get("title").(string)
	content := d.Get("content").(string)
	contentFormat := d.Get("content_format").(string)
	publishStatus := d.Get("publish_status").(string)

	// Create the request payload
	postData := map[string]interface{}{
		"title":         title,
		"contentFormat": contentFormat,
		"content":       content,
		"publishStatus": publishStatus,
		// Add other fields as necessary
	}
	payloadBytes, err := json.Marshal(postData)
	if err != nil {
		return diag.FromErr(err)
	}

	// Make the API request
	url := fmt.Sprintf("https://api.medium.com/v1/users/%s/posts", userId)
	req, err := http.NewRequest("POST", url, bytes.NewReader(payloadBytes))
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return diag.Errorf("API call to Medium failed: %s", string(body))
	}

	// Ideally, parse the response to get the post ID and set it as the resource ID
	// For example purposes, we'll just generate a placeholder
	d.SetId("medium_post_id_placeholder")

	return nil
}

func resourceMediumPostRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMediumPostUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMediumPostDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
