// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Revokes a specified subscription in Amazon DataZone.
func (c *Client) RevokeSubscription(ctx context.Context, params *RevokeSubscriptionInput, optFns ...func(*Options)) (*RevokeSubscriptionOutput, error) {
	if params == nil {
		params = &RevokeSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokeSubscription", params, optFns, c.addOperationRevokeSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokeSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeSubscriptionInput struct {

	// The identifier of the Amazon DataZone domain where you want to revoke a
	// subscription.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the revoked subscription.
	//
	// This member is required.
	Identifier *string

	// Specifies whether permissions are retained when the subscription is revoked.
	RetainPermissions *bool

	noSmithyDocumentSerde
}

type RevokeSubscriptionOutput struct {

	// The timestamp of when the subscription was revoked.
	//
	// This member is required.
	CreatedAt *time.Time

	// The identifier of the user who revoked the subscription.
	//
	// This member is required.
	CreatedBy *string

	// The identifier of the Amazon DataZone domain where you want to revoke a
	// subscription.
	//
	// This member is required.
	DomainId *string

	// The identifier of the revoked subscription.
	//
	// This member is required.
	Id *string

	// The status of the revoked subscription.
	//
	// This member is required.
	Status types.SubscriptionStatus

	// The subscribed listing of the revoked subscription.
	//
	// This member is required.
	SubscribedListing *types.SubscribedListing

	// The subscribed principal of the revoked subscription.
	//
	// This member is required.
	SubscribedPrincipal types.SubscribedPrincipal

	// The timestamp of when the subscription was revoked.
	//
	// This member is required.
	UpdatedAt *time.Time

	// Specifies whether permissions are retained when the subscription is revoked.
	RetainPermissions *bool

	// The identifier of the subscription request for the revoked subscription.
	SubscriptionRequestId *string

	// The Amazon DataZone user who revoked the subscription.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRevokeSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRevokeSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRevokeSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RevokeSubscription"); err != nil {
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
	if err = addOpRevokeSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRevokeSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RevokeSubscription",
	}
}
