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

// Retrieves the cost anomaly subscription objects for your account. You can
// filter using a list of cost anomaly monitor Amazon Resource Names (ARNs).
func (c *Client) GetAnomalySubscriptions(ctx context.Context, params *GetAnomalySubscriptionsInput, optFns ...func(*Options)) (*GetAnomalySubscriptionsOutput, error) {
	if params == nil {
		params = &GetAnomalySubscriptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAnomalySubscriptions", params, optFns, c.addOperationGetAnomalySubscriptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAnomalySubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAnomalySubscriptionsInput struct {

	// The number of entries a paginated response contains.
	MaxResults *int32

	// Cost anomaly monitor ARNs.
	MonitorArn *string

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	// A list of cost anomaly subscription ARNs.
	SubscriptionArnList []string

	noSmithyDocumentSerde
}

type GetAnomalySubscriptionsOutput struct {

	// A list of cost anomaly subscriptions that includes the detailed metadata for
	// each one.
	//
	// This member is required.
	AnomalySubscriptions []types.AnomalySubscription

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAnomalySubscriptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAnomalySubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAnomalySubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAnomalySubscriptions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAnomalySubscriptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAnomalySubscriptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAnomalySubscriptions",
	}
}
