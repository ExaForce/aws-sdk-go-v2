// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enable or disable the automatic warm-up feature for dedicated IP addresses.
func (c *Client) PutAccountDedicatedIpWarmupAttributes(ctx context.Context, params *PutAccountDedicatedIpWarmupAttributesInput, optFns ...func(*Options)) (*PutAccountDedicatedIpWarmupAttributesOutput, error) {
	if params == nil {
		params = &PutAccountDedicatedIpWarmupAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccountDedicatedIpWarmupAttributes", params, optFns, c.addOperationPutAccountDedicatedIpWarmupAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccountDedicatedIpWarmupAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to enable or disable the automatic IP address warm-up feature.
type PutAccountDedicatedIpWarmupAttributesInput struct {

	// Enables or disables the automatic warm-up feature for dedicated IP addresses
	// that are associated with your Amazon Pinpoint account in the current AWS Region.
	// Set to true to enable the automatic warm-up feature, or set to false to disable
	// it.
	AutoWarmupEnabled bool

	noSmithyDocumentSerde
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type PutAccountDedicatedIpWarmupAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAccountDedicatedIpWarmupAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutAccountDedicatedIpWarmupAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutAccountDedicatedIpWarmupAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutAccountDedicatedIpWarmupAttributes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountDedicatedIpWarmupAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAccountDedicatedIpWarmupAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutAccountDedicatedIpWarmupAttributes",
	}
}
