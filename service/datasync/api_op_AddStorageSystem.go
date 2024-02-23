// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Web Services resource for an on-premises storage system that
// you want DataSync Discovery to collect information about.
func (c *Client) AddStorageSystem(ctx context.Context, params *AddStorageSystemInput, optFns ...func(*Options)) (*AddStorageSystemOutput, error) {
	if params == nil {
		params = &AddStorageSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddStorageSystem", params, optFns, c.addOperationAddStorageSystemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddStorageSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddStorageSystemInput struct {

	// Specifies the Amazon Resource Name (ARN) of the DataSync agent that connects to
	// and reads from your on-premises storage system's management interface. You can
	// only specify one ARN.
	//
	// This member is required.
	AgentArns []string

	// Specifies a client token to make sure requests with this API operation are
	// idempotent. If you don't specify a client token, DataSync generates one for you
	// automatically.
	//
	// This member is required.
	ClientToken *string

	// Specifies the user name and password for accessing your on-premises storage
	// system's management interface.
	//
	// This member is required.
	Credentials *types.Credentials

	// Specifies the server name and network port required to connect with the
	// management interface of your on-premises storage system.
	//
	// This member is required.
	ServerConfiguration *types.DiscoveryServerConfiguration

	// Specifies the type of on-premises storage system that you want DataSync
	// Discovery to collect information about. DataSync Discovery currently supports
	// NetApp Fabric-Attached Storage (FAS) and All Flash FAS (AFF) systems running
	// ONTAP 9.7 or later.
	//
	// This member is required.
	SystemType types.DiscoverySystemType

	// Specifies the ARN of the Amazon CloudWatch log group for monitoring and logging
	// discovery job events.
	CloudWatchLogGroupArn *string

	// Specifies a familiar name for your on-premises storage system.
	Name *string

	// Specifies labels that help you categorize, filter, and search for your Amazon
	// Web Services resources. We recommend creating at least a name tag for your
	// on-premises storage system.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

type AddStorageSystemOutput struct {

	// The ARN of the on-premises storage system that you can use with DataSync
	// Discovery.
	//
	// This member is required.
	StorageSystemArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddStorageSystemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddStorageSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddStorageSystem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddStorageSystem"); err != nil {
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
	if err = addEndpointPrefix_opAddStorageSystemMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opAddStorageSystemMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpAddStorageSystemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddStorageSystem(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opAddStorageSystemMiddleware struct {
}

func (*endpointPrefix_opAddStorageSystemMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opAddStorageSystemMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "discovery-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opAddStorageSystemMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opAddStorageSystemMiddleware{}, "ResolveEndpointV2", middleware.After)
}

type idempotencyToken_initializeOpAddStorageSystem struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpAddStorageSystem) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpAddStorageSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*AddStorageSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *AddStorageSystemInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opAddStorageSystemMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpAddStorageSystem{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opAddStorageSystem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddStorageSystem",
	}
}
