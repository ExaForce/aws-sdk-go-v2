// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// With this operation, your users can update one or more of their attributes with
// their own credentials. You authorize this API request with the user's access
// token. To delete an attribute from your user, submit the attribute in your API
// request with a blank value. Custom attribute values in this request must include
// the custom: prefix. Authorize this action with a signed-in user's access token.
// It must include the scope aws.cognito.signin.user.admin . Amazon Cognito doesn't
// evaluate Identity and Access Management (IAM) policies in requests for this API
// operation. For this operation, you can't use IAM credentials to authorize
// requests, and you can't grant IAM permissions in policies. For more information
// about authorization models in Amazon Cognito, see Using the Amazon Cognito user
// pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
// . This action might generate an SMS text message. Starting June 1, 2021, US
// telecom carriers require you to register an origination phone number before you
// can send SMS messages to US phone numbers. If you use SMS text messages in
// Amazon Cognito, you must register a phone number with Amazon Pinpoint (https://console.aws.amazon.com/pinpoint/home/)
// . Amazon Cognito uses the registered number automatically. Otherwise, Amazon
// Cognito users who must receive SMS messages might not be able to sign up,
// activate their accounts, or sign in. If you have never used SMS text messages
// with Amazon Cognito or any other Amazon Web Service, Amazon Simple Notification
// Service might place your account in the SMS sandbox. In sandbox mode (https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html)
// , you can send messages only to verified phone numbers. After you test your app
// while in the sandbox environment, you can move out of the sandbox and into
// production. For more information, see SMS message settings for Amazon Cognito
// user pools (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-sms-settings.html)
// in the Amazon Cognito Developer Guide.
func (c *Client) UpdateUserAttributes(ctx context.Context, params *UpdateUserAttributesInput, optFns ...func(*Options)) (*UpdateUserAttributesOutput, error) {
	if params == nil {
		params = &UpdateUserAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUserAttributes", params, optFns, c.addOperationUpdateUserAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUserAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to update user attributes.
type UpdateUserAttributesInput struct {

	// A valid access token that Amazon Cognito issued to the user whose user
	// attributes you want to update.
	//
	// This member is required.
	AccessToken *string

	// An array of name-value pairs representing user attributes. For custom
	// attributes, you must prepend the custom: prefix to the attribute name. If you
	// have set an attribute to require verification before Amazon Cognito updates its
	// value, this request doesn’t immediately update the value of that attribute.
	// After your user receives and responds to a verification message to verify the
	// new value, Amazon Cognito updates the attribute value. Your user can sign in and
	// receive messages with the original attribute value until they verify the new
	// value.
	//
	// This member is required.
	UserAttributes []types.AttributeType

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action initiates. You create custom workflows by assigning
	// Lambda functions to user pool triggers. When you use the UpdateUserAttributes
	// API action, Amazon Cognito invokes the function that is assigned to the custom
	// message trigger. When Amazon Cognito invokes this function, it passes a JSON
	// payload, which the function receives as input. This payload contains a
	// clientMetadata attribute, which provides the data that you assigned to the
	// ClientMetadata parameter in your UpdateUserAttributes request. In your function
	// code in Lambda, you can process the clientMetadata value to enhance your
	// workflow for your specific needs. For more information, see Customizing user
	// pool Workflows with Lambda Triggers (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. When you use the ClientMetadata
	// parameter, remember that Amazon Cognito won't do the following:
	//   - Store the ClientMetadata value. This data is available only to Lambda
	//   triggers that are assigned to a user pool to support custom workflows. If your
	//   user pool configuration doesn't include triggers, the ClientMetadata parameter
	//   serves no purpose.
	//   - Validate the ClientMetadata value.
	//   - Encrypt the ClientMetadata value. Don't use Amazon Cognito to provide
	//   sensitive information.
	ClientMetadata map[string]string

	noSmithyDocumentSerde
}

// Represents the response from the server for the request to update user
// attributes.
type UpdateUserAttributesOutput struct {

	// The code delivery details list from the server for the request to update user
	// attributes.
	CodeDeliveryDetailsList []types.CodeDeliveryDetailsType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateUserAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateUserAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateUserAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateUserAttributes"); err != nil {
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
	if err = addOpUpdateUserAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateUserAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateUserAttributes",
	}
}
