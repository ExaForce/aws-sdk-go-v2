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

// Gets a particular detector version.
func (c *Client) GetDetectorVersion(ctx context.Context, params *GetDetectorVersionInput, optFns ...func(*Options)) (*GetDetectorVersionOutput, error) {
	if params == nil {
		params = &GetDetectorVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDetectorVersion", params, optFns, c.addOperationGetDetectorVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDetectorVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDetectorVersionInput struct {

	// The detector ID.
	//
	// This member is required.
	DetectorId *string

	// The detector version ID.
	//
	// This member is required.
	DetectorVersionId *string

	noSmithyDocumentSerde
}

type GetDetectorVersionOutput struct {

	// The detector version ARN.
	Arn *string

	// The timestamp when the detector version was created.
	CreatedTime *string

	// The detector version description.
	Description *string

	// The detector ID.
	DetectorId *string

	// The detector version ID.
	DetectorVersionId *string

	// The Amazon SageMaker model endpoints included in the detector version.
	ExternalModelEndpoints []string

	// The timestamp when the detector version was last updated.
	LastUpdatedTime *string

	// The model versions included in the detector version.
	ModelVersions []types.ModelVersion

	// The execution mode of the rule in the dectector FIRST_MATCHED indicates that
	// Amazon Fraud Detector evaluates rules sequentially, first to last, stopping at
	// the first matched rule. Amazon Fraud dectector then provides the outcomes for
	// that single rule. ALL_MATCHED indicates that Amazon Fraud Detector evaluates
	// all rules and returns the outcomes for all matched rules. You can define and
	// edit the rule mode at the detector version level, when it is in draft status.
	RuleExecutionMode types.RuleExecutionMode

	// The rules included in the detector version.
	Rules []types.Rule

	// The status of the detector version.
	Status types.DetectorVersionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDetectorVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDetectorVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDetectorVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDetectorVersion"); err != nil {
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
	if err = addOpGetDetectorVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDetectorVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDetectorVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDetectorVersion",
	}
}
