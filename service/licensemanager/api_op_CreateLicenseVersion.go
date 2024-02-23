// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new version of the specified license.
func (c *Client) CreateLicenseVersion(ctx context.Context, params *CreateLicenseVersionInput, optFns ...func(*Options)) (*CreateLicenseVersionOutput, error) {
	if params == nil {
		params = &CreateLicenseVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLicenseVersion", params, optFns, c.addOperationCreateLicenseVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLicenseVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLicenseVersionInput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	//
	// This member is required.
	ClientToken *string

	// Configuration for consumption of the license. Choose a provisional
	// configuration for workloads running with continuous connectivity. Choose a
	// borrow configuration for workloads with offline usage.
	//
	// This member is required.
	ConsumptionConfiguration *types.ConsumptionConfiguration

	// License entitlements.
	//
	// This member is required.
	Entitlements []types.Entitlement

	// Home Region of the license.
	//
	// This member is required.
	HomeRegion *string

	// License issuer.
	//
	// This member is required.
	Issuer *types.Issuer

	// Amazon Resource Name (ARN) of the license.
	//
	// This member is required.
	LicenseArn *string

	// License name.
	//
	// This member is required.
	LicenseName *string

	// Product name.
	//
	// This member is required.
	ProductName *string

	// License status.
	//
	// This member is required.
	Status types.LicenseStatus

	// Date and time range during which the license is valid, in ISO8601-UTC format.
	//
	// This member is required.
	Validity *types.DatetimeRange

	// Information about the license.
	LicenseMetadata []types.Metadata

	// Current version of the license.
	SourceVersion *string

	noSmithyDocumentSerde
}

type CreateLicenseVersionOutput struct {

	// License ARN.
	LicenseArn *string

	// License status.
	Status types.LicenseStatus

	// New version of the license.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLicenseVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLicenseVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLicenseVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLicenseVersion"); err != nil {
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
	if err = addOpCreateLicenseVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLicenseVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLicenseVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLicenseVersion",
	}
}
