// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers an Alexa-enabled device built by an Original Equipment Manufacturer
// (OEM) using Alexa Voice Service (AVS).
//
// Deprecated: Alexa For Business is no longer supported
func (c *Client) RegisterAVSDevice(ctx context.Context, params *RegisterAVSDeviceInput, optFns ...func(*Options)) (*RegisterAVSDeviceOutput, error) {
	if params == nil {
		params = &RegisterAVSDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterAVSDevice", params, optFns, c.addOperationRegisterAVSDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterAVSDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterAVSDeviceInput struct {

	// The device type ID for your AVS device generated by Amazon when the OEM creates
	// a new product on Amazon's Developer Console.
	//
	// This member is required.
	AmazonId *string

	// The client ID of the OEM used for code-based linking authorization on an AVS
	// device.
	//
	// This member is required.
	ClientId *string

	// The product ID used to identify your AVS device during authorization.
	//
	// This member is required.
	ProductId *string

	// The code that is obtained after your AVS device has made a POST request to LWA
	// as a part of the Device Authorization Request component of the OAuth code-based
	// linking specification.
	//
	// This member is required.
	UserCode *string

	// The key generated by the OEM that uniquely identifies a specified instance of
	// your AVS device.
	DeviceSerialNumber *string

	// The Amazon Resource Name (ARN) of the room with which to associate your AVS
	// device.
	RoomArn *string

	// The tags to be added to the specified resource. Do not provide system tags.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type RegisterAVSDeviceOutput struct {

	// The ARN of the device.
	DeviceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterAVSDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterAVSDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterAVSDevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterAVSDevice"); err != nil {
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
	if err = addOpRegisterAVSDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterAVSDevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterAVSDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterAVSDevice",
	}
}
