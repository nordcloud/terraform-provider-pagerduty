package pagerduty

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nordcloud/go-pagerduty/pagerduty"
)

func dataSourcePagerDutyService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePagerDutyServiceRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourcePagerDutyServiceRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*pagerduty.Client)

	log.Printf("[INFO] Reading PagerDuty service")

	searchName := d.Get("name").(string)

	o := &pagerduty.ListServicesOptions{
		Query: searchName,
	}

	return resource.Retry(3*time.Minute, func() *resource.RetryError {
		resp, _, err := client.Services.List(o)
		if err != nil {
			errResp := handleNotFoundError(err, d)
			if errResp != nil {
				time.Sleep(10 * time.Second)
				return resource.RetryableError(errResp)
			}

			return nil
		}

		var found *pagerduty.Service

		for _, service := range resp.Services {
			if service.Name == searchName {
				found = service
				break
			}
		}

		if found == nil {
			return resource.NonRetryableError(
				fmt.Errorf("Unable to locate any service with the name: %s", searchName),
			)
		}

		d.SetId(found.ID)
		d.Set("name", found.Name)

		return nil
	})
}
