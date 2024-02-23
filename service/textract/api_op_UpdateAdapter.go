// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Update the configuration for an adapter. FeatureTypes configurations cannot be
// updated. At least one new parameter must be specified as an argument.
func (c *Client) UpdateAdapter(ctx context.Context, params *UpdateAdapterInput, optFns ...func(*Options)) (*UpdateAdapterOutput, error) {
	if params == nil {
		params = &UpdateAdapterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAdapter", params, optFns, c.addOperationUpdateAdapterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAdapterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAdapterInput struct {

	// A string containing a unique ID for the adapter that will be updated.
	//
	// This member is required.
	AdapterId *string

	// The new name to be applied to the adapter.
	AdapterName *string

	// The new auto-update status to be applied to the adapter.
	AutoUpdate types.AutoUpdate

	// The new description to be applied to the adapter.
	Description *string

	noSmithyDocumentSerde
}

type UpdateAdapterOutput struct {

	// A string containing a unique ID for the adapter that has been updated.
	AdapterId *string

	// A string containing the name of the adapter that has been updated.
	AdapterName *string

	// The auto-update status of the adapter that has been updated.
	AutoUpdate types.AutoUpdate

	// An object specifying the creation time of the the adapter that has been updated.
	CreationTime *time.Time

	// A string containing the description of the adapter that has been updated.
	Description *string

	// List of the targeted feature types for the updated adapter.
	FeatureTypes []types.FeatureType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAdapterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAdapter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAdapter{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAdapter"); err != nil {
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
	if err = addOpUpdateAdapterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAdapter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAdapter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAdapter",
	}
}
