// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the identity attribute order for a specific TypedLinkFacet . For more
// information, see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink)
// .
func (c *Client) GetTypedLinkFacetInformation(ctx context.Context, params *GetTypedLinkFacetInformationInput, optFns ...func(*Options)) (*GetTypedLinkFacetInformationOutput, error) {
	if params == nil {
		params = &GetTypedLinkFacetInformationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTypedLinkFacetInformation", params, optFns, c.addOperationGetTypedLinkFacetInformationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTypedLinkFacetInformationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTypedLinkFacetInformationInput struct {

	// The unique name of the typed link facet.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) that is associated with the schema. For more
	// information, see arns .
	//
	// This member is required.
	SchemaArn *string

	noSmithyDocumentSerde
}

type GetTypedLinkFacetInformationOutput struct {

	// The order of identity attributes for the facet, from most significant to least
	// significant. The ability to filter typed links considers the order that the
	// attributes are defined on the typed link facet. When providing ranges to typed
	// link selection, any inexact ranges must be specified at the end. Any attributes
	// that do not have a range specified are presumed to match the entire range.
	// Filters are interpreted in the order of the attributes on the typed link facet,
	// not the order in which they are supplied to any API calls. For more information
	// about identity attributes, see Typed Links (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink)
	// .
	IdentityAttributeOrder []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTypedLinkFacetInformationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTypedLinkFacetInformation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTypedLinkFacetInformation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetTypedLinkFacetInformation"); err != nil {
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
	if err = addOpGetTypedLinkFacetInformationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTypedLinkFacetInformation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTypedLinkFacetInformation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetTypedLinkFacetInformation",
	}
}
