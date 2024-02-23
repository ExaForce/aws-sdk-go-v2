// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the current configuration items for resources that are present in your
// Config aggregator. The operation also returns a list of resources that are not
// processed in the current request. If there are no unprocessed resources, the
// operation returns an empty unprocessedResourceIdentifiers list.
//   - The API does not return results for deleted resources.
//   - The API does not return tags and relationships.
func (c *Client) BatchGetAggregateResourceConfig(ctx context.Context, params *BatchGetAggregateResourceConfigInput, optFns ...func(*Options)) (*BatchGetAggregateResourceConfigOutput, error) {
	if params == nil {
		params = &BatchGetAggregateResourceConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetAggregateResourceConfig", params, optFns, c.addOperationBatchGetAggregateResourceConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetAggregateResourceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetAggregateResourceConfigInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// A list of aggregate ResourceIdentifiers objects.
	//
	// This member is required.
	ResourceIdentifiers []types.AggregateResourceIdentifier

	noSmithyDocumentSerde
}

type BatchGetAggregateResourceConfigOutput struct {

	// A list that contains the current configuration of one or more resources.
	BaseConfigurationItems []types.BaseConfigurationItem

	// A list of resource identifiers that were not processed with current scope. The
	// list is empty if all the resources are processed.
	UnprocessedResourceIdentifiers []types.AggregateResourceIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetAggregateResourceConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetAggregateResourceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetAggregateResourceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetAggregateResourceConfig"); err != nil {
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
	if err = addOpBatchGetAggregateResourceConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetAggregateResourceConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetAggregateResourceConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetAggregateResourceConfig",
	}
}
