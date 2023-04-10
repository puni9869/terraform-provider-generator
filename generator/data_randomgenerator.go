package generator

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceRandomGenerator() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"number": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	count := d.Get("number").(string)
	if err := d.Set("number", count); err != nil {
		fmt.Println("[DEBUG] error in data source")
		diag.Errorf("error in getting number. %s", err)
		return diags
	}

	fmt.Println("[DEBUG] outputting  the value")

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
