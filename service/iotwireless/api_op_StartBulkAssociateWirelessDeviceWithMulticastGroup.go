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

// Starts a bulk association of all qualifying wireless devices with a multicast
// group.
func (c *Client) StartBulkAssociateWirelessDeviceWithMulticastGroup(ctx context.Context, params *StartBulkAssociateWirelessDeviceWithMulticastGroupInput, optFns ...func(*Options)) (*StartBulkAssociateWirelessDeviceWithMulticastGroupOutput, error) {
	if params == nil {
		params = &StartBulkAssociateWirelessDeviceWithMulticastGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartBulkAssociateWirelessDeviceWithMulticastGroup", params, optFns, c.addOperationStartBulkAssociateWirelessDeviceWithMulticastGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartBulkAssociateWirelessDeviceWithMulticastGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartBulkAssociateWirelessDeviceWithMulticastGroupInput struct {

	// The ID of the multicast group.
	//
	// This member is required.
	Id *string

	// Query string used to search for wireless devices as part of the bulk associate
	// and disassociate process.
	QueryString *string

	// The tag to attach to the specified resource. Tags are metadata that you can use
	// to manage a resource.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type StartBulkAssociateWirelessDeviceWithMulticastGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartBulkAssociateWirelessDeviceWithMulticastGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartBulkAssociateWirelessDeviceWithMulticastGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartBulkAssociateWirelessDeviceWithMulticastGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartBulkAssociateWirelessDeviceWithMulticastGroup"); err != nil {
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
	if err = addOpStartBulkAssociateWirelessDeviceWithMulticastGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartBulkAssociateWirelessDeviceWithMulticastGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartBulkAssociateWirelessDeviceWithMulticastGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartBulkAssociateWirelessDeviceWithMulticastGroup",
	}
}
