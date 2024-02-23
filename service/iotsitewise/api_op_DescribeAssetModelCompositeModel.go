// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about an asset model composite model (also known as an
// asset model component). For more information, see Custom composite models
// (Components) (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/custom-composite-models.html)
// in the IoT SiteWise User Guide.
func (c *Client) DescribeAssetModelCompositeModel(ctx context.Context, params *DescribeAssetModelCompositeModelInput, optFns ...func(*Options)) (*DescribeAssetModelCompositeModelOutput, error) {
	if params == nil {
		params = &DescribeAssetModelCompositeModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAssetModelCompositeModel", params, optFns, c.addOperationDescribeAssetModelCompositeModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAssetModelCompositeModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAssetModelCompositeModelInput struct {

	// The ID of a composite model on this asset model. This can be either the actual
	// ID in UUID format, or else externalId: followed by the external ID, if it has
	// one. For more information, see Referencing objects with external IDs (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-id-references)
	// in the IoT SiteWise User Guide.
	//
	// This member is required.
	AssetModelCompositeModelId *string

	// The ID of the asset model. This can be either the actual ID in UUID format, or
	// else externalId: followed by the external ID, if it has one. For more
	// information, see Referencing objects with external IDs (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/object-ids.html#external-id-references)
	// in the IoT SiteWise User Guide.
	//
	// This member is required.
	AssetModelId *string

	noSmithyDocumentSerde
}

type DescribeAssetModelCompositeModelOutput struct {

	// The description for the composite model.
	//
	// This member is required.
	AssetModelCompositeModelDescription *string

	// The ID of a composite model on this asset model.
	//
	// This member is required.
	AssetModelCompositeModelId *string

	// The unique, friendly name for the composite model.
	//
	// This member is required.
	AssetModelCompositeModelName *string

	// The path to the composite model listing the parent composite models.
	//
	// This member is required.
	AssetModelCompositeModelPath []types.AssetModelCompositeModelPathSegment

	// The property definitions of the composite model.
	//
	// This member is required.
	AssetModelCompositeModelProperties []types.AssetModelProperty

	// The list of composite model summaries for the composite model.
	//
	// This member is required.
	AssetModelCompositeModelSummaries []types.AssetModelCompositeModelSummary

	// The composite model type. Valid values are AWS/ALARM , CUSTOM , or
	// AWS/L4E_ANOMALY .
	//
	// This member is required.
	AssetModelCompositeModelType *string

	// The ID of the asset model, in UUID format.
	//
	// This member is required.
	AssetModelId *string

	// The available actions for a composite model on this asset model.
	ActionDefinitions []types.ActionDefinition

	// The external ID of a composite model on this asset model.
	AssetModelCompositeModelExternalId *string

	// Metadata for the composition relationship established by using
	// composedAssetModelId in CreateAssetModelCompositeModel (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_CreateAssetModelCompositeModel.html)
	// . For instance, an array detailing the path of the composition relationship for
	// this composite model.
	CompositionDetails *types.CompositionDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAssetModelCompositeModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAssetModelCompositeModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAssetModelCompositeModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAssetModelCompositeModel"); err != nil {
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
	if err = addEndpointPrefix_opDescribeAssetModelCompositeModelMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeAssetModelCompositeModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAssetModelCompositeModel(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeAssetModelCompositeModelMiddleware struct {
}

func (*endpointPrefix_opDescribeAssetModelCompositeModelMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeAssetModelCompositeModelMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
func addEndpointPrefix_opDescribeAssetModelCompositeModelMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDescribeAssetModelCompositeModelMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAssetModelCompositeModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAssetModelCompositeModel",
	}
}
