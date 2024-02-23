// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers the IAM role that grants Amazon Inspector access to AWS Services
// needed to perform security assessments.
func (c *Client) RegisterCrossAccountAccessRole(ctx context.Context, params *RegisterCrossAccountAccessRoleInput, optFns ...func(*Options)) (*RegisterCrossAccountAccessRoleOutput, error) {
	if params == nil {
		params = &RegisterCrossAccountAccessRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterCrossAccountAccessRole", params, optFns, c.addOperationRegisterCrossAccountAccessRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterCrossAccountAccessRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterCrossAccountAccessRoleInput struct {

	// The ARN of the IAM role that grants Amazon Inspector access to AWS Services
	// needed to perform security assessments.
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

type RegisterCrossAccountAccessRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterCrossAccountAccessRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterCrossAccountAccessRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterCrossAccountAccessRole{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterCrossAccountAccessRole"); err != nil {
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
	if err = addOpRegisterCrossAccountAccessRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterCrossAccountAccessRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterCrossAccountAccessRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterCrossAccountAccessRole",
	}
}
