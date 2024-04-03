package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/flovouin/terraform-provider-metabase/metabase"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var providerConfig = fmt.Sprintf(`
provider "metabase" {
  endpoint = "%s"
  username = "%s"
  password = "%s"
}
`,
	os.Getenv("METABASE_URL"),
	os.Getenv("METABASE_USERNAME"),
	os.Getenv("METABASE_PASSWORD"),
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"metabase": providerserver.NewProtocol6WithError(New("test")()),
}

var testAccMetabaseClient, _ = metabase.MakeAuthenticatedClientWithUsernameAndPassword(
	context.Background(),
	os.Getenv("METABASE_URL"),
	os.Getenv("METABASE_USERNAME"),
	os.Getenv("METABASE_PASSWORD"),
)
