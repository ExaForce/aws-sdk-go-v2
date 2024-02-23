// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an association between a Merged API and source API using the Merged
// API's identifier and the association ID.
func (c *Client) DisassociateSourceGraphqlApi(ctx context.Context, params *DisassociateSourceGraphqlApiInput, optFns ...func(*Options)) (*DisassociateSourceGraphqlApiOutput, error) {
	if params == nil {
		params = &DisassociateSourceGraphqlApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateSourceGraphqlApi", params, optFns, c.addOperationDisassociateSourceGraphqlApiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateSourceGraphqlApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateSourceGraphqlApiInput struct {

	// The ID generated by the AppSync service for the source API association.
	//
	// This member is required.
	AssociationId *string

	// The identifier of the AppSync Merged API. This is generated by the AppSync
	// service. In most cases, Merged APIs (especially in your account) only require
	// the API ID value or ARN of the merged API. However, Merged APIs in other
	// accounts (cross-account use cases) strictly require the full resource ARN of the
	// merged API.
	//
	// This member is required.
	MergedApiIdentifier *string

	noSmithyDocumentSerde
}

type DisassociateSourceGraphqlApiOutput struct {

	// The state of the source API association.
	SourceApiAssociationStatus types.SourceApiAssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateSourceGraphqlApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateSourceGraphqlApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateSourceGraphqlApi{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisassociateSourceGraphqlApi"); err != nil {
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
	if err = addOpDisassociateSourceGraphqlApiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateSourceGraphqlApi(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateSourceGraphqlApi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisassociateSourceGraphqlApi",
	}
}
