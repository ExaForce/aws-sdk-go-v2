// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// For a batch of security controls and standards, identifies whether each control
// is currently enabled or disabled in a standard.
func (c *Client) BatchGetStandardsControlAssociations(ctx context.Context, params *BatchGetStandardsControlAssociationsInput, optFns ...func(*Options)) (*BatchGetStandardsControlAssociationsOutput, error) {
	if params == nil {
		params = &BatchGetStandardsControlAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetStandardsControlAssociations", params, optFns, c.addOperationBatchGetStandardsControlAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetStandardsControlAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetStandardsControlAssociationsInput struct {

	// An array with one or more objects that includes a security control (identified
	// with SecurityControlId , SecurityControlArn , or a mix of both parameters) and
	// the Amazon Resource Name (ARN) of a standard. This field is used to query the
	// enablement status of a control in a specified standard. The security control ID
	// or ARN is the same across standards.
	//
	// This member is required.
	StandardsControlAssociationIds []types.StandardsControlAssociationId

	noSmithyDocumentSerde
}

type BatchGetStandardsControlAssociationsOutput struct {

	// Provides the enablement status of a security control in a specified standard
	// and other details for the control in relation to the specified standard.
	//
	// This member is required.
	StandardsControlAssociationDetails []types.StandardsControlAssociationDetail

	// A security control (identified with SecurityControlId , SecurityControlArn , or
	// a mix of both parameters) whose enablement status in a specified standard cannot
	// be returned.
	UnprocessedAssociations []types.UnprocessedStandardsControlAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetStandardsControlAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetStandardsControlAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetStandardsControlAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetStandardsControlAssociations"); err != nil {
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
	if err = addOpBatchGetStandardsControlAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetStandardsControlAssociations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetStandardsControlAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetStandardsControlAssociations",
	}
}
