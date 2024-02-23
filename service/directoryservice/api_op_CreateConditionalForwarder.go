// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a conditional forwarder associated with your Amazon Web Services
// directory. Conditional forwarders are required in order to set up a trust
// relationship with another domain. The conditional forwarder points to the
// trusted domain.
func (c *Client) CreateConditionalForwarder(ctx context.Context, params *CreateConditionalForwarderInput, optFns ...func(*Options)) (*CreateConditionalForwarderOutput, error) {
	if params == nil {
		params = &CreateConditionalForwarderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConditionalForwarder", params, optFns, c.addOperationCreateConditionalForwarderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConditionalForwarderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Initiates the creation of a conditional forwarder for your Directory Service
// for Microsoft Active Directory. Conditional forwarders are required in order to
// set up a trust relationship with another domain.
type CreateConditionalForwarderInput struct {

	// The directory ID of the Amazon Web Services directory for which you are
	// creating the conditional forwarder.
	//
	// This member is required.
	DirectoryId *string

	// The IP addresses of the remote DNS server associated with RemoteDomainName.
	//
	// This member is required.
	DnsIpAddrs []string

	// The fully qualified domain name (FQDN) of the remote domain with which you will
	// set up a trust relationship.
	//
	// This member is required.
	RemoteDomainName *string

	noSmithyDocumentSerde
}

// The result of a CreateConditinalForwarder request.
type CreateConditionalForwarderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConditionalForwarderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateConditionalForwarder{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateConditionalForwarder{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateConditionalForwarder"); err != nil {
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
	if err = addOpCreateConditionalForwarderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConditionalForwarder(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConditionalForwarder(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateConditionalForwarder",
	}
}
