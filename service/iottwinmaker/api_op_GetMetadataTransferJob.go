// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets a nmetadata transfer job.
func (c *Client) GetMetadataTransferJob(ctx context.Context, params *GetMetadataTransferJobInput, optFns ...func(*Options)) (*GetMetadataTransferJobOutput, error) {
	if params == nil {
		params = &GetMetadataTransferJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMetadataTransferJob", params, optFns, c.addOperationGetMetadataTransferJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMetadataTransferJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMetadataTransferJobInput struct {

	// The metadata transfer job Id.
	//
	// This member is required.
	MetadataTransferJobId *string

	noSmithyDocumentSerde
}

type GetMetadataTransferJobOutput struct {

	// The metadata transfer job ARN.
	//
	// This member is required.
	Arn *string

	// The metadata transfer job's creation DateTime property.
	//
	// This member is required.
	CreationDateTime *time.Time

	// The metadata transfer job's destination.
	//
	// This member is required.
	Destination *types.DestinationConfiguration

	// The metadata transfer job Id.
	//
	// This member is required.
	MetadataTransferJobId *string

	// The metadata transfer job's role.
	//
	// This member is required.
	MetadataTransferJobRole *string

	// The metadata transfer job's sources.
	//
	// This member is required.
	Sources []types.SourceConfiguration

	// The metadata transfer job's status.
	//
	// This member is required.
	Status *types.MetadataTransferJobStatus

	// The metadata transfer job's update DateTime property.
	//
	// This member is required.
	UpdateDateTime *time.Time

	// The metadata transfer job description.
	Description *string

	// The metadata transfer job's progress.
	Progress *types.MetadataTransferJobProgress

	// The metadata transfer job's report URL.
	ReportUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMetadataTransferJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMetadataTransferJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMetadataTransferJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMetadataTransferJob"); err != nil {
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
	if err = addEndpointPrefix_opGetMetadataTransferJobMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetMetadataTransferJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMetadataTransferJob(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetMetadataTransferJobMiddleware struct {
}

func (*endpointPrefix_opGetMetadataTransferJobMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetMetadataTransferJobMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetMetadataTransferJobMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetMetadataTransferJobMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetMetadataTransferJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMetadataTransferJob",
	}
}
