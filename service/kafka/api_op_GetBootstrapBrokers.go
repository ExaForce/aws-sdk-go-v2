// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A list of brokers that a client application can use to bootstrap.
func (c *Client) GetBootstrapBrokers(ctx context.Context, params *GetBootstrapBrokersInput, optFns ...func(*Options)) (*GetBootstrapBrokersOutput, error) {
	if params == nil {
		params = &GetBootstrapBrokersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBootstrapBrokers", params, optFns, c.addOperationGetBootstrapBrokersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBootstrapBrokersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBootstrapBrokersInput struct {

	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	//
	// This member is required.
	ClusterArn *string

	noSmithyDocumentSerde
}

type GetBootstrapBrokersOutput struct {

	// A string containing one or more hostname:port pairs.
	BootstrapBrokerString *string

	// A string that contains one or more DNS names (or IP addresses) and SASL IAM
	// port pairs.
	BootstrapBrokerStringPublicSaslIam *string

	// A string containing one or more DNS names (or IP) and Sasl Scram port pairs.
	BootstrapBrokerStringPublicSaslScram *string

	// A string containing one or more DNS names (or IP) and TLS port pairs.
	BootstrapBrokerStringPublicTls *string

	// A string that contains one or more DNS names (or IP addresses) and SASL IAM
	// port pairs.
	BootstrapBrokerStringSaslIam *string

	// A string containing one or more DNS names (or IP) and Sasl Scram port pairs.
	BootstrapBrokerStringSaslScram *string

	// A string containing one or more DNS names (or IP) and TLS port pairs.
	BootstrapBrokerStringTls *string

	// A string containing one or more DNS names (or IP) and SASL/IAM port pairs for
	// VPC connectivity.
	BootstrapBrokerStringVpcConnectivitySaslIam *string

	// A string containing one or more DNS names (or IP) and SASL/SCRAM port pairs for
	// VPC connectivity.
	BootstrapBrokerStringVpcConnectivitySaslScram *string

	// A string containing one or more DNS names (or IP) and TLS port pairs for VPC
	// connectivity.
	BootstrapBrokerStringVpcConnectivityTls *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBootstrapBrokersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBootstrapBrokers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBootstrapBrokers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetBootstrapBrokers"); err != nil {
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
	if err = addOpGetBootstrapBrokersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBootstrapBrokers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBootstrapBrokers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetBootstrapBrokers",
	}
}
