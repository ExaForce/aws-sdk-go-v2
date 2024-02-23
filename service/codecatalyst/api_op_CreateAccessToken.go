// Code generated by smithy-go-codegen DO NOT EDIT.

package codecatalyst

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a personal access token (PAT) for the current user. A personal access
// token (PAT) is similar to a password. It is associated with your user identity
// for use across all spaces and projects in Amazon CodeCatalyst. You use PATs to
// access CodeCatalyst from resources that include integrated development
// environments (IDEs) and Git-based source repositories. PATs represent you in
// Amazon CodeCatalyst and you can manage them in your user settings.For more
// information, see Managing personal access tokens in Amazon CodeCatalyst (https://docs.aws.amazon.com/codecatalyst/latest/userguide/ipa-tokens-keys.html)
// .
func (c *Client) CreateAccessToken(ctx context.Context, params *CreateAccessTokenInput, optFns ...func(*Options)) (*CreateAccessTokenOutput, error) {
	if params == nil {
		params = &CreateAccessTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccessToken", params, optFns, c.addOperationCreateAccessTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccessTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessTokenInput struct {

	// The friendly name of the personal access token.
	//
	// This member is required.
	Name *string

	// The date and time the personal access token expires, in coordinated universal
	// time (UTC) timestamp format as specified in RFC 3339 (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
	// .
	ExpiresTime *time.Time

	noSmithyDocumentSerde
}

type CreateAccessTokenOutput struct {

	// The system-generated unique ID of the access token.
	//
	// This member is required.
	AccessTokenId *string

	// The date and time the personal access token expires, in coordinated universal
	// time (UTC) timestamp format as specified in RFC 3339 (https://www.rfc-editor.org/rfc/rfc3339#section-5.6)
	// . If not specified, the default is one year from creation.
	//
	// This member is required.
	ExpiresTime *time.Time

	// The friendly name of the personal access token.
	//
	// This member is required.
	Name *string

	// The secret value of the personal access token.
	//
	// This member is required.
	Secret *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAccessTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAccessToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAccessToken{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAccessToken"); err != nil {
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
	if err = addOpCreateAccessTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAccessToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAccessToken",
	}
}
