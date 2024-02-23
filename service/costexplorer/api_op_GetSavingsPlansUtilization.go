// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the Savings Plans utilization for your account across date ranges
// with daily or monthly granularity. Management account in an organization have
// access to member accounts. You can use GetDimensionValues in SAVINGS_PLANS to
// determine the possible dimension values. You can't group by any dimension values
// for GetSavingsPlansUtilization .
func (c *Client) GetSavingsPlansUtilization(ctx context.Context, params *GetSavingsPlansUtilizationInput, optFns ...func(*Options)) (*GetSavingsPlansUtilizationOutput, error) {
	if params == nil {
		params = &GetSavingsPlansUtilizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSavingsPlansUtilization", params, optFns, c.addOperationGetSavingsPlansUtilizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSavingsPlansUtilizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSavingsPlansUtilizationInput struct {

	// The time period that you want the usage and costs for. The Start date must be
	// within 13 months. The End date must be after the Start date, and before the
	// current date. Future dates can't be used as an End date.
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// Filters Savings Plans utilization coverage data for active Savings Plans
	// dimensions. You can filter data with the following dimensions:
	//   - LINKED_ACCOUNT
	//   - SAVINGS_PLAN_ARN
	//   - SAVINGS_PLANS_TYPE
	//   - REGION
	//   - PAYMENT_OPTION
	//   - INSTANCE_TYPE_FAMILY
	// GetSavingsPlansUtilization uses the same Expression (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_Expression.html)
	// object as the other operations, but only AND is supported among each dimension.
	Filter *types.Expression

	// The granularity of the Amazon Web Services utillization data for your Savings
	// Plans. The GetSavingsPlansUtilization operation supports only DAILY and MONTHLY
	// granularities.
	Granularity types.Granularity

	// The value that you want to sort the data by. The following values are supported
	// for Key :
	//   - UtilizationPercentage
	//   - TotalCommitment
	//   - UsedCommitment
	//   - UnusedCommitment
	//   - NetSavings
	// The supported values for SortOrder are ASCENDING and DESCENDING .
	SortBy *types.SortDefinition

	noSmithyDocumentSerde
}

type GetSavingsPlansUtilizationOutput struct {

	// The total amount of cost/commitment that you used your Savings Plans,
	// regardless of date ranges.
	//
	// This member is required.
	Total *types.SavingsPlansUtilizationAggregates

	// The amount of cost/commitment that you used your Savings Plans. You can use it
	// to specify date ranges.
	SavingsPlansUtilizationsByTime []types.SavingsPlansUtilizationByTime

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSavingsPlansUtilizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSavingsPlansUtilization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSavingsPlansUtilization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSavingsPlansUtilization"); err != nil {
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
	if err = addOpGetSavingsPlansUtilizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSavingsPlansUtilization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSavingsPlansUtilization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSavingsPlansUtilization",
	}
}
