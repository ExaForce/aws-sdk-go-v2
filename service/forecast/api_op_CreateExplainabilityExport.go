// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exports an Explainability resource created by the CreateExplainability
// operation. Exported files are exported to an Amazon Simple Storage Service
// (Amazon S3) bucket. You must specify a DataDestination object that includes an
// Amazon S3 bucket and an Identity and Access Management (IAM) role that Amazon
// Forecast can assume to access the Amazon S3 bucket. For more information, see
// aws-forecast-iam-roles . The Status of the export job must be ACTIVE before you
// can access the export in your Amazon S3 bucket. To get the status, use the
// DescribeExplainabilityExport operation.
func (c *Client) CreateExplainabilityExport(ctx context.Context, params *CreateExplainabilityExportInput, optFns ...func(*Options)) (*CreateExplainabilityExportOutput, error) {
	if params == nil {
		params = &CreateExplainabilityExportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateExplainabilityExport", params, optFns, c.addOperationCreateExplainabilityExportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateExplainabilityExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateExplainabilityExportInput struct {

	// The destination for an export job. Provide an S3 path, an Identity and Access
	// Management (IAM) role that allows Amazon Forecast to access the location, and an
	// Key Management Service (KMS) key (optional).
	//
	// This member is required.
	Destination *types.DataDestination

	// The Amazon Resource Name (ARN) of the Explainability to export.
	//
	// This member is required.
	ExplainabilityArn *string

	// A unique name for the Explainability export.
	//
	// This member is required.
	ExplainabilityExportName *string

	// The format of the exported data, CSV or PARQUET.
	Format *string

	// Optional metadata to help you categorize and organize your resources. Each tag
	// consists of a key and an optional value, both of which you define. Tag keys and
	// values are case sensitive. The following restrictions apply to tags:
	//   - For each resource, each tag key must be unique and each tag key must have
	//   one value.
	//   - Maximum number of tags per resource: 50.
	//   - Maximum key length: 128 Unicode characters in UTF-8.
	//   - Maximum value length: 256 Unicode characters in UTF-8.
	//   - Accepted characters: all letters and numbers, spaces representable in
	//   UTF-8, and + - = . _ : / @. If your tagging schema is used across other services
	//   and resources, the character restrictions of those services also apply.
	//   - Key prefixes cannot include any upper or lowercase combination of aws: or
	//   AWS: . Values can have this prefix. If a tag value has aws as its prefix but
	//   the key does not, Forecast considers it to be a user tag and will count against
	//   the limit of 50 tags. Tags with only the key prefix of aws do not count
	//   against your tags per resource limit. You cannot edit or delete tag keys with
	//   this prefix.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateExplainabilityExportOutput struct {

	// The Amazon Resource Name (ARN) of the export.
	ExplainabilityExportArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateExplainabilityExportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateExplainabilityExport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateExplainabilityExport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateExplainabilityExport"); err != nil {
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
	if err = addOpCreateExplainabilityExportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateExplainabilityExport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateExplainabilityExport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateExplainabilityExport",
	}
}
