// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a live source's configuration.
func (c *Client) UpdateLiveSource(ctx context.Context, params *UpdateLiveSourceInput, optFns ...func(*Options)) (*UpdateLiveSourceOutput, error) {
	if params == nil {
		params = &UpdateLiveSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLiveSource", params, optFns, c.addOperationUpdateLiveSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLiveSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLiveSourceInput struct {

	// A list of HTTP package configurations for the live source on this account.
	//
	// This member is required.
	HttpPackageConfigurations []types.HttpPackageConfiguration

	// The name of the live source.
	//
	// This member is required.
	LiveSourceName *string

	// The name of the source location associated with this Live Source.
	//
	// This member is required.
	SourceLocationName *string

	noSmithyDocumentSerde
}

type UpdateLiveSourceOutput struct {

	// The Amazon Resource Name (ARN) associated with this live source.
	Arn *string

	// The timestamp that indicates when the live source was created.
	CreationTime *time.Time

	// A list of HTTP package configurations for the live source on this account.
	HttpPackageConfigurations []types.HttpPackageConfiguration

	// The timestamp that indicates when the live source was last modified.
	LastModifiedTime *time.Time

	// The name of the live source.
	LiveSourceName *string

	// The name of the source location associated with the live source.
	SourceLocationName *string

	// The tags to assign to the live source. Tags are key-value pairs that you can
	// associate with Amazon resources to help with organization, access control, and
	// cost tracking. For more information, see Tagging AWS Elemental MediaTailor
	// Resources (https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html) .
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLiveSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLiveSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLiveSource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateLiveSource"); err != nil {
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
	if err = addOpUpdateLiveSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLiveSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLiveSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateLiveSource",
	}
}
