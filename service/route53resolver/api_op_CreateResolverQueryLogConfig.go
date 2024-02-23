// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Resolver query logging configuration, which defines where you want
// Resolver to save DNS query logs that originate in your VPCs. Resolver can log
// queries only for VPCs that are in the same Region as the query logging
// configuration. To specify which VPCs you want to log queries for, you use
// AssociateResolverQueryLogConfig . For more information, see
// AssociateResolverQueryLogConfig (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_AssociateResolverQueryLogConfig.html)
// . You can optionally use Resource Access Manager (RAM) to share a query logging
// configuration with other Amazon Web Services accounts. The other accounts can
// then associate VPCs with the configuration. The query logs that Resolver creates
// for a configuration include all DNS queries that originate in all VPCs that are
// associated with the configuration.
func (c *Client) CreateResolverQueryLogConfig(ctx context.Context, params *CreateResolverQueryLogConfigInput, optFns ...func(*Options)) (*CreateResolverQueryLogConfigOutput, error) {
	if params == nil {
		params = &CreateResolverQueryLogConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResolverQueryLogConfig", params, optFns, c.addOperationCreateResolverQueryLogConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResolverQueryLogConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResolverQueryLogConfigInput struct {

	// A unique string that identifies the request and that allows failed requests to
	// be retried without the risk of running the operation twice. CreatorRequestId
	// can be any unique string, for example, a date/time stamp.
	//
	// This member is required.
	CreatorRequestId *string

	// The ARN of the resource that you want Resolver to send query logs. You can send
	// query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data
	// Firehose delivery stream. Examples of valid values include the following:
	//   - S3 bucket: arn:aws:s3:::examplebucket You can optionally append a file
	//   prefix to the end of the ARN. arn:aws:s3:::examplebucket/development/
	//   - CloudWatch Logs log group:
	//   arn:aws:logs:us-west-1:123456789012:log-group:/mystack-testgroup-12ABC1AB12A1:*
	//   - Kinesis Data Firehose delivery stream:
	//   arn:aws:kinesis:us-east-2:0123456789:stream/my_stream_name
	//
	// This member is required.
	DestinationArn *string

	// The name that you want to give the query logging configuration.
	//
	// This member is required.
	Name *string

	// A list of the tag keys and values that you want to associate with the query
	// logging configuration.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateResolverQueryLogConfigOutput struct {

	// Information about the CreateResolverQueryLogConfig request, including the
	// status of the request.
	ResolverQueryLogConfig *types.ResolverQueryLogConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResolverQueryLogConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateResolverQueryLogConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateResolverQueryLogConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateResolverQueryLogConfig"); err != nil {
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
	if err = addIdempotencyToken_opCreateResolverQueryLogConfigMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateResolverQueryLogConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResolverQueryLogConfig(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateResolverQueryLogConfig struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateResolverQueryLogConfig) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateResolverQueryLogConfig) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateResolverQueryLogConfigInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateResolverQueryLogConfigInput ")
	}

	if input.CreatorRequestId == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.CreatorRequestId = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateResolverQueryLogConfigMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateResolverQueryLogConfig{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateResolverQueryLogConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateResolverQueryLogConfig",
	}
}
