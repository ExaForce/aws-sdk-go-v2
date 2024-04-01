// Code generated by smithy-go-codegen DO NOT EDIT.

package deadline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/deadline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get a budget.
func (c *Client) GetBudget(ctx context.Context, params *GetBudgetInput, optFns ...func(*Options)) (*GetBudgetOutput, error) {
	if params == nil {
		params = &GetBudgetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBudget", params, optFns, c.addOperationGetBudgetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBudgetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBudgetInput struct {

	// The budget ID.
	//
	// This member is required.
	BudgetId *string

	// The farm ID of the farm connected to the budget.
	//
	// This member is required.
	FarmId *string

	noSmithyDocumentSerde
}

type GetBudgetOutput struct {

	// The budget actions for the budget.
	//
	// This member is required.
	Actions []types.ResponseBudgetAction

	// The consumed usage limit for the budget.
	//
	// This member is required.
	ApproximateDollarLimit *float32

	// The budget ID.
	//
	// This member is required.
	BudgetId *string

	// The date and time the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The user or system that created this resource.
	//
	// This member is required.
	CreatedBy *string

	// The display name of the budget.
	//
	// This member is required.
	DisplayName *string

	// The budget schedule.
	//
	// This member is required.
	Schedule types.BudgetSchedule

	// The status of the budget.
	//   - ACTIVE –Get a budget being evaluated.
	//   - INACTIVE –Get an inactive budget. This can include expired, canceled, or
	//   deleted statuses.
	//
	// This member is required.
	Status types.BudgetStatus

	// The resource that the budget is tracking usage for.
	//
	// This member is required.
	UsageTrackingResource types.UsageTrackingResource

	// The usages of the budget.
	//
	// This member is required.
	Usages *types.ConsumedUsages

	// The description of the budget.
	Description *string

	// The date and time the queue stopped.
	QueueStoppedAt *time.Time

	// The date and time the resource was updated.
	UpdatedAt *time.Time

	// The user or system that updated this resource.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBudgetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBudget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBudget{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetBudget"); err != nil {
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
	if err = addEndpointPrefix_opGetBudgetMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetBudgetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBudget(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetBudgetMiddleware struct {
}

func (*endpointPrefix_opGetBudgetMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetBudgetMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
func addEndpointPrefix_opGetBudgetMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetBudgetMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetBudget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetBudget",
	}
}
