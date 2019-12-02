package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider ...
func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"hiera_lookup": dataSourceHieraLookup(),
		},
	}
}

func dataSourceHieraLookup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHieraLookupRead,

		Schema: map[string]*schema.Schema{
			"workspace": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},
			"instance_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceHieraLookupRead(d *schema.ResourceData, meta interface{}) error {
	workspace, workspaceOk := d.GetOk("workspace")
	if !workspaceOk {
		return fmt.Errorf("workspace is needed")
	}
	hieraData, err := get(workspace.(string))
	if err != nil {
		return err
	}
	d.SetId(workspace.(string))
	d.Set("instance_name", hieraData["instance_name"])
	d.Set("instance_type", hieraData["instance_type"])
	return nil
}

func get(workspace string) (map[string]string, error) {

	// prefix := "http://localhost:8080"
	prefix := "https://hiera-tfmdd2vwoq-uc.a.run.app"
	url := fmt.Sprintf("%s/lookup/%s", prefix, workspace)
	log.Printf("hiera URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var hieraData map[string]string
	err = json.Unmarshal(bs, &hieraData)
	if err != nil {
		return nil, err
	}
	log.Printf("hiera data: %+v", hieraData)

	return hieraData, nil
}
