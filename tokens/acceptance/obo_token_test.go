package acceptance

import (
	"github.com/databricks/terraform-provider-databricks/internal/acceptance"
	"github.com/databricks/terraform-provider-databricks/qa"

	"testing"
)

func TestAccAwsOboTokenResource(t *testing.T) {
	qa.RequireCloudEnv(t, "aws")
	acceptance.Test(t, []acceptance.Step{
		{
			Template: `
			resource "databricks_service_principal" "this" {
				display_name = "tf-{var.RANDOM}"
			}

			data "databricks_group" "admins" {
				display_name = "admins"
			}
			
			resource "databricks_group_member" "this" {
				group_id = data.databricks_group.admins.id
				member_id = databricks_service_principal.this.id
			}

			resource "databricks_obo_token" "this" {
				depends_on = [databricks_group_member.this]
				application_id = databricks_service_principal.this.application_id
				comment = "PAT on behalf of ${databricks_service_principal.this.display_name}"
				lifetime_seconds = 3600
			}`,
		},
	})
}
