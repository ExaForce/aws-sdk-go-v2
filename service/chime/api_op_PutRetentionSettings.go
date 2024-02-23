// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Puts retention settings for the specified Amazon Chime Enterprise account. We
// recommend using AWS CloudTrail to monitor usage of this API for your account.
// For more information, see Logging Amazon Chime API Calls with AWS CloudTrail (https://docs.aws.amazon.com/chime/latest/ag/cloudtrail.html)
// in the Amazon Chime Administration Guide. To turn off existing retention
// settings, remove the number of days from the corresponding RetentionDays field
// in the RetentionSettings object. For more information about retention settings,
// see Managing Chat Retention Policies (https://docs.aws.amazon.com/chime/latest/ag/chat-retention.html)
// in the Amazon Chime Administration Guide.
func (c *Client) PutRetentionSettings(ctx context.Context, params *PutRetentionSettingsInput, optFns ...func(*Options)) (*PutRetentionSettingsOutput, error) {
	if params == nil {
		params = &PutRetentionSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRetentionSettings", params, optFns, c.addOperationPutRetentionSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRetentionSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRetentionSettingsInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The retention settings.
	//
	// This member is required.
	RetentionSettings *types.RetentionSettings

	noSmithyDocumentSerde
}

type PutRetentionSettingsOutput struct {

	// The timestamp representing the time at which the specified items are
	// permanently deleted, in ISO 8601 format.
	InitiateDeletionTimestamp *time.Time

	// The retention settings.
	RetentionSettings *types.RetentionSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRetentionSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutRetentionSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutRetentionSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutRetentionSettings"); err != nil {
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
	if err = addOpPutRetentionSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRetentionSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRetentionSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutRetentionSettings",
	}
}
