// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/tnb/document"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Instantiates a network instance. A network instance is a single network created
// in Amazon Web Services TNB that can be deployed and on which life-cycle
// operations (like terminate, update, and delete) can be performed. Before you can
// instantiate a network instance, you have to create a network instance. For more
// information, see CreateSolNetworkInstance (https://docs.aws.amazon.com/tnb/latest/APIReference/API_CreateSolNetworkInstance.html)
// .
func (c *Client) InstantiateSolNetworkInstance(ctx context.Context, params *InstantiateSolNetworkInstanceInput, optFns ...func(*Options)) (*InstantiateSolNetworkInstanceOutput, error) {
	if params == nil {
		params = &InstantiateSolNetworkInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InstantiateSolNetworkInstance", params, optFns, c.addOperationInstantiateSolNetworkInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InstantiateSolNetworkInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InstantiateSolNetworkInstanceInput struct {

	// ID of the network instance.
	//
	// This member is required.
	NsInstanceId *string

	// Provides values for the configurable properties.
	AdditionalParamsForNs document.Interface

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value. When you use this API, the tags are
	// transferred to the network operation that is created. Use tags to search and
	// filter your resources or track your Amazon Web Services costs.
	Tags map[string]string

	noSmithyDocumentSerde
}

type InstantiateSolNetworkInstanceOutput struct {

	// The identifier of the network operation.
	//
	// This member is required.
	NsLcmOpOccId *string

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value. When you use this API, the tags are
	// transferred to the network operation that is created. Use tags to search and
	// filter your resources or track your Amazon Web Services costs.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInstantiateSolNetworkInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInstantiateSolNetworkInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInstantiateSolNetworkInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InstantiateSolNetworkInstance"); err != nil {
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
	if err = addOpInstantiateSolNetworkInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInstantiateSolNetworkInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInstantiateSolNetworkInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InstantiateSolNetworkInstance",
	}
}
