// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates cluster broker volume size (or) sets cluster storage mode to TIERED.
func (c *Client) UpdateStorage(ctx context.Context, params *UpdateStorageInput, optFns ...func(*Options)) (*UpdateStorageOutput, error) {
	if params == nil {
		params = &UpdateStorageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateStorage", params, optFns, c.addOperationUpdateStorageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateStorageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request object for UpdateStorage api. Its used to update the storage attributes
// for the cluster.
type UpdateStorageInput struct {

	// The Amazon Resource Name (ARN) of the cluster to be updated.
	//
	// This member is required.
	ClusterArn *string

	// The version of cluster to update from. A successful operation will then
	// generate a new version.
	//
	// This member is required.
	CurrentVersion *string

	// EBS volume provisioned throughput information.
	ProvisionedThroughput *types.ProvisionedThroughput

	// Controls storage mode for supported storage tiers.
	StorageMode types.StorageMode

	// size of the EBS volume to update.
	VolumeSizeGB *int32

	noSmithyDocumentSerde
}

type UpdateStorageOutput struct {

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// The Amazon Resource Name (ARN) of the cluster operation.
	ClusterOperationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateStorageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateStorage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateStorage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateStorage"); err != nil {
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
	if err = addOpUpdateStorageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateStorage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateStorage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateStorage",
	}
}
