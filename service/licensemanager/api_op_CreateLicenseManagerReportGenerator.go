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

// Creates a report generator.
func (c *Client) CreateLicenseManagerReportGenerator(ctx context.Context, params *CreateLicenseManagerReportGeneratorInput, optFns ...func(*Options)) (*CreateLicenseManagerReportGeneratorOutput, error) {
	if params == nil {
		params = &CreateLicenseManagerReportGeneratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLicenseManagerReportGenerator", params, optFns, c.addOperationCreateLicenseManagerReportGeneratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLicenseManagerReportGeneratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLicenseManagerReportGeneratorInput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	//
	// This member is required.
	ClientToken *string

	// Defines the type of license configuration the report generator tracks.
	//
	// This member is required.
	ReportContext *types.ReportContext

	// Frequency by which reports are generated. Reports can be generated daily,
	// monthly, or weekly.
	//
	// This member is required.
	ReportFrequency *types.ReportFrequency

	// Name of the report generator.
	//
	// This member is required.
	ReportGeneratorName *string

	// Type of reports to generate. The following report types an be generated:
	//   - License configuration report - Reports the number and details of consumed
	//   licenses for a license configuration.
	//   - Resource report - Reports the tracked licenses and resource consumption for
	//   a license configuration.
	//
	// This member is required.
	Type []types.ReportType

	// Description of the report generator.
	Description *string

	// Tags to add to the report generator.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateLicenseManagerReportGeneratorOutput struct {

	// The Amazon Resource Name (ARN) of the new report generator.
	LicenseManagerReportGeneratorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLicenseManagerReportGeneratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLicenseManagerReportGenerator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLicenseManagerReportGenerator{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLicenseManagerReportGenerator"); err != nil {
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
	if err = addOpCreateLicenseManagerReportGeneratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLicenseManagerReportGenerator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLicenseManagerReportGenerator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLicenseManagerReportGenerator",
	}
}
