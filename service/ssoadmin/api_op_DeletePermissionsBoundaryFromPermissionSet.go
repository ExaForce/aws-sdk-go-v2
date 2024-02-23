// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the permissions boundary from a specified PermissionSet .
func (c *Client) DeletePermissionsBoundaryFromPermissionSet(ctx context.Context, params *DeletePermissionsBoundaryFromPermissionSetInput, optFns ...func(*Options)) (*DeletePermissionsBoundaryFromPermissionSetOutput, error) {
	if params == nil {
		params = &DeletePermissionsBoundaryFromPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePermissionsBoundaryFromPermissionSet", params, optFns, c.addOperationDeletePermissionsBoundaryFromPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePermissionsBoundaryFromPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePermissionsBoundaryFromPermissionSetInput struct {

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the PermissionSet .
	//
	// This member is required.
	PermissionSetArn *string

	noSmithyDocumentSerde
}

type DeletePermissionsBoundaryFromPermissionSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePermissionsBoundaryFromPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePermissionsBoundaryFromPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePermissionsBoundaryFromPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeletePermissionsBoundaryFromPermissionSet"); err != nil {
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
	if err = addOpDeletePermissionsBoundaryFromPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePermissionsBoundaryFromPermissionSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePermissionsBoundaryFromPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeletePermissionsBoundaryFromPermissionSet",
	}
}
