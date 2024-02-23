// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides access to an Amazon OpenSearch Service domain through the use of an
// interface VPC endpoint.
func (c *Client) AuthorizeVpcEndpointAccess(ctx context.Context, params *AuthorizeVpcEndpointAccessInput, optFns ...func(*Options)) (*AuthorizeVpcEndpointAccessOutput, error) {
	if params == nil {
		params = &AuthorizeVpcEndpointAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AuthorizeVpcEndpointAccess", params, optFns, c.addOperationAuthorizeVpcEndpointAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AuthorizeVpcEndpointAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AuthorizeVpcEndpointAccessInput struct {

	// The Amazon Web Services account ID to grant access to.
	//
	// This member is required.
	Account *string

	// The name of the OpenSearch Service domain to provide access to.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

type AuthorizeVpcEndpointAccessOutput struct {

	// Information about the Amazon Web Services account or service that was provided
	// access to the domain.
	//
	// This member is required.
	AuthorizedPrincipal *types.AuthorizedPrincipal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAuthorizeVpcEndpointAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAuthorizeVpcEndpointAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAuthorizeVpcEndpointAccess{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AuthorizeVpcEndpointAccess"); err != nil {
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
	if err = addOpAuthorizeVpcEndpointAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAuthorizeVpcEndpointAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAuthorizeVpcEndpointAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AuthorizeVpcEndpointAccess",
	}
}
