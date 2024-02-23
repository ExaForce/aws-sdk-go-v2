// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about a parallel data resource.
func (c *Client) GetParallelData(ctx context.Context, params *GetParallelDataInput, optFns ...func(*Options)) (*GetParallelDataOutput, error) {
	if params == nil {
		params = &GetParallelDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetParallelData", params, optFns, c.addOperationGetParallelDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetParallelDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetParallelDataInput struct {

	// The name of the parallel data resource that is being retrieved.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetParallelDataOutput struct {

	// The Amazon S3 location of a file that provides any errors or warnings that were
	// produced by your input file. This file was created when Amazon Translate
	// attempted to create a parallel data resource. The location is returned as a
	// presigned URL to that has a 30-minute expiration.
	AuxiliaryDataLocation *types.ParallelDataDataLocation

	// The Amazon S3 location of the most recent parallel data input file that was
	// successfully imported into Amazon Translate. The location is returned as a
	// presigned URL that has a 30-minute expiration. Amazon Translate doesn't scan all
	// input files for the risk of CSV injection attacks. CSV injection occurs when a
	// .csv or .tsv file is altered so that a record contains malicious code. The
	// record begins with a special character, such as =, +, -, or @. When the file is
	// opened in a spreadsheet program, the program might interpret the record as a
	// formula and run the code within it. Before you download an input file from
	// Amazon S3, ensure that you recognize the file and trust its creator.
	DataLocation *types.ParallelDataDataLocation

	// The Amazon S3 location of a file that provides any errors or warnings that were
	// produced by your input file. This file was created when Amazon Translate
	// attempted to update a parallel data resource. The location is returned as a
	// presigned URL to that has a 30-minute expiration.
	LatestUpdateAttemptAuxiliaryDataLocation *types.ParallelDataDataLocation

	// The properties of the parallel data resource that is being retrieved.
	ParallelDataProperties *types.ParallelDataProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetParallelDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetParallelData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetParallelData{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetParallelData"); err != nil {
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
	if err = addOpGetParallelDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetParallelData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetParallelData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetParallelData",
	}
}
