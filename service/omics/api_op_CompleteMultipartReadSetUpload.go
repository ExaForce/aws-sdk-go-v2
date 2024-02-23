// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Concludes a multipart upload once you have uploaded all the components.
func (c *Client) CompleteMultipartReadSetUpload(ctx context.Context, params *CompleteMultipartReadSetUploadInput, optFns ...func(*Options)) (*CompleteMultipartReadSetUploadOutput, error) {
	if params == nil {
		params = &CompleteMultipartReadSetUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CompleteMultipartReadSetUpload", params, optFns, c.addOperationCompleteMultipartReadSetUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CompleteMultipartReadSetUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CompleteMultipartReadSetUploadInput struct {

	// The individual uploads or parts of a multipart upload.
	//
	// This member is required.
	Parts []types.CompleteReadSetUploadPartListItem

	// The sequence store ID for the store involved in the multipart upload.
	//
	// This member is required.
	SequenceStoreId *string

	// The ID for the multipart upload.
	//
	// This member is required.
	UploadId *string

	noSmithyDocumentSerde
}

type CompleteMultipartReadSetUploadOutput struct {

	// The read set ID created for an uploaded read set.
	//
	// This member is required.
	ReadSetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCompleteMultipartReadSetUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCompleteMultipartReadSetUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCompleteMultipartReadSetUpload{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CompleteMultipartReadSetUpload"); err != nil {
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
	if err = addEndpointPrefix_opCompleteMultipartReadSetUploadMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCompleteMultipartReadSetUploadValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCompleteMultipartReadSetUpload(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCompleteMultipartReadSetUploadMiddleware struct {
}

func (*endpointPrefix_opCompleteMultipartReadSetUploadMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCompleteMultipartReadSetUploadMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "storage-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opCompleteMultipartReadSetUploadMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opCompleteMultipartReadSetUploadMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opCompleteMultipartReadSetUpload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CompleteMultipartReadSetUpload",
	}
}
