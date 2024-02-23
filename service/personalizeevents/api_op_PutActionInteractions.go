// Code generated by smithy-go-codegen DO NOT EDIT.

package personalizeevents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/personalizeevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Records action interaction event data. An action interaction event is an
// interaction between a user and an action. For example, a user taking an action,
// such a enrolling in a membership program or downloading your app. For more
// information about recording action interactions, see Recording action
// interaction events (https://docs.aws.amazon.com/personalize/latest/dg/recording-action-interaction-events.html)
// . For more information about actions in an Actions dataset, see Actions dataset (https://docs.aws.amazon.com/personalize/latest/dg/actions-datasets.html)
// .
func (c *Client) PutActionInteractions(ctx context.Context, params *PutActionInteractionsInput, optFns ...func(*Options)) (*PutActionInteractionsOutput, error) {
	if params == nil {
		params = &PutActionInteractionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutActionInteractions", params, optFns, c.addOperationPutActionInteractionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutActionInteractionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutActionInteractionsInput struct {

	// A list of action interaction events from the session.
	//
	// This member is required.
	ActionInteractions []types.ActionInteraction

	// The ID of your action interaction event tracker. When you create an Action
	// interactions dataset, Amazon Personalize creates an action interaction event
	// tracker for you. For more information, see Action interaction event tracker ID (https://docs.aws.amazon.com/personalize/latest/dg/action-interaction-tracker-id.html)
	// .
	//
	// This member is required.
	TrackingId *string

	noSmithyDocumentSerde
}

type PutActionInteractionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutActionInteractionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutActionInteractions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutActionInteractions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutActionInteractions"); err != nil {
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
	if err = addOpPutActionInteractionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutActionInteractions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutActionInteractions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutActionInteractions",
	}
}
