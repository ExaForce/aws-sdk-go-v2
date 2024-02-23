// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the Firewall Manager policy administrator as a tenant administrator of a
// third-party firewall service. A tenant is an instance of the third-party
// firewall service that's associated with your Amazon Web Services customer
// account.
func (c *Client) AssociateThirdPartyFirewall(ctx context.Context, params *AssociateThirdPartyFirewallInput, optFns ...func(*Options)) (*AssociateThirdPartyFirewallOutput, error) {
	if params == nil {
		params = &AssociateThirdPartyFirewallInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateThirdPartyFirewall", params, optFns, c.addOperationAssociateThirdPartyFirewallMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateThirdPartyFirewallOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateThirdPartyFirewallInput struct {

	// The name of the third-party firewall vendor.
	//
	// This member is required.
	ThirdPartyFirewall types.ThirdPartyFirewall

	noSmithyDocumentSerde
}

type AssociateThirdPartyFirewallOutput struct {

	// The current status for setting a Firewall Manager policy administrator's
	// account as an administrator of the third-party firewall tenant.
	//   - ONBOARDING - The Firewall Manager policy administrator is being designated
	//   as a tenant administrator.
	//   - ONBOARD_COMPLETE - The Firewall Manager policy administrator is designated
	//   as a tenant administrator.
	//   - OFFBOARDING - The Firewall Manager policy administrator is being removed as
	//   a tenant administrator.
	//   - OFFBOARD_COMPLETE - The Firewall Manager policy administrator has been
	//   removed as a tenant administrator.
	//   - NOT_EXIST - The Firewall Manager policy administrator doesn't exist as a
	//   tenant administrator.
	ThirdPartyFirewallStatus types.ThirdPartyFirewallAssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateThirdPartyFirewallMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateThirdPartyFirewall{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateThirdPartyFirewall{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateThirdPartyFirewall"); err != nil {
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
	if err = addOpAssociateThirdPartyFirewallValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateThirdPartyFirewall(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateThirdPartyFirewall(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateThirdPartyFirewall",
	}
}
