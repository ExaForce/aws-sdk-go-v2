// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/tnb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the details of a network function instance, including the instantation
// state and metadata from the function package descriptor in the network function
// package. A network function instance is a function in a function package .
func (c *Client) GetSolFunctionInstance(ctx context.Context, params *GetSolFunctionInstanceInput, optFns ...func(*Options)) (*GetSolFunctionInstanceOutput, error) {
	if params == nil {
		params = &GetSolFunctionInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSolFunctionInstance", params, optFns, c.addOperationGetSolFunctionInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSolFunctionInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSolFunctionInstanceInput struct {

	// ID of the network function.
	//
	// This member is required.
	VnfInstanceId *string

	noSmithyDocumentSerde
}

type GetSolFunctionInstanceOutput struct {

	// Network function instance ARN.
	//
	// This member is required.
	Arn *string

	// Network function instance ID.
	//
	// This member is required.
	Id *string

	// Network function instantiation state.
	//
	// This member is required.
	InstantiationState types.VnfInstantiationState

	// The metadata of a network function instance. A network function instance is a
	// function in a function package .
	//
	// This member is required.
	Metadata *types.GetSolFunctionInstanceMetadata

	// Network instance ID.
	//
	// This member is required.
	NsInstanceId *string

	// Function package ID.
	//
	// This member is required.
	VnfPkgId *string

	// Function package descriptor ID.
	//
	// This member is required.
	VnfdId *string

	// Information about the network function. A network function instance is a
	// function in a function package .
	InstantiatedVnfInfo *types.GetSolVnfInfo

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value. You can use tags to search and filter
	// your resources or track your Amazon Web Services costs.
	Tags map[string]string

	// Network function product name.
	VnfProductName *string

	// Network function provider.
	VnfProvider *string

	// Function package descriptor version.
	VnfdVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSolFunctionInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSolFunctionInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSolFunctionInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSolFunctionInstance"); err != nil {
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
	if err = addOpGetSolFunctionInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSolFunctionInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSolFunctionInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSolFunctionInstance",
	}
}
