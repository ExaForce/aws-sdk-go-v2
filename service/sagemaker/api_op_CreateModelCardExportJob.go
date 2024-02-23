// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon SageMaker Model Card export job.
func (c *Client) CreateModelCardExportJob(ctx context.Context, params *CreateModelCardExportJobInput, optFns ...func(*Options)) (*CreateModelCardExportJobOutput, error) {
	if params == nil {
		params = &CreateModelCardExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModelCardExportJob", params, optFns, c.addOperationCreateModelCardExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelCardExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelCardExportJobInput struct {

	// The name of the model card export job.
	//
	// This member is required.
	ModelCardExportJobName *string

	// The name or Amazon Resource Name (ARN) of the model card to export.
	//
	// This member is required.
	ModelCardName *string

	// The model card output configuration that specifies the Amazon S3 path for
	// exporting.
	//
	// This member is required.
	OutputConfig *types.ModelCardExportOutputConfig

	// The version of the model card to export. If a version is not provided, then the
	// latest version of the model card is exported.
	ModelCardVersion *int32

	noSmithyDocumentSerde
}

type CreateModelCardExportJobOutput struct {

	// The Amazon Resource Name (ARN) of the model card export job.
	//
	// This member is required.
	ModelCardExportJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateModelCardExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateModelCardExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateModelCardExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateModelCardExportJob"); err != nil {
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
	if err = addOpCreateModelCardExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModelCardExportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateModelCardExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateModelCardExportJob",
	}
}
