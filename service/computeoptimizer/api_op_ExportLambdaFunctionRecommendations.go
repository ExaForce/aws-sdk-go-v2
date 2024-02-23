// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exports optimization recommendations for Lambda functions. Recommendations are
// exported in a comma-separated values (.csv) file, and its metadata in a
// JavaScript Object Notation (JSON) (.json) file, to an existing Amazon Simple
// Storage Service (Amazon S3) bucket that you specify. For more information, see
// Exporting Recommendations (https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html)
// in the Compute Optimizer User Guide. You can have only one Lambda function
// export job in progress per Amazon Web Services Region.
func (c *Client) ExportLambdaFunctionRecommendations(ctx context.Context, params *ExportLambdaFunctionRecommendationsInput, optFns ...func(*Options)) (*ExportLambdaFunctionRecommendationsOutput, error) {
	if params == nil {
		params = &ExportLambdaFunctionRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportLambdaFunctionRecommendations", params, optFns, c.addOperationExportLambdaFunctionRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportLambdaFunctionRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportLambdaFunctionRecommendationsInput struct {

	// Describes the destination Amazon Simple Storage Service (Amazon S3) bucket name
	// and key prefix for a recommendations export job. You must create the destination
	// Amazon S3 bucket for your recommendations export before you create the export
	// job. Compute Optimizer does not create the S3 bucket for you. After you create
	// the S3 bucket, ensure that it has the required permission policy to allow
	// Compute Optimizer to write the export file to it. If you plan to specify an
	// object prefix when you create the export job, you must include the object prefix
	// in the policy that you add to the S3 bucket. For more information, see Amazon
	// S3 Bucket Policy for Compute Optimizer (https://docs.aws.amazon.com/compute-optimizer/latest/ug/create-s3-bucket-policy-for-compute-optimizer.html)
	// in the Compute Optimizer User Guide.
	//
	// This member is required.
	S3DestinationConfig *types.S3DestinationConfig

	// The IDs of the Amazon Web Services accounts for which to export Lambda function
	// recommendations. If your account is the management account of an organization,
	// use this parameter to specify the member account for which you want to export
	// recommendations. This parameter cannot be specified together with the include
	// member accounts parameter. The parameters are mutually exclusive.
	// Recommendations for member accounts are not included in the export if this
	// parameter, or the include member accounts parameter, is omitted. You can specify
	// multiple account IDs per request.
	AccountIds []string

	// The recommendations data to include in the export file. For more information
	// about the fields that can be exported, see Exported files (https://docs.aws.amazon.com/compute-optimizer/latest/ug/exporting-recommendations.html#exported-files)
	// in the Compute Optimizer User Guide.
	FieldsToExport []types.ExportableLambdaFunctionField

	// The format of the export file. The only export file format currently supported
	// is Csv .
	FileFormat types.FileFormat

	// An array of objects to specify a filter that exports a more specific set of
	// Lambda function recommendations.
	Filters []types.LambdaFunctionRecommendationFilter

	// Indicates whether to include recommendations for resources in all member
	// accounts of the organization if your account is the management account of an
	// organization. The member accounts must also be opted in to Compute Optimizer,
	// and trusted access for Compute Optimizer must be enabled in the organization
	// account. For more information, see Compute Optimizer and Amazon Web Services
	// Organizations trusted access (https://docs.aws.amazon.com/compute-optimizer/latest/ug/security-iam.html#trusted-service-access)
	// in the Compute Optimizer User Guide. Recommendations for member accounts of the
	// organization are not included in the export file if this parameter is omitted.
	// This parameter cannot be specified together with the account IDs parameter. The
	// parameters are mutually exclusive. Recommendations for member accounts are not
	// included in the export if this parameter, or the account IDs parameter, is
	// omitted.
	IncludeMemberAccounts bool

	noSmithyDocumentSerde
}

type ExportLambdaFunctionRecommendationsOutput struct {

	// The identification number of the export job. Use the
	// DescribeRecommendationExportJobs action, and specify the job ID to view the
	// status of an export job.
	JobId *string

	// Describes the destination Amazon Simple Storage Service (Amazon S3) bucket name
	// and object keys of a recommendations export file, and its associated metadata
	// file.
	S3Destination *types.S3Destination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportLambdaFunctionRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpExportLambdaFunctionRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpExportLambdaFunctionRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ExportLambdaFunctionRecommendations"); err != nil {
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
	if err = addOpExportLambdaFunctionRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportLambdaFunctionRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportLambdaFunctionRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ExportLambdaFunctionRecommendations",
	}
}
