// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a collection in an AWS Region. You can add faces to the collection
// using the IndexFaces operation. For example, you might create collections, one
// for each of your application users. A user can then index faces using the
// IndexFaces operation and persist results in a specific collection. Then, a user
// can search the collection for faces in the user-specific container. When you
// create a collection, it is associated with the latest version of the face model
// version. Collection names are case-sensitive. This operation requires
// permissions to perform the rekognition:CreateCollection action. If you want to
// tag your collection, you also require permission to perform the
// rekognition:TagResource operation.
func (c *Client) CreateCollection(ctx context.Context, params *CreateCollectionInput, optFns ...func(*Options)) (*CreateCollectionOutput, error) {
	if params == nil {
		params = &CreateCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCollection", params, optFns, c.addOperationCreateCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCollectionInput struct {

	// ID for the collection that you are creating.
	//
	// This member is required.
	CollectionId *string

	// A set of tags (key-value pairs) that you want to attach to the collection.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateCollectionOutput struct {

	// Amazon Resource Name (ARN) of the collection. You can use this to manage
	// permissions on your resources.
	CollectionArn *string

	// Version number of the face detection model associated with the collection you
	// are creating.
	FaceModelVersion *string

	// HTTP status code indicating the result of the operation.
	StatusCode *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCollection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCollection"); err != nil {
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
	if err = addOpCreateCollectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCollection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCollection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCollection",
	}
}
