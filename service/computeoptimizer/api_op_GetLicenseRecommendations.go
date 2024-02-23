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

// Returns license recommendations for Amazon EC2 instances that run on a specific
// license. Compute Optimizer generates recommendations for licenses that meet a
// specific set of requirements. For more information, see the Supported resources
// and requirements (https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html)
// in the Compute Optimizer User Guide.
func (c *Client) GetLicenseRecommendations(ctx context.Context, params *GetLicenseRecommendationsInput, optFns ...func(*Options)) (*GetLicenseRecommendationsOutput, error) {
	if params == nil {
		params = &GetLicenseRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLicenseRecommendations", params, optFns, c.addOperationGetLicenseRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLicenseRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLicenseRecommendationsInput struct {

	// The ID of the Amazon Web Services account for which to return license
	// recommendations. If your account is the management account of an organization,
	// use this parameter to specify the member account for which you want to return
	// license recommendations. Only one account ID can be specified per request.
	AccountIds []string

	// An array of objects to specify a filter that returns a more specific list of
	// license recommendations.
	Filters []types.LicenseRecommendationFilter

	// The maximum number of license recommendations to return with a single request.
	// To retrieve the remaining results, make another request with the returned
	// nextToken value.
	MaxResults *int32

	// The token to advance to the next page of license recommendations.
	NextToken *string

	// The ARN that identifies the Amazon EC2 instance. The following is the format of
	// the ARN: arn:aws:ec2:region:aws_account_id:instance/instance-id
	ResourceArns []string

	noSmithyDocumentSerde
}

type GetLicenseRecommendationsOutput struct {

	// An array of objects that describe errors of the request.
	Errors []types.GetRecommendationError

	// An array of objects that describe license recommendations.
	LicenseRecommendations []types.LicenseRecommendation

	// The token to use to advance to the next page of license recommendations.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLicenseRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetLicenseRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetLicenseRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetLicenseRecommendations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLicenseRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLicenseRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetLicenseRecommendations",
	}
}
