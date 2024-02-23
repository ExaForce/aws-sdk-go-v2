// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a key group. When you update a key group, all the fields are updated
// with the values provided in the request. You cannot update some fields
// independent of others. To update a key group:
//   - Get the current key group with GetKeyGroup or GetKeyGroupConfig .
//   - Locally modify the fields in the key group that you want to update. For
//     example, add or remove public key IDs.
//   - Call UpdateKeyGroup with the entire key group object, including the fields
//     that you modified and those that you didn't.
func (c *Client) UpdateKeyGroup(ctx context.Context, params *UpdateKeyGroupInput, optFns ...func(*Options)) (*UpdateKeyGroupOutput, error) {
	if params == nil {
		params = &UpdateKeyGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateKeyGroup", params, optFns, c.addOperationUpdateKeyGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateKeyGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateKeyGroupInput struct {

	// The identifier of the key group that you are updating.
	//
	// This member is required.
	Id *string

	// The key group configuration.
	//
	// This member is required.
	KeyGroupConfig *types.KeyGroupConfig

	// The version of the key group that you are updating. The version is the key
	// group's ETag value.
	IfMatch *string

	noSmithyDocumentSerde
}

type UpdateKeyGroupOutput struct {

	// The identifier for this version of the key group.
	ETag *string

	// The key group that was just updated.
	KeyGroup *types.KeyGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateKeyGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateKeyGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateKeyGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateKeyGroup"); err != nil {
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
	if err = addOpUpdateKeyGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateKeyGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateKeyGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateKeyGroup",
	}
}
