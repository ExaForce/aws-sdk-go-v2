// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers an on-premises instance. Only one IAM ARN (an IAM session ARN or IAM
// user ARN) is supported in the request. You cannot use both.
func (c *Client) RegisterOnPremisesInstance(ctx context.Context, params *RegisterOnPremisesInstanceInput, optFns ...func(*Options)) (*RegisterOnPremisesInstanceOutput, error) {
	if params == nil {
		params = &RegisterOnPremisesInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterOnPremisesInstance", params, optFns, c.addOperationRegisterOnPremisesInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterOnPremisesInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of the register on-premises instance operation.
type RegisterOnPremisesInstanceInput struct {

	// The name of the on-premises instance to register.
	//
	// This member is required.
	InstanceName *string

	// The ARN of the IAM session to associate with the on-premises instance.
	IamSessionArn *string

	// The ARN of the user to associate with the on-premises instance.
	IamUserArn *string

	noSmithyDocumentSerde
}

type RegisterOnPremisesInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterOnPremisesInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterOnPremisesInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterOnPremisesInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RegisterOnPremisesInstance"); err != nil {
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
	if err = addOpRegisterOnPremisesInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterOnPremisesInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterOnPremisesInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RegisterOnPremisesInstance",
	}
}
