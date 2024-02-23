// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Verifies the specified user attributes in the user pool. If your user pool
// requires verification before Amazon Cognito updates the attribute value,
// VerifyUserAttribute updates the affected attribute to its pending value. For
// more information, see UserAttributeUpdateSettingsType (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_UserAttributeUpdateSettingsType.html)
// . Authorize this action with a signed-in user's access token. It must include
// the scope aws.cognito.signin.user.admin . Amazon Cognito doesn't evaluate
// Identity and Access Management (IAM) policies in requests for this API
// operation. For this operation, you can't use IAM credentials to authorize
// requests, and you can't grant IAM permissions in policies. For more information
// about authorization models in Amazon Cognito, see Using the Amazon Cognito user
// pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
// .
func (c *Client) VerifyUserAttribute(ctx context.Context, params *VerifyUserAttributeInput, optFns ...func(*Options)) (*VerifyUserAttributeOutput, error) {
	if params == nil {
		params = &VerifyUserAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "VerifyUserAttribute", params, optFns, c.addOperationVerifyUserAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*VerifyUserAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to verify user attributes.
type VerifyUserAttributeInput struct {

	// A valid access token that Amazon Cognito issued to the user whose user
	// attributes you want to verify.
	//
	// This member is required.
	AccessToken *string

	// The attribute name in the request to verify user attributes.
	//
	// This member is required.
	AttributeName *string

	// The verification code in the request to verify user attributes.
	//
	// This member is required.
	Code *string

	noSmithyDocumentSerde
}

// A container representing the response from the server from the request to
// verify user attributes.
type VerifyUserAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationVerifyUserAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpVerifyUserAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpVerifyUserAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "VerifyUserAttribute"); err != nil {
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
	if err = addOpVerifyUserAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opVerifyUserAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opVerifyUserAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "VerifyUserAttribute",
	}
}
