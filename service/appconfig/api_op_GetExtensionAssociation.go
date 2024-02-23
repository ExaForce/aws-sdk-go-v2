// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about an AppConfig extension association. For more
// information about extensions and associations, see Working with AppConfig
// extensions (https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html)
// in the AppConfig User Guide.
func (c *Client) GetExtensionAssociation(ctx context.Context, params *GetExtensionAssociationInput, optFns ...func(*Options)) (*GetExtensionAssociationOutput, error) {
	if params == nil {
		params = &GetExtensionAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetExtensionAssociation", params, optFns, c.addOperationGetExtensionAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetExtensionAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetExtensionAssociationInput struct {

	// The extension association ID to get.
	//
	// This member is required.
	ExtensionAssociationId *string

	noSmithyDocumentSerde
}

type GetExtensionAssociationOutput struct {

	// The system-generated Amazon Resource Name (ARN) for the extension.
	Arn *string

	// The ARN of the extension defined in the association.
	ExtensionArn *string

	// The version number for the extension defined in the association.
	ExtensionVersionNumber int32

	// The system-generated ID for the association.
	Id *string

	// The parameter names and values defined in the association.
	Parameters map[string]string

	// The ARNs of applications, configuration profiles, or environments defined in
	// the association.
	ResourceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetExtensionAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetExtensionAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetExtensionAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetExtensionAssociation"); err != nil {
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
	if err = addOpGetExtensionAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetExtensionAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetExtensionAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetExtensionAssociation",
	}
}
