// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The GetQualificationScore operation returns the value of a Worker's
// Qualification for a given Qualification type. To get a Worker's Qualification,
// you must know the Worker's ID. The Worker's ID is included in the assignment
// data returned by the ListAssignmentsForHIT operation. Only the owner of a
// Qualification type can query the value of a Worker's Qualification of that type.
func (c *Client) GetQualificationScore(ctx context.Context, params *GetQualificationScoreInput, optFns ...func(*Options)) (*GetQualificationScoreOutput, error) {
	if params == nil {
		params = &GetQualificationScoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQualificationScore", params, optFns, c.addOperationGetQualificationScoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQualificationScoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQualificationScoreInput struct {

	// The ID of the QualificationType.
	//
	// This member is required.
	QualificationTypeId *string

	// The ID of the Worker whose Qualification is being updated.
	//
	// This member is required.
	WorkerId *string

	noSmithyDocumentSerde
}

type GetQualificationScoreOutput struct {

	// The Qualification data structure of the Qualification assigned to a user,
	// including the Qualification type and the value (score).
	Qualification *types.Qualification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetQualificationScoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetQualificationScore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetQualificationScore{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetQualificationScore"); err != nil {
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
	if err = addOpGetQualificationScoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQualificationScore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetQualificationScore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetQualificationScore",
	}
}
