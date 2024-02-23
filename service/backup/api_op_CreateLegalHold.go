// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This action creates a legal hold on a recovery point (backup). A legal hold is
// a restraint on altering or deleting a backup until an authorized user cancels
// the legal hold. Any actions to delete or disassociate a recovery point will fail
// with an error if one or more active legal holds are on the recovery point.
func (c *Client) CreateLegalHold(ctx context.Context, params *CreateLegalHoldInput, optFns ...func(*Options)) (*CreateLegalHoldOutput, error) {
	if params == nil {
		params = &CreateLegalHoldInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLegalHold", params, optFns, c.addOperationCreateLegalHoldMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLegalHoldOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLegalHoldInput struct {

	// This is the string description of the legal hold.
	//
	// This member is required.
	Description *string

	// This is the string title of the legal hold.
	//
	// This member is required.
	Title *string

	// This is a user-chosen string used to distinguish between otherwise identical
	// calls. Retrying a successful request with the same idempotency token results in
	// a success message with no action taken.
	IdempotencyToken *string

	// This specifies criteria to assign a set of resources, such as resource types or
	// backup vaults.
	RecoveryPointSelection *types.RecoveryPointSelection

	// Optional tags to include. A tag is a key-value pair you can use to manage,
	// filter, and search for your resources. Allowed characters include UTF-8 letters,
	// numbers, spaces, and the following characters: + - = . _ : /.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateLegalHoldOutput struct {

	// Time in number format when legal hold was created.
	CreationDate *time.Time

	// This is the returned string description of the legal hold.
	Description *string

	// This is the ARN (Amazon Resource Number) of the created legal hold.
	LegalHoldArn *string

	// Legal hold ID returned for the specified legal hold on a recovery point.
	LegalHoldId *string

	// This specifies criteria to assign a set of resources, such as resource types or
	// backup vaults.
	RecoveryPointSelection *types.RecoveryPointSelection

	// This displays the status of the legal hold returned after creating the legal
	// hold. Statuses can be ACTIVE , PENDING , CANCELED , CANCELING , or FAILED .
	Status types.LegalHoldStatus

	// This is the string title of the legal hold returned after creating the legal
	// hold.
	Title *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLegalHoldMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLegalHold{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLegalHold{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLegalHold"); err != nil {
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
	if err = addOpCreateLegalHoldValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLegalHold(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLegalHold(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLegalHold",
	}
}
