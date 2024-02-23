// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new domain association for an Amplify app. This action associates a
// custom domain with the Amplify app
func (c *Client) CreateDomainAssociation(ctx context.Context, params *CreateDomainAssociationInput, optFns ...func(*Options)) (*CreateDomainAssociationOutput, error) {
	if params == nil {
		params = &CreateDomainAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomainAssociation", params, optFns, c.addOperationCreateDomainAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the create domain association request.
type CreateDomainAssociationInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The domain name for the domain association.
	//
	// This member is required.
	DomainName *string

	// The setting for the subdomain.
	//
	// This member is required.
	SubDomainSettings []types.SubDomainSetting

	// Sets the branch patterns for automatic subdomain creation.
	AutoSubDomainCreationPatterns []string

	// The required AWS Identity and Access Management (IAM) service role for the
	// Amazon Resource Name (ARN) for automatically creating subdomains.
	AutoSubDomainIAMRole *string

	// The type of SSL/TLS certificate to use for your custom domain. If you don't
	// specify a certificate type, Amplify uses the default certificate that it
	// provisions and manages for you.
	CertificateSettings *types.CertificateSettings

	// Enables the automated creation of subdomains for branches.
	EnableAutoSubDomain *bool

	noSmithyDocumentSerde
}

// The result structure for the create domain association request.
type CreateDomainAssociationOutput struct {

	// Describes the structure of a domain association, which associates a custom
	// domain with an Amplify app.
	//
	// This member is required.
	DomainAssociation *types.DomainAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDomainAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDomainAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDomainAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDomainAssociation"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateDomainAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomainAssociation(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateDomainAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDomainAssociation",
	}
}
