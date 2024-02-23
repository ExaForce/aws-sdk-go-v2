// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a collection of resources, including folders and documents. The only
// CollectionType supported is SHARED_WITH_ME .
func (c *Client) GetResources(ctx context.Context, params *GetResourcesInput, optFns ...func(*Options)) (*GetResourcesOutput, error) {
	if params == nil {
		params = &GetResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResources", params, optFns, c.addOperationGetResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourcesInput struct {

	// The Amazon WorkDocs authentication token. Not required when using Amazon Web
	// Services administrator credentials to access the API.
	AuthenticationToken *string

	// The collection type.
	CollectionType types.ResourceCollectionType

	// The maximum number of resources to return.
	Limit *int32

	// The marker for the next set of results. This marker was received from a
	// previous call.
	Marker *string

	// The user ID for the resource collection. This is a required field for accessing
	// the API operation using IAM credentials.
	UserId *string

	noSmithyDocumentSerde
}

type GetResourcesOutput struct {

	// The documents in the specified collection.
	Documents []types.DocumentMetadata

	// The folders in the specified folder.
	Folders []types.FolderMetadata

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetResources"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetResources",
	}
}
