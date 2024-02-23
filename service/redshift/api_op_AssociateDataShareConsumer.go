// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// From a datashare consumer account, associates a datashare with the account
// (AssociateEntireAccount) or the specified namespace (ConsumerArn). If you make
// this association, the consumer can consume the datashare.
func (c *Client) AssociateDataShareConsumer(ctx context.Context, params *AssociateDataShareConsumerInput, optFns ...func(*Options)) (*AssociateDataShareConsumerOutput, error) {
	if params == nil {
		params = &AssociateDataShareConsumerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateDataShareConsumer", params, optFns, c.addOperationAssociateDataShareConsumerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateDataShareConsumerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateDataShareConsumerInput struct {

	// The Amazon Resource Name (ARN) of the datashare that the consumer is to use
	// with the account or the namespace.
	//
	// This member is required.
	DataShareArn *string

	// If set to true, allows write operations for a datashare.
	AllowWrites *bool

	// A value that specifies whether the datashare is associated with the entire
	// account.
	AssociateEntireAccount *bool

	// The Amazon Resource Name (ARN) of the consumer that is associated with the
	// datashare.
	ConsumerArn *string

	// From a datashare consumer account, associates a datashare with all existing and
	// future namespaces in the specified Amazon Web Services Region.
	ConsumerRegion *string

	noSmithyDocumentSerde
}

type AssociateDataShareConsumerOutput struct {

	// A value that specifies whether the datashare can be shared to a publicly
	// accessible cluster.
	AllowPubliclyAccessibleConsumers *bool

	// An Amazon Resource Name (ARN) that references the datashare that is owned by a
	// specific namespace of the producer cluster. A datashare ARN is in the
	// arn:aws:redshift:{region}:{account-id}:{datashare}:{namespace-guid}/{datashare-name}
	// format.
	DataShareArn *string

	// A value that specifies when the datashare has an association between producer
	// and data consumers.
	DataShareAssociations []types.DataShareAssociation

	// The identifier of a datashare to show its managing entity.
	ManagedBy *string

	// The Amazon Resource Name (ARN) of the producer.
	ProducerArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateDataShareConsumerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAssociateDataShareConsumer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAssociateDataShareConsumer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateDataShareConsumer"); err != nil {
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
	if err = addOpAssociateDataShareConsumerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateDataShareConsumer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateDataShareConsumer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateDataShareConsumer",
	}
}
