package fwprovider

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV1"
	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/terraform-providers/terraform-provider-datadog/datadog/internal/utils"
)

var (
	_ resource.ResourceWithConfigure = &IdpMetadataResource{}
)

func NewIdpMetadataResource() resource.Resource {
	return &IdpMetadataResource{}
}

type IdpMetadataResource struct {
	ApiV1 *datadogV1.OrganizationsApi
	ApiV2 *datadogV2.OrganizationsApi
	Auth  context.Context
}

type IdpMetadataModel struct {
	ID                 types.String `tfsdk:"id"`
	PublicId           types.String `tfsdk:"public_id"`
	FederationMetadata types.String `tfsdk:"idp_federation_metadata_xml"`
	SamlIdpEndpoint    types.String `tfsdk:"saml_idp_endpoint"`
	SamlLoginUrl       types.String `tfsdk:"saml_login_url"`
}

func (r *IdpMetadataResource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	providerData, ok := request.ProviderData.(*FrameworkProvider)
	if !ok {
		response.Diagnostics.AddError("Unexpected Resource Configure Type", "")
		return
	}

	r.ApiV1 = providerData.DatadogApiInstances.GetOrganizationsApiV1()
	r.ApiV2 = providerData.DatadogApiInstances.GetOrganizationsApiV2()
	r.Auth = providerData.Auth
}

func (r *IdpMetadataResource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "idp_metadata"
}

func (r *IdpMetadataResource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description: "Provides a Datadog Idp Metadata resource. This can be used to create and manage Datadog SAML metadata.",
		Attributes: map[string]schema.Attribute{
			"public_id": schema.StringAttribute{
				Description: "The `public_id` of the organization you are operating within.",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"idp_federation_metadata_xml": schema.StringAttribute{
				Description: "XXXXXXXXXXXXXXX",
				Required:    true,
				Sensitive:   true,
			},
			"saml_idp_endpoint": schema.StringAttribute{
				Description: "We're only able to read and see if the SAML metadata is uploaded. If SAML IdP endpoint has changed, it indicates that a change has been made to the Metadata file outside of Terraform.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"saml_login_url": schema.StringAttribute{
				Description: "We're only able to read and see if the SAML metadata is uploaded. If SAML login url has changed, it indicates that a change has been made to the Metadata file outside of Terraform.",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			// Resource ID
			"id": utils.ResourceIDAttribute(),
		},
	}
}

func (r *IdpMetadataResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan IdpMetadataModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	idpFile, err := r.contentToFilePointer(plan.FederationMetadata.ValueString())
	if err != nil {
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error uploading idp metadata"))
		return
	}
	_, err = r.ApiV2.UploadIdPMetadata(ctx, *datadogV2.NewUploadIdPMetadataOptionalParameters().WithIdpFile(idpFile))
	if err != nil {
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error uploading idp metadata"))
		return
	}

	resp, httpResponse, err := r.ApiV1.GetOrg(ctx, plan.PublicId.ValueString())
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error verifying idp metadata upload"))
			return
		}
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error getting organization"))
		return
	}

	r.updateState(&plan, &resp)
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, idpFile)
	plan.FederationMetadata = types.StringValue(buf.String())

	response.Diagnostics.Append(response.State.Set(ctx, &plan)...)
}

func (r *IdpMetadataResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	var state IdpMetadataModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	resp, httpResponse, err := r.ApiV1.GetOrg(ctx, state.PublicId.ValueString())
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error verifying idp metadata existence"))
			return
		}
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error getting organization"))
		return
	}

	r.updateState(&state, &resp)

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (r *IdpMetadataResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var state IdpMetadataModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	idpFile, err := r.contentToFilePointer(state.FederationMetadata.ValueString())
	if err != nil {
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error uploading idp metadata"))
		return
	}
	_, err = r.ApiV2.UploadIdPMetadata(ctx, *datadogV2.NewUploadIdPMetadataOptionalParameters().WithIdpFile(idpFile))
	if err != nil {
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error uploading idp metadata"))
		return
	}

	resp, httpResponse, err := r.ApiV1.GetOrg(ctx, state.PublicId.ValueString())
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error verifying idp metadata upload"))
			return
		}
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error getting organization"))
		return
	}

	r.updateState(&state, &resp)
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, idpFile)
	state.FederationMetadata = types.StringValue(buf.String())

	response.Diagnostics.Append(response.State.Set(ctx, &state)...)
}

func (r *IdpMetadataResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state IdpMetadataModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}
	emptyMetadata := strings.TrimSpace(`
<?xml version="1.0" encoding="utf-8"?>
<EntityDescriptor ID="deleted" xmlns="urn:oasis:names:tc:SAML:2.0:metadata">
	<IDPSSODescriptor protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
		<SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect" Location="https://deleted.idp/sso/redirect" />
		<SingleSignOnService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://deleted.idp/sso/post" />
		<KeyDescriptor use="signing">
			<ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
				<ds:X509Data>
					<ds:X509Certificate>
						DELETED
					</ds:X509Certificate>
				</ds:X509Data>
			</ds:KeyInfo>
		</KeyDescriptor>
	</IDPSSODescriptor>
	<SPSSODescriptor protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
		<AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://deleted.url/acs" index="0" isDefault="true" />
	</SPSSODescriptor>
	</EntityDescriptor>
`)
	content, err := r.contentToFilePointer(emptyMetadata)
	if err != nil {
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error generating delete idp content body"))
		return
	}

	_, err = r.ApiV2.UploadIdPMetadata(ctx, *datadogV2.NewUploadIdPMetadataOptionalParameters().WithIdpFile(content))
	if err != nil {
		response.Diagnostics.Append(utils.FrameworkErrorDiag(err, "error deleting idp metadata"))
		return
	}
}

// /* Nothing to import as it's not possible to read the uploaded SAML manifest content */
// func (r *IdpMetadataResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
// 	resource.ImportStatePassthroughID(ctx, frameworkPath.Root("id"), request, response)
// }

func (r *IdpMetadataResource) contentToFilePointer(content string) (*os.File, error) {
	// Create a temporary file
	file, err := os.CreateTemp("", "temp-file-content")
	if err != nil {
		return nil, err
	}
	defer os.Remove(file.Name()) // Remove the temporary file when done

	// Write the string content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return nil, err
	}

	// Seek back to the beginning of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (r *IdpMetadataResource) updateState(state *IdpMetadataModel, o *datadogV1.OrganizationResponse) {
	state.SamlIdpEndpoint = types.StringValue(o.Org.Settings.GetSamlIdpEndpoint())
	state.SamlLoginUrl = types.StringValue(o.Org.Settings.GetSamlLoginUrl())
}
