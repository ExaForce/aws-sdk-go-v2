// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Revokes the permission of another Amazon Web Services account to be able to put
// events to the specified event bus. Specify the account to revoke by the
// StatementId value that you associated with the account when you granted it
// permission with PutPermission . You can find the StatementId by using
// DescribeEventBus (https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_DescribeEventBus.html)
// .
func (c *Client) RemovePermission(ctx context.Context, params *RemovePermissionInput, optFns ...func(*Options)) (*RemovePermissionOutput, error) {
	if params == nil {
		params = &RemovePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemovePermission", params, optFns, c.addOperationRemovePermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemovePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemovePermissionInput struct {

	// The name of the event bus to revoke permissions for. If you omit this, the
	// default event bus is used.
	EventBusName *string

	// Specifies whether to remove all permissions.
	RemoveAllPermissions bool

	// The statement ID corresponding to the account that is no longer allowed to put
	// events to the default event bus.
	StatementId *string

	noSmithyDocumentSerde
}

type RemovePermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemovePermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRemovePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemovePermission{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RemovePermission"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemovePermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemovePermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RemovePermission",
	}
}
