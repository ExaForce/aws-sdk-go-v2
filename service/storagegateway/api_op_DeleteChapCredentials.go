// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes Challenge-Handshake Authentication Protocol (CHAP) credentials for a
// specified iSCSI target and initiator pair. This operation is supported in volume
// and tape gateway types.
func (c *Client) DeleteChapCredentials(ctx context.Context, params *DeleteChapCredentialsInput, optFns ...func(*Options)) (*DeleteChapCredentialsOutput, error) {
	if params == nil {
		params = &DeleteChapCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteChapCredentials", params, optFns, c.addOperationDeleteChapCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteChapCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing one or more of the following fields:
//   - DeleteChapCredentialsInput$InitiatorName
//   - DeleteChapCredentialsInput$TargetARN
type DeleteChapCredentialsInput struct {

	// The iSCSI initiator that connects to the target.
	//
	// This member is required.
	InitiatorName *string

	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the
	// DescribeStorediSCSIVolumes operation to return to retrieve the TargetARN for
	// specified VolumeARN.
	//
	// This member is required.
	TargetARN *string

	noSmithyDocumentSerde
}

// A JSON object containing the following fields:
type DeleteChapCredentialsOutput struct {

	// The iSCSI initiator that connects to the target.
	InitiatorName *string

	// The Amazon Resource Name (ARN) of the target.
	TargetARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteChapCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteChapCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteChapCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteChapCredentials"); err != nil {
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
	if err = addOpDeleteChapCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteChapCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteChapCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteChapCredentials",
	}
}
