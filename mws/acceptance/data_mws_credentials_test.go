package acceptance

import (
	"fmt"
	"os"
	"testing"

	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataSourceMwsCredentials(t *testing.T) {
	cloudEnv := os.Getenv("CLOUD_ENV")
	if cloudEnv != "MWS" {
		t.Skip("Cannot run test on non-MWS environment")
	}
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			data "databricks_mws_credentials" "this" {
			}`,
			Check: func(s *terraform.State) error {
				r, ok := s.RootModule().Resources["data.databricks_mws_credentials.this"]
				if !ok {
					return fmt.Errorf("data not found in state")
				}
				ids := r.Primary.Attributes["ids.%"]
				if ids == "" {
					return fmt.Errorf("ids is empty: %v", r.Primary.Attributes)
				}
				return nil
			},
		},
	})
}
