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

// This action is no longer supported. You can use it to configure only SMS MFA.
// You can't use it to configure time-based one-time password (TOTP) software token
// MFA. To configure either type of MFA, use SetUserMFAPreference (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_SetUserMFAPreference.html)
// instead. Authorize this action with a signed-in user's access token. It must
// include the scope aws.cognito.signin.user.admin . Amazon Cognito doesn't
// evaluate Identity and Access Management (IAM) policies in requests for this API
// operation. For this operation, you can't use IAM credentials to authorize
// requests, and you can't grant IAM permissions in policies. For more information
// about authorization models in Amazon Cognito, see Using the Amazon Cognito user
// pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
// .
func (c *Client) SetUserSettings(ctx context.Context, params *SetUserSettingsInput, optFns ...func(*Options)) (*SetUserSettingsOutput, error) {
	if params == nil {
		params = &SetUserSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetUserSettings", params, optFns, c.addOperationSetUserSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetUserSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to set user settings.
type SetUserSettingsInput struct {

	// A valid access token that Amazon Cognito issued to the user whose user settings
	// you want to configure.
	//
	// This member is required.
	AccessToken *string

	// You can use this parameter only to set an SMS configuration that uses SMS for
	// delivery.
	//
	// This member is required.
	MFAOptions []types.MFAOptionType

	noSmithyDocumentSerde
}

// The response from the server for a set user settings request.
type SetUserSettingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetUserSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetUserSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetUserSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetUserSettings"); err != nil {
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
	if err = addOpSetUserSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetUserSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetUserSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetUserSettings",
	}
}
