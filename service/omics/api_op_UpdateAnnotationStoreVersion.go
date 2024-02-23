// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the description of an annotation store version.
func (c *Client) UpdateAnnotationStoreVersion(ctx context.Context, params *UpdateAnnotationStoreVersionInput, optFns ...func(*Options)) (*UpdateAnnotationStoreVersionOutput, error) {
	if params == nil {
		params = &UpdateAnnotationStoreVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAnnotationStoreVersion", params, optFns, c.addOperationUpdateAnnotationStoreVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAnnotationStoreVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAnnotationStoreVersionInput struct {

	// The name of an annotation store.
	//
	// This member is required.
	Name *string

	// The name of an annotation store version.
	//
	// This member is required.
	VersionName *string

	// The description of an annotation store.
	Description *string

	noSmithyDocumentSerde
}

type UpdateAnnotationStoreVersionOutput struct {

	// The time stamp for when an annotation store version was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The description of an annotation store version.
	//
	// This member is required.
	Description *string

	// The annotation store version ID.
	//
	// This member is required.
	Id *string

	// The name of an annotation store.
	//
	// This member is required.
	Name *string

	// The status of an annotation store version.
	//
	// This member is required.
	Status types.VersionStatus

	// The annotation store ID.
	//
	// This member is required.
	StoreId *string

	// The time stamp for when an annotation store version was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The name of an annotation store version.
	//
	// This member is required.
	VersionName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAnnotationStoreVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAnnotationStoreVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAnnotationStoreVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAnnotationStoreVersion"); err != nil {
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
	if err = addEndpointPrefix_opUpdateAnnotationStoreVersionMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateAnnotationStoreVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAnnotationStoreVersion(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateAnnotationStoreVersionMiddleware struct {
}

func (*endpointPrefix_opUpdateAnnotationStoreVersionMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateAnnotationStoreVersionMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "analytics-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opUpdateAnnotationStoreVersionMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opUpdateAnnotationStoreVersionMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opUpdateAnnotationStoreVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAnnotationStoreVersion",
	}
}
