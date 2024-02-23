// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets free trial status for multiple Amazon Web Services accounts.
func (c *Client) BatchGetFreeTrialInfo(ctx context.Context, params *BatchGetFreeTrialInfoInput, optFns ...func(*Options)) (*BatchGetFreeTrialInfoOutput, error) {
	if params == nil {
		params = &BatchGetFreeTrialInfoInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetFreeTrialInfo", params, optFns, c.addOperationBatchGetFreeTrialInfoMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetFreeTrialInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetFreeTrialInfoInput struct {

	// The account IDs to get free trial status for.
	//
	// This member is required.
	AccountIds []string

	noSmithyDocumentSerde
}

type BatchGetFreeTrialInfoOutput struct {

	// An array of objects that provide Amazon Inspector free trial details for each
	// of the requested accounts.
	//
	// This member is required.
	Accounts []types.FreeTrialAccountInfo

	// An array of objects detailing any accounts that free trial data could not be
	// returned for.
	//
	// This member is required.
	FailedAccounts []types.FreeTrialInfoError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetFreeTrialInfoMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetFreeTrialInfo{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetFreeTrialInfo{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetFreeTrialInfo"); err != nil {
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
	if err = addOpBatchGetFreeTrialInfoValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetFreeTrialInfo(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetFreeTrialInfo(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetFreeTrialInfo",
	}
}
