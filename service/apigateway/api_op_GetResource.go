// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists information about a resource.
func (c *Client) GetResource(ctx context.Context, params *GetResourceInput, optFns ...func(*Options)) (*GetResourceOutput, error) {
	if params == nil {
		params = &GetResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResource", params, optFns, c.addOperationGetResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to list information about a resource.
type GetResourceInput struct {

	// The identifier for the Resource resource.
	//
	// This member is required.
	ResourceId *string

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// A query parameter to retrieve the specified resources embedded in the returned
	// Resource representation in the response. This embed parameter value is a list
	// of comma-separated strings. Currently, the request supports only retrieval of
	// the embedded Method resources this way. The query parameter value must be a
	// single-valued list and contain the "methods" string. For example, GET
	// /restapis/{restapi_id}/resources/{resource_id}?embed=methods .
	Embed []string

	noSmithyDocumentSerde
}

// Represents an API resource.
type GetResourceOutput struct {

	// The resource's identifier.
	Id *string

	// The parent resource's identifier.
	ParentId *string

	// The full path for this resource.
	Path *string

	// The last path segment for this resource.
	PathPart *string

	// Gets an API resource's method of a given HTTP verb.
	ResourceMethods map[string]types.Method

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetResource"); err != nil {
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
	if err = addOpGetResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResource(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

func newServiceMetadataMiddleware_opGetResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetResource",
	}
}
