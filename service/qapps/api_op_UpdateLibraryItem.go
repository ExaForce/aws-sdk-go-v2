// Code generated by smithy-go-codegen DO NOT EDIT.

package qapps

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/qapps/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates the metadata and status of a library item for an Amazon Q App.
func (c *Client) UpdateLibraryItem(ctx context.Context, params *UpdateLibraryItemInput, optFns ...func(*Options)) (*UpdateLibraryItemOutput, error) {
	if params == nil {
		params = &UpdateLibraryItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLibraryItem", params, optFns, c.addOperationUpdateLibraryItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLibraryItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLibraryItemInput struct {

	// The unique identifier of the Amazon Q Business application environment instance.
	//
	// This member is required.
	InstanceId *string

	// The unique identifier of the library item to update.
	//
	// This member is required.
	LibraryItemId *string

	// The new categories to associate with the library item.
	Categories []string

	// The new status to set for the library item, such as "Published" or "Hidden".
	Status types.LibraryItemStatus

	noSmithyDocumentSerde
}

type UpdateLibraryItemOutput struct {

	// The unique identifier of the Q App associated with the library item.
	//
	// This member is required.
	AppId *string

	// The version of the Q App associated with the library item.
	//
	// This member is required.
	AppVersion *int32

	// The categories associated with the updated library item.
	//
	// This member is required.
	Categories []types.Category

	// The date and time the library item was originally created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The user who originally created the library item.
	//
	// This member is required.
	CreatedBy *string

	// The unique identifier of the updated library item.
	//
	// This member is required.
	LibraryItemId *string

	// The number of ratings the library item has received.
	//
	// This member is required.
	RatingCount *int32

	// The new status of the updated library item.
	//
	// This member is required.
	Status *string

	// Whether the current user has rated the library item.
	IsRatedByUser *bool

	// The date and time the library item was last updated.
	UpdatedAt *time.Time

	// The user who last updated the library item.
	UpdatedBy *string

	// The number of users who have the associated Q App.
	UserCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLibraryItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLibraryItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLibraryItem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLibraryItem"); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateLibraryItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLibraryItem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLibraryItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLibraryItem",
	}
}
