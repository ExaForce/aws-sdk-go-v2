// Code generated by smithy-go-codegen DO NOT EDIT.

package backupgateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the gateway virtual machine (VM) software. The request immediately
// triggers the software update. When you make this request, you get a 200 OK
// success response immediately. However, it might take some time for the update to
// complete.
func (c *Client) UpdateGatewaySoftwareNow(ctx context.Context, params *UpdateGatewaySoftwareNowInput, optFns ...func(*Options)) (*UpdateGatewaySoftwareNowOutput, error) {
	if params == nil {
		params = &UpdateGatewaySoftwareNowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGatewaySoftwareNow", params, optFns, c.addOperationUpdateGatewaySoftwareNowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGatewaySoftwareNowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGatewaySoftwareNowInput struct {

	// The Amazon Resource Name (ARN) of the gateway to be updated.
	//
	// This member is required.
	GatewayArn *string

	noSmithyDocumentSerde
}

type UpdateGatewaySoftwareNowOutput struct {

	// The Amazon Resource Name (ARN) of the gateway you updated.
	GatewayArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGatewaySoftwareNowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateGatewaySoftwareNow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateGatewaySoftwareNow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateGatewaySoftwareNow"); err != nil {
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
	if err = addOpUpdateGatewaySoftwareNowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGatewaySoftwareNow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateGatewaySoftwareNow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateGatewaySoftwareNow",
	}
}
