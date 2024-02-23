// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides details regarding the entity used with the connector, with a
// description of the data model for each field in that entity.
func (c *Client) DescribeConnectorEntity(ctx context.Context, params *DescribeConnectorEntityInput, optFns ...func(*Options)) (*DescribeConnectorEntityOutput, error) {
	if params == nil {
		params = &DescribeConnectorEntityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConnectorEntity", params, optFns, c.addOperationDescribeConnectorEntityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConnectorEntityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConnectorEntityInput struct {

	// The entity name for that connector.
	//
	// This member is required.
	ConnectorEntityName *string

	// The version of the API that's used by the connector.
	ApiVersion *string

	// The name of the connector profile. The name is unique for each ConnectorProfile
	// in the Amazon Web Services account.
	ConnectorProfileName *string

	// The type of connector application, such as Salesforce, Amplitude, and so on.
	ConnectorType types.ConnectorType

	noSmithyDocumentSerde
}

type DescribeConnectorEntityOutput struct {

	// Describes the fields for that connector entity. For example, for an account
	// entity, the fields would be account name, account ID, and so on.
	//
	// This member is required.
	ConnectorEntityFields []types.ConnectorEntityField

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConnectorEntityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeConnectorEntity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeConnectorEntity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeConnectorEntity"); err != nil {
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
	if err = addOpDescribeConnectorEntityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConnectorEntity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConnectorEntity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeConnectorEntity",
	}
}
