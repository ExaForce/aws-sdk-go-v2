// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a ledger, including its state, permissions mode,
// encryption at rest settings, and when it was created.
func (c *Client) DescribeLedger(ctx context.Context, params *DescribeLedgerInput, optFns ...func(*Options)) (*DescribeLedgerOutput, error) {
	if params == nil {
		params = &DescribeLedgerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLedger", params, optFns, c.addOperationDescribeLedgerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLedgerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLedgerInput struct {

	// The name of the ledger that you want to describe.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DescribeLedgerOutput struct {

	// The Amazon Resource Name (ARN) for the ledger.
	Arn *string

	// The date and time, in epoch time format, when the ledger was created. (Epoch
	// time format is the number of seconds elapsed since 12:00:00 AM January 1, 1970
	// UTC.)
	CreationDateTime *time.Time

	// Specifies whether the ledger is protected from being deleted by any user. If
	// not defined during ledger creation, this feature is enabled ( true ) by default.
	// If deletion protection is enabled, you must first disable it before you can
	// delete the ledger. You can disable it by calling the UpdateLedger operation to
	// set this parameter to false .
	DeletionProtection *bool

	// Information about the encryption of data at rest in the ledger. This includes
	// the current status, the KMS key, and when the key became inaccessible (in the
	// case of an error).
	EncryptionDescription *types.LedgerEncryptionDescription

	// The name of the ledger.
	Name *string

	// The permissions mode of the ledger.
	PermissionsMode types.PermissionsMode

	// The current status of the ledger.
	State types.LedgerState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLedgerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeLedger{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeLedger{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeLedger"); err != nil {
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
	if err = addOpDescribeLedgerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLedger(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLedger(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeLedger",
	}
}
