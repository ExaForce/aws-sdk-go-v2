// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Submits a request to process a query statement. This operation generates work
// units that can be retrieved with the GetWorkUnits operation as soon as the
// query state is WORKUNITS_AVAILABLE or FINISHED.
func (c *Client) StartQueryPlanning(ctx context.Context, params *StartQueryPlanningInput, optFns ...func(*Options)) (*StartQueryPlanningOutput, error) {
	if params == nil {
		params = &StartQueryPlanningInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartQueryPlanning", params, optFns, c.addOperationStartQueryPlanningMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartQueryPlanningOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartQueryPlanningInput struct {

	// A structure containing information about the query plan.
	//
	// This member is required.
	QueryPlanningContext *types.QueryPlanningContext

	// A PartiQL query statement used as an input to the planner service.
	//
	// This member is required.
	QueryString *string

	noSmithyDocumentSerde
}

// A structure for the output.
type StartQueryPlanningOutput struct {

	// The ID of the plan query operation can be used to fetch the actual work unit
	// descriptors that are produced as the result of the operation. The ID is also
	// used to get the query state and as an input to the Execute operation.
	//
	// This member is required.
	QueryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartQueryPlanningMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartQueryPlanning{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartQueryPlanning{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartQueryPlanning"); err != nil {
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
	if err = addEndpointPrefix_opStartQueryPlanningMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStartQueryPlanningValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartQueryPlanning(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opStartQueryPlanningMiddleware struct {
}

func (*endpointPrefix_opStartQueryPlanningMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opStartQueryPlanningMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "query-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opStartQueryPlanningMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opStartQueryPlanningMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opStartQueryPlanning(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartQueryPlanning",
	}
}
