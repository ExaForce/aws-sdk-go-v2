// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enable portfolio sharing feature through Organizations. This API will allow
// Service Catalog to receive updates on your organization in order to sync your
// shares with the current structure. This API can only be called by the management
// account in the organization. When you call this API, Service Catalog calls
// organizations:EnableAWSServiceAccess on your behalf so that your shares stay in
// sync with any changes in your Organizations structure. Note that a delegated
// administrator is not authorized to invoke EnableAWSOrganizationsAccess . If you
// have previously disabled Organizations access for Service Catalog, and then
// enable access again, the portfolio access permissions might not sync with the
// latest changes to the organization structure. Specifically, accounts that you
// removed from the organization after disabling Service Catalog access, and before
// you enabled access again, can retain access to the previously shared portfolio.
// As a result, an account that has been removed from the organization might still
// be able to create or manage Amazon Web Services resources when it is no longer
// authorized to do so. Amazon Web Services is working to resolve this issue.
func (c *Client) EnableAWSOrganizationsAccess(ctx context.Context, params *EnableAWSOrganizationsAccessInput, optFns ...func(*Options)) (*EnableAWSOrganizationsAccessOutput, error) {
	if params == nil {
		params = &EnableAWSOrganizationsAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableAWSOrganizationsAccess", params, optFns, c.addOperationEnableAWSOrganizationsAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableAWSOrganizationsAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableAWSOrganizationsAccessInput struct {
	noSmithyDocumentSerde
}

type EnableAWSOrganizationsAccessOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableAWSOrganizationsAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableAWSOrganizationsAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableAWSOrganizationsAccess{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EnableAWSOrganizationsAccess"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableAWSOrganizationsAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableAWSOrganizationsAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EnableAWSOrganizationsAccess",
	}
}
