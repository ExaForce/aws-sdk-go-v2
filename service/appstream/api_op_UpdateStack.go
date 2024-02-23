// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified fields for the specified stack.
func (c *Client) UpdateStack(ctx context.Context, params *UpdateStackInput, optFns ...func(*Options)) (*UpdateStackOutput, error) {
	if params == nil {
		params = &UpdateStackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStack", params, optFns, c.addOperationUpdateStackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateStackInput struct {

	// The name of the stack.
	//
	// This member is required.
	Name *string

	// The list of interface VPC endpoint (interface endpoint) objects. Users of the
	// stack can connect to AppStream 2.0 only through the specified endpoints.
	AccessEndpoints []types.AccessEndpoint

	// The persistent application settings for users of a stack. When these settings
	// are enabled, changes that users make to applications and Windows settings are
	// automatically saved after each session and applied to the next session.
	ApplicationSettings *types.ApplicationSettings

	// The stack attributes to delete.
	AttributesToDelete []types.StackAttribute

	// Deletes the storage connectors currently enabled for the stack.
	//
	// Deprecated: This member has been deprecated.
	DeleteStorageConnectors *bool

	// The description to display.
	Description *string

	// The stack name to display.
	DisplayName *string

	// The domains where AppStream 2.0 streaming sessions can be embedded in an
	// iframe. You must approve the domains that you want to host embedded AppStream
	// 2.0 streaming sessions.
	EmbedHostDomains []string

	// The URL that users are redirected to after they choose the Send Feedback link.
	// If no URL is specified, no Send Feedback link is displayed.
	FeedbackURL *string

	// The URL that users are redirected to after their streaming session ends.
	RedirectURL *string

	// The storage connectors to enable.
	StorageConnectors []types.StorageConnector

	// The streaming protocol you want your stack to prefer. This can be UDP or TCP.
	// Currently, UDP is only supported in the Windows native client.
	StreamingExperienceSettings *types.StreamingExperienceSettings

	// The actions that are enabled or disabled for users during their streaming
	// sessions. By default, these actions are enabled.
	UserSettings []types.UserSetting

	noSmithyDocumentSerde
}

type UpdateStackOutput struct {

	// Information about the stack.
	Stack *types.Stack

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateStackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateStack{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateStack{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateStack"); err != nil {
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
	if err = addOpUpdateStackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStack(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateStack(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateStack",
	}
}
