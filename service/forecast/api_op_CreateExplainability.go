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

// Explainability is only available for Forecasts and Predictors generated from an
// AutoPredictor ( CreateAutoPredictor ) Creates an Amazon Forecast Explainability.
// Explainability helps you better understand how the attributes in your datasets
// impact forecast. Amazon Forecast uses a metric called Impact scores to quantify
// the relative impact of each attribute and determine whether they increase or
// decrease forecast values. To enable Forecast Explainability, your predictor must
// include at least one of the following: related time series, item metadata, or
// additional datasets like Holidays and the Weather Index. CreateExplainability
// accepts either a Predictor ARN or Forecast ARN. To receive aggregated Impact
// scores for all time series and time points in your datasets, provide a Predictor
// ARN. To receive Impact scores for specific time series and time points, provide
// a Forecast ARN. CreateExplainability with a Predictor ARN You can only have one
// Explainability resource per predictor. If you already enabled ExplainPredictor
// in CreateAutoPredictor , that predictor already has an Explainability resource.
// The following parameters are required when providing a Predictor ARN:
//   - ExplainabilityName - A unique name for the Explainability.
//   - ResourceArn - The Arn of the predictor.
//   - TimePointGranularity - Must be set to “ALL”.
//   - TimeSeriesGranularity - Must be set to “ALL”.
//
// Do not specify a value for the following parameters:
//   - DataSource - Only valid when TimeSeriesGranularity is “SPECIFIC”.
//   - Schema - Only valid when TimeSeriesGranularity is “SPECIFIC”.
//   - StartDateTime - Only valid when TimePointGranularity is “SPECIFIC”.
//   - EndDateTime - Only valid when TimePointGranularity is “SPECIFIC”.
//
// CreateExplainability with a Forecast ARN You can specify a maximum of 50 time
// series and 500 time points. The following parameters are required when providing
// a Predictor ARN:
//   - ExplainabilityName - A unique name for the Explainability.
//   - ResourceArn - The Arn of the forecast.
//   - TimePointGranularity - Either “ALL” or “SPECIFIC”.
//   - TimeSeriesGranularity - Either “ALL” or “SPECIFIC”.
//
// If you set TimeSeriesGranularity to “SPECIFIC”, you must also provide the
// following:
//   - DataSource - The S3 location of the CSV file specifying your time series.
//   - Schema - The Schema defines the attributes and attribute types listed in the
//     Data Source.
//
// If you set TimePointGranularity to “SPECIFIC”, you must also provide the
// following:
//   - StartDateTime - The first timestamp in the range of time points.
//   - EndDateTime - The last timestamp in the range of time points.
func (c *Client) CreateExplainability(ctx context.Context, params *CreateExplainabilityInput, optFns ...func(*Options)) (*CreateExplainabilityOutput, error) {
	if params == nil {
		params = &CreateExplainabilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateExplainability", params, optFns, c.addOperationCreateExplainabilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateExplainabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateExplainabilityInput struct {

	// The configuration settings that define the granularity of time series and time
	// points for the Explainability.
	//
	// This member is required.
	ExplainabilityConfig *types.ExplainabilityConfig

	// A unique name for the Explainability.
	//
	// This member is required.
	ExplainabilityName *string

	// The Amazon Resource Name (ARN) of the Predictor or Forecast used to create the
	// Explainability.
	//
	// This member is required.
	ResourceArn *string

	// The source of your data, an Identity and Access Management (IAM) role that
	// allows Amazon Forecast to access the data and, optionally, an Key Management
	// Service (KMS) key.
	DataSource *types.DataSource

	// Create an Explainability visualization that is viewable within the Amazon Web
	// Services console.
	EnableVisualization *bool

	// If TimePointGranularity is set to SPECIFIC , define the last time point for the
	// Explainability. Use the following timestamp format: yyyy-MM-ddTHH:mm:ss
	// (example: 2015-01-01T20:00:00)
	EndDateTime *string

	// Defines the fields of a dataset.
	Schema *types.Schema

	// If TimePointGranularity is set to SPECIFIC , define the first point for the
	// Explainability. Use the following timestamp format: yyyy-MM-ddTHH:mm:ss
	// (example: 2015-01-01T20:00:00)
	StartDateTime *string

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

type CreateExplainabilityOutput struct {

	// The Amazon Resource Name (ARN) of the Explainability.
	ExplainabilityArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateExplainabilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateExplainability{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateExplainability{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateExplainability"); err != nil {
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
	if err = addOpCreateExplainabilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateExplainability(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateExplainability(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateExplainability",
	}
}
