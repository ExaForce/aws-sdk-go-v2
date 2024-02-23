// Code generated by smithy-go-codegen DO NOT EDIT.

package iottwinmaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a scene.
func (c *Client) GetScene(ctx context.Context, params *GetSceneInput, optFns ...func(*Options)) (*GetSceneOutput, error) {
	if params == nil {
		params = &GetSceneInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetScene", params, optFns, c.addOperationGetSceneMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSceneOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSceneInput struct {

	// The ID of the scene.
	//
	// This member is required.
	SceneId *string

	// The ID of the workspace that contains the scene.
	//
	// This member is required.
	WorkspaceId *string

	noSmithyDocumentSerde
}

type GetSceneOutput struct {

	// The ARN of the scene.
	//
	// This member is required.
	Arn *string

	// The relative path that specifies the location of the content definition file.
	//
	// This member is required.
	ContentLocation *string

	// The date and time when the scene was created.
	//
	// This member is required.
	CreationDateTime *time.Time

	// The ID of the scene.
	//
	// This member is required.
	SceneId *string

	// The date and time when the scene was last updated.
	//
	// This member is required.
	UpdateDateTime *time.Time

	// The ID of the workspace that contains the scene.
	//
	// This member is required.
	WorkspaceId *string

	// A list of capabilities that the scene uses to render.
	Capabilities []string

	// The description of the scene.
	Description *string

	// The SceneResponse error.
	Error *types.SceneError

	// The generated scene metadata.
	GeneratedSceneMetadata map[string]string

	// The response metadata.
	SceneMetadata map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSceneMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetScene{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetScene{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetScene"); err != nil {
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
	if err = addEndpointPrefix_opGetSceneMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetSceneValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetScene(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetSceneMiddleware struct {
}

func (*endpointPrefix_opGetSceneMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetSceneMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetSceneMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetSceneMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetScene(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetScene",
	}
}
