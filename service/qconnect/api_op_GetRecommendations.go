// Code generated by smithy-go-codegen DO NOT EDIT.

package qconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves recommendations for the specified session. To avoid retrieving the
// same recommendations in subsequent calls, use NotifyRecommendationsReceived (https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_NotifyRecommendationsReceived.html)
// . This API supports long-polling behavior with the waitTimeSeconds parameter.
// Short poll is the default behavior and only returns recommendations already
// available. To perform a manual query against an assistant, use QueryAssistant (https://docs.aws.amazon.com/amazon-q-connect/latest/APIReference/API_QueryAssistant.html)
// .
//
// Deprecated: GetRecommendations API will be discontinued starting June 1, 2024.
// To receive generative responses after March 1, 2024 you will need to create a
// new Assistant in the Connect console and integrate the Amazon Q in Connect
// JavaScript library (amazon-q-connectjs) into your applications.
func (c *Client) GetRecommendations(ctx context.Context, params *GetRecommendationsInput, optFns ...func(*Options)) (*GetRecommendationsOutput, error) {
	if params == nil {
		params = &GetRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRecommendations", params, optFns, c.addOperationGetRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRecommendationsInput struct {

	// The identifier of the Amazon Q assistant. Can be either the ID or the ARN. URLs
	// cannot contain the ARN.
	//
	// This member is required.
	AssistantId *string

	// The identifier of the session. Can be either the ID or the ARN. URLs cannot
	// contain the ARN.
	//
	// This member is required.
	SessionId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The duration (in seconds) for which the call waits for a recommendation to be
	// made available before returning. If a recommendation is available, the call
	// returns sooner than WaitTimeSeconds . If no messages are available and the wait
	// time expires, the call returns successfully with an empty list.
	WaitTimeSeconds int32

	noSmithyDocumentSerde
}

type GetRecommendationsOutput struct {

	// The recommendations.
	//
	// This member is required.
	Recommendations []types.RecommendationData

	// The triggers corresponding to recommendations.
	Triggers []types.RecommendationTrigger

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRecommendations"); err != nil {
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
	if err = addOpGetRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRecommendations",
	}
}
