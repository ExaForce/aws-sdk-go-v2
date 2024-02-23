// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets details of the past fraud predictions for the specified event ID, event
// type, detector ID, and detector version ID that was generated in the specified
// time period.
func (c *Client) GetEventPredictionMetadata(ctx context.Context, params *GetEventPredictionMetadataInput, optFns ...func(*Options)) (*GetEventPredictionMetadataOutput, error) {
	if params == nil {
		params = &GetEventPredictionMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEventPredictionMetadata", params, optFns, c.addOperationGetEventPredictionMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEventPredictionMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEventPredictionMetadataInput struct {

	// The detector ID.
	//
	// This member is required.
	DetectorId *string

	// The detector version ID.
	//
	// This member is required.
	DetectorVersionId *string

	// The event ID.
	//
	// This member is required.
	EventId *string

	// The event type associated with the detector specified for the prediction.
	//
	// This member is required.
	EventTypeName *string

	// The timestamp that defines when the prediction was generated. The timestamp
	// must be specified using ISO 8601 standard in UTC. We recommend calling
	// ListEventPredictions (https://docs.aws.amazon.com/frauddetector/latest/api/API_ListEventPredictions.html)
	// first, and using the predictionTimestamp value in the response to provide an
	// accurate prediction timestamp value.
	//
	// This member is required.
	PredictionTimestamp *string

	noSmithyDocumentSerde
}

type GetEventPredictionMetadataOutput struct {

	// The detector ID.
	DetectorId *string

	// The detector version ID.
	DetectorVersionId *string

	// The status of the detector version.
	DetectorVersionStatus *string

	// The entity ID.
	EntityId *string

	// The entity type.
	EntityType *string

	// External (Amazon SageMaker) models that were evaluated for generating
	// predictions.
	EvaluatedExternalModels []types.EvaluatedExternalModel

	// Model versions that were evaluated for generating predictions.
	EvaluatedModelVersions []types.EvaluatedModelVersion

	// The event ID.
	EventId *string

	// The timestamp for when the prediction was generated for the associated event ID.
	EventTimestamp *string

	// The event type associated with the detector specified for this prediction.
	EventTypeName *string

	// A list of event variables that influenced the prediction scores.
	EventVariables []types.EventVariableSummary

	// The outcomes of the matched rule, based on the rule execution mode.
	Outcomes []string

	// The timestamp that defines when the prediction was generated.
	PredictionTimestamp *string

	// The execution mode of the rule used for evaluating variable values.
	RuleExecutionMode types.RuleExecutionMode

	// List of rules associated with the detector version that were used for
	// evaluating variable values.
	Rules []types.EvaluatedRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEventPredictionMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetEventPredictionMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetEventPredictionMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEventPredictionMetadata"); err != nil {
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
	if err = addOpGetEventPredictionMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEventPredictionMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEventPredictionMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEventPredictionMetadata",
	}
}
