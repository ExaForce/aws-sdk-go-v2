// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update the event configuration for a particular resource identifier.
func (c *Client) UpdateResourceEventConfiguration(ctx context.Context, params *UpdateResourceEventConfigurationInput, optFns ...func(*Options)) (*UpdateResourceEventConfigurationOutput, error) {
	if params == nil {
		params = &UpdateResourceEventConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResourceEventConfiguration", params, optFns, c.addOperationUpdateResourceEventConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResourceEventConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResourceEventConfigurationInput struct {

	// Resource identifier to opt in for event messaging.
	//
	// This member is required.
	Identifier *string

	// Identifier type of the particular resource identifier for event configuration.
	//
	// This member is required.
	IdentifierType types.IdentifierType

	// Event configuration for the connection status event.
	ConnectionStatus *types.ConnectionStatusEventConfiguration

	// Event configuration for the device registration state event.
	DeviceRegistrationState *types.DeviceRegistrationStateEventConfiguration

	// Event configuration for the join event.
	Join *types.JoinEventConfiguration

	// Event configuration for the message delivery status event.
	MessageDeliveryStatus *types.MessageDeliveryStatusEventConfiguration

	// Partner type of the resource if the identifier type is PartnerAccountId
	PartnerType types.EventNotificationPartnerType

	// Event configuration for the proximity event.
	Proximity *types.ProximityEventConfiguration

	noSmithyDocumentSerde
}

type UpdateResourceEventConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateResourceEventConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateResourceEventConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateResourceEventConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateResourceEventConfiguration"); err != nil {
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
	if err = addOpUpdateResourceEventConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResourceEventConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateResourceEventConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateResourceEventConfiguration",
	}
}
