// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a flow for the specified Amazon Connect instance. You can also create
// and update flows using the Amazon Connect Flow language (https://docs.aws.amazon.com/connect/latest/APIReference/flow-language.html)
// .
func (c *Client) CreateContactFlow(ctx context.Context, params *CreateContactFlowInput, optFns ...func(*Options)) (*CreateContactFlowOutput, error) {
	if params == nil {
		params = &CreateContactFlowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateContactFlow", params, optFns, c.addOperationCreateContactFlowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateContactFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContactFlowInput struct {

	// The JSON string that represents the content of the flow. For an example, see
	// Example flow in Amazon Connect Flow language (https://docs.aws.amazon.com/connect/latest/APIReference/flow-language-example.html)
	// . Length Constraints: Minimum length of 1. Maximum length of 256000.
	//
	// This member is required.
	Content *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The name of the flow.
	//
	// This member is required.
	Name *string

	// The type of the flow. For descriptions of the available types, see Choose a
	// flow type (https://docs.aws.amazon.com/connect/latest/adminguide/create-contact-flow.html#contact-flow-types)
	// in the Amazon Connect Administrator Guide.
	//
	// This member is required.
	Type types.ContactFlowType

	// The description of the flow.
	Description *string

	// The tags used to organize, track, or control access for this resource. For
	// example, { "Tags": {"key1":"value1", "key2":"value2"} }.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateContactFlowOutput struct {

	// The Amazon Resource Name (ARN) of the flow.
	ContactFlowArn *string

	// The identifier of the flow.
	ContactFlowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateContactFlowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateContactFlow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateContactFlow{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateContactFlow"); err != nil {
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
	if err = addOpCreateContactFlowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContactFlow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateContactFlow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateContactFlow",
	}
}
