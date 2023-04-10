package generator

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRandomGenerator() *schema.Resource {
	return &schema.Resource{
		CreateContext: randomgeneratorCreate,
		ReadContext:   randomgeneratorRead,
		UpdateContext: randomgeneratorUpdate,
		DeleteContext: randomgeneratorDelete,

		Schema: map[string]*schema.Schema{
			"number": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func randomgeneratorCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	count := d.Get("number").(string)
	d.SetId(count)
	// https://www.uuidtools.com/api/generate/v1/count/uuid_count
	resp, err := http.Get("https://www.uuidtools.com/api/generate/v1/count/" + count)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERROR] %s: randomgeneratorCreate filed", err)
	}
	defer resp.Body.Close()
	log.Printf("[ERROR] %s: randomgeneratorCreate successfully", err)
	fmt.Println(string(body))
	return diags
}

func randomgeneratorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	return nil
}

func randomgeneratorUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return randomgeneratorRead(ctx, d, m)
}

func randomgeneratorDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	d.SetId("")
	return diags
}
