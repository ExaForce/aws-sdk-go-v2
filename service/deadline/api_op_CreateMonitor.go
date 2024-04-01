// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Web Services Deadline Cloud monitor that you can use to view
// your farms, queues, and fleets. After you submit a job, you can track the
// progress of the tasks and steps that make up the job, and then download the
// job's results.
func (c *Client) CreateMonitor(ctx context.Context, params *CreateMonitorInput, optFns ...func(*Options)) (*CreateMonitorOutput, error) {
	if params == nil {
		params = &CreateMonitorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMonitor", params, optFns, c.addOperationCreateMonitorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMonitorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMonitorInput struct {

	// The name that you give the monitor that is displayed in the Deadline Cloud
	// console.
	//
	// This member is required.
	DisplayName *string

	// The Amazon Resource Name (ARN) of the IAM Identity Center instance that
	// authenticates monitor users.
	//
	// This member is required.
	IdentityCenterInstanceArn *string

	// The Amazon Resource Name (ARN) of the IAM role that the monitor uses to connect
	// to Deadline Cloud. Every user that signs in to the monitor using IAM Identity
	// Center uses this role to access Deadline Cloud resources.
	//
	// This member is required.
	RoleArn *string

	// The subdomain to use when creating the monitor URL. The full URL of the monitor
	// is subdomain.Region.deadlinecloud.amazonaws.com.
	//
	// This member is required.
	Subdomain *string

	// The unique token which the server uses to recognize retries of the same request.
	ClientToken *string

	noSmithyDocumentSerde
}

type CreateMonitorOutput struct {

	// The Amazon Resource Name (ARN) that IAM Identity Center assigns to the monitor.
	//
	// This member is required.
	IdentityCenterApplicationArn *string

	// The unique identifier of the monitor.
	//
	// This member is required.
	MonitorId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMonitorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateMonitor"); err != nil {
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
	if err = addEndpointPrefix_opCreateMonitorMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateMonitorMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMonitorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMonitor(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opCreateMonitorMiddleware struct {
}

func (*endpointPrefix_opCreateMonitorMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateMonitorMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "management." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opCreateMonitorMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opCreateMonitorMiddleware{}, "ResolveEndpointV2", middleware.After)
}

type idempotencyToken_initializeOpCreateMonitor struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMonitor) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMonitor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMonitorInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMonitorInput ")
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
func addIdempotencyToken_opCreateMonitorMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMonitor{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMonitor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateMonitor",
	}
}
