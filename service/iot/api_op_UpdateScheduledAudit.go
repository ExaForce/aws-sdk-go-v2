// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a scheduled audit, including which checks are performed and how often
// the audit takes place. Requires permission to access the UpdateScheduledAudit (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) UpdateScheduledAudit(ctx context.Context, params *UpdateScheduledAuditInput, optFns ...func(*Options)) (*UpdateScheduledAuditOutput, error) {
	if params == nil {
		params = &UpdateScheduledAuditInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateScheduledAudit", params, optFns, c.addOperationUpdateScheduledAuditMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateScheduledAuditOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateScheduledAuditInput struct {

	// The name of the scheduled audit. (Max. 128 chars)
	//
	// This member is required.
	ScheduledAuditName *string

	// The day of the month on which the scheduled audit takes place. This can be 1
	// through 31 or LAST . This field is required if the frequency parameter is set
	// to MONTHLY . If days 29-31 are specified, and the month does not have that many
	// days, the audit takes place on the "LAST" day of the month.
	DayOfMonth *string

	// The day of the week on which the scheduled audit takes place. This can be one
	// of SUN , MON , TUE , WED , THU , FRI , or SAT . This field is required if the
	// "frequency" parameter is set to WEEKLY or BIWEEKLY .
	DayOfWeek types.DayOfWeek

	// How often the scheduled audit takes place, either DAILY , WEEKLY , BIWEEKLY , or
	// MONTHLY . The start time of each audit is determined by the system.
	Frequency types.AuditFrequency

	// Which checks are performed during the scheduled audit. Checks must be enabled
	// for your account. (Use DescribeAccountAuditConfiguration to see the list of all
	// checks, including those that are enabled or use UpdateAccountAuditConfiguration
	// to select which checks are enabled.)
	TargetCheckNames []string

	noSmithyDocumentSerde
}

type UpdateScheduledAuditOutput struct {

	// The ARN of the scheduled audit.
	ScheduledAuditArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateScheduledAuditMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateScheduledAudit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateScheduledAudit{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateScheduledAudit"); err != nil {
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
	if err = addOpUpdateScheduledAuditValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateScheduledAudit(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateScheduledAudit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateScheduledAudit",
	}
}
