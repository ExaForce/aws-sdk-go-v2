// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get a farm.
func (c *Client) GetFarm(ctx context.Context, params *GetFarmInput, optFns ...func(*Options)) (*GetFarmOutput, error) {
	if params == nil {
		params = &GetFarmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFarm", params, optFns, c.addOperationGetFarmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFarmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFarmInput struct {

	// The farm ID of the farm.
	//
	// This member is required.
	FarmId *string

	noSmithyDocumentSerde
}

type GetFarmOutput struct {

	// The date and time the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The user or system that created this resource.
	//
	// This member is required.
	CreatedBy *string

	// The display name of the farm.
	//
	// This field can store any content. Escape or encode this content before
	// displaying it on a webpage or any other system that might interpret the content
	// of this field.
	//
	// This member is required.
	DisplayName *string

	// The farm ID of the farm to get.
	//
	// This member is required.
	FarmId *string

	// The ARN of the KMS key used on the farm.
	//
	// This member is required.
	KmsKeyArn *string

	// The description of the farm.
	//
	// This field can store any content. Escape or encode this content before
	// displaying it on a webpage or any other system that might interpret the content
	// of this field.
	Description *string

	// The date and time the resource was updated.
	UpdatedAt *time.Time

	// The user or system that updated this resource.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFarmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFarm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFarm{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetFarm"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGetFarmMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetFarmValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFarm(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetFarmMiddleware struct {
}

func (*endpointPrefix_opGetFarmMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetFarmMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
func addEndpointPrefix_opGetFarmMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetFarmMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetFarm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetFarm",
	}
}
