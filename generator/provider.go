package generator

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"generator": resourceRandomGenerator(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"generator": dataSourceRandomGenerator(),
		},
	}
}
