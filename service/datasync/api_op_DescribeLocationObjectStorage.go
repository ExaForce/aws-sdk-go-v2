// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides details about how an DataSync transfer location for an object storage
// system is configured.
func (c *Client) DescribeLocationObjectStorage(ctx context.Context, params *DescribeLocationObjectStorageInput, optFns ...func(*Options)) (*DescribeLocationObjectStorageOutput, error) {
	if params == nil {
		params = &DescribeLocationObjectStorageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLocationObjectStorage", params, optFns, c.addOperationDescribeLocationObjectStorageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLocationObjectStorageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeLocationObjectStorageRequest
type DescribeLocationObjectStorageInput struct {

	// Specifies the Amazon Resource Name (ARN) of the object storage system location.
	//
	// This member is required.
	LocationArn *string

	noSmithyDocumentSerde
}

// DescribeLocationObjectStorageResponse
type DescribeLocationObjectStorageOutput struct {

	// The access key (for example, a user name) required to authenticate with the
	// object storage system.
	AccessKey *string

	// The ARNs of the DataSync agents that can connect with your object storage
	// system.
	AgentArns []string

	// The time that the location was created.
	CreationTime *time.Time

	// The ARN of the object storage system location.
	LocationArn *string

	// The URI of the object storage system location.
	LocationUri *string

	// The self-signed certificate that DataSync uses to securely authenticate with
	// your object storage system.
	ServerCertificate []byte

	// The port that your object storage server accepts inbound network traffic on
	// (for example, port 443).
	ServerPort *int32

	// The protocol that your object storage system uses to communicate.
	ServerProtocol types.ObjectStorageServerProtocol

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLocationObjectStorageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLocationObjectStorage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLocationObjectStorage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeLocationObjectStorage"); err != nil {
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
	if err = addOpDescribeLocationObjectStorageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLocationObjectStorage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLocationObjectStorage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeLocationObjectStorage",
	}
}
