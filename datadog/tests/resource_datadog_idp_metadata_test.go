package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
)

func TestAccDatadogIdpMetadataResource_main(t *testing.T) {
	// t.Parallel()
	// if isRecording() || isReplaying() {
	// 	t.Skip("This test doesn't support recording or replaying")
	// }
	// _, _, accProviders := testAccFrameworkMuxProviders(context.Background(), t)

	// resource.Test(t, resource.TestCase{
	// 	PreCheck:                 func() { testAccPreCheck(t) },
	// 	ProtoV5ProviderFactories: accProviders,
	// 	// CheckDestroy:             testAccCheckDatadogApiKeyDestroy(providers.frameworkProvider),
	// 	Steps: []resource.TestStep{
	// 		{
	// 			Config: testAccResourceIdpMetadataConfig(),
	// 			Check: resource.ComposeAggregateTestCheckFunc(
	// 				resource.TestCheckResourceAttr("datadog_idp_metadata.this", "name", "Gruntwork AS"),
	// 			),
	// 		},
	// 	},
	// })
	os.Setenv("DD_SITE", "datadoghq.eu")
	ctx := datadog.NewDefaultContext(context.Background())
	configuration := datadog.NewConfiguration()
	apiClient := datadog.NewAPIClient(configuration)
	api := datadogV2.NewOrganizationsApi(apiClient)
	// r, err := api.UploadIdPMetadata(ctx, *datadogV2.NewUploadIdPMetadataOptionalParameters().WithIdpFile(func() *os.File {
	// 	fp, _ := os.Open("fixtures/organizations/saml_configurations/valid_idp_metadata.xml")
	// 	return fp
	// }()))
	r, err := api.UploadIdPMetadata(ctx, *datadogV2.NewUploadIdPMetadataOptionalParameters().WithIdpFile(getEmptyFile()))
	// r, err := api.UploadIdPMetadata(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.UploadIdPMetadata`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

}

func getEmptyFile() *os.File {
	file, err := os.CreateTemp("", "empty_file.xml")
	if err != nil {
		return nil
	}
	defer file.Close()

	return file
}

// func testAccResourceIdpMetadataConfig() string {
// 	return `
// provider "datadog" {
// 	api_url = "https://api.datadoghq.eu/"
// }

// resource "datadog_idp_metadata" "this" {
// 	name = "Gruntwork AS"

// 	idp_federation_metadata_xml = "{}"
// }`
// }
