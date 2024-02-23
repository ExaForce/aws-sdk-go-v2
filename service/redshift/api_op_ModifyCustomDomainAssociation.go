// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Contains information for changing a custom domain association.
func (c *Client) ModifyCustomDomainAssociation(ctx context.Context, params *ModifyCustomDomainAssociationInput, optFns ...func(*Options)) (*ModifyCustomDomainAssociationOutput, error) {
	if params == nil {
		params = &ModifyCustomDomainAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyCustomDomainAssociation", params, optFns, c.addOperationModifyCustomDomainAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyCustomDomainAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyCustomDomainAssociationInput struct {

	// The identifier of the cluster to change a custom domain association for.
	//
	// This member is required.
	ClusterIdentifier *string

	// The certificate Amazon Resource Name (ARN) for the changed custom domain
	// association.
	//
	// This member is required.
	CustomDomainCertificateArn *string

	// The custom domain name for a changed custom domain association.
	//
	// This member is required.
	CustomDomainName *string

	noSmithyDocumentSerde
}

type ModifyCustomDomainAssociationOutput struct {

	// The identifier of the cluster associated with the result for the changed custom
	// domain association.
	ClusterIdentifier *string

	// The certificate expiration time associated with the result for the changed
	// custom domain association.
	CustomDomainCertExpiryTime *string

	// The certificate Amazon Resource Name (ARN) associated with the result for the
	// changed custom domain association.
	CustomDomainCertificateArn *string

	// The custom domain name associated with the result for the changed custom domain
	// association.
	CustomDomainName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyCustomDomainAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyCustomDomainAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyCustomDomainAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyCustomDomainAssociation"); err != nil {
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
	if err = addOpModifyCustomDomainAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCustomDomainAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyCustomDomainAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyCustomDomainAssociation",
	}
}
