// Code generated by smithy-go-codegen DO NOT EDIT.

package mediastore

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediastore/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the cross-origin resource sharing (CORS) configuration on a container so
// that the container can service cross-origin requests. For example, you might
// want to enable a request whose origin is http://www.example.com to access your
// AWS Elemental MediaStore container at my.example.container.com by using the
// browser's XMLHttpRequest capability. To enable CORS on a container, you attach a
// CORS policy to the container. In the CORS policy, you configure rules that
// identify origins and the HTTP methods that can be executed on your container.
// The policy can contain up to 398,000 characters. You can add up to 100 rules to
// a CORS policy. If more than one rule applies, the service uses the first
// applicable rule listed. To learn more about CORS, see Cross-Origin Resource
// Sharing (CORS) in AWS Elemental MediaStore (https://docs.aws.amazon.com/mediastore/latest/ug/cors-policy.html)
// .
func (c *Client) PutCorsPolicy(ctx context.Context, params *PutCorsPolicyInput, optFns ...func(*Options)) (*PutCorsPolicyOutput, error) {
	if params == nil {
		params = &PutCorsPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutCorsPolicy", params, optFns, c.addOperationPutCorsPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutCorsPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutCorsPolicyInput struct {

	// The name of the container that you want to assign the CORS policy to.
	//
	// This member is required.
	ContainerName *string

	// The CORS policy to apply to the container.
	//
	// This member is required.
	CorsPolicy []types.CorsRule

	noSmithyDocumentSerde
}

type PutCorsPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutCorsPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutCorsPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutCorsPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutCorsPolicy"); err != nil {
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
	if err = addOpPutCorsPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutCorsPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutCorsPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutCorsPolicy",
	}
}
