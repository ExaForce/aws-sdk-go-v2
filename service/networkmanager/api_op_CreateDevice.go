// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new device in a global network. If you specify both a site ID and a
// location, the location of the site is used for visualization in the Network
// Manager console.
func (c *Client) CreateDevice(ctx context.Context, params *CreateDeviceInput, optFns ...func(*Options)) (*CreateDeviceOutput, error) {
	if params == nil {
		params = &CreateDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDevice", params, optFns, c.addOperationCreateDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDeviceInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The Amazon Web Services location of the device, if applicable. For an
	// on-premises device, you can omit this parameter.
	AWSLocation *types.AWSLocation

	// A description of the device. Constraints: Maximum length of 256 characters.
	Description *string

	// The location of the device.
	Location *types.Location

	// The model of the device. Constraints: Maximum length of 128 characters.
	Model *string

	// The serial number of the device. Constraints: Maximum length of 128 characters.
	SerialNumber *string

	// The ID of the site.
	SiteId *string

	// The tags to apply to the resource during creation.
	Tags []types.Tag

	// The type of the device.
	Type *string

	// The vendor of the device. Constraints: Maximum length of 128 characters.
	Vendor *string

	noSmithyDocumentSerde
}

type CreateDeviceOutput struct {

	// Information about the device.
	Device *types.Device

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDevice"); err != nil {
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
	if err = addOpCreateDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDevice",
	}
}
