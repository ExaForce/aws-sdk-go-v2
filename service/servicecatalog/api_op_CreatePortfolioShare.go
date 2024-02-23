// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Shares the specified portfolio with the specified account or organization node.
// Shares to an organization node can only be created by the management account of
// an organization or by a delegated administrator. You can share portfolios to an
// organization, an organizational unit, or a specific account. Note that if a
// delegated admin is de-registered, they can no longer create portfolio shares.
// AWSOrganizationsAccess must be enabled in order to create a portfolio share to
// an organization node. You can't share a shared resource, including portfolios
// that contain a shared product. If the portfolio share with the specified account
// or organization node already exists, this action will have no effect and will
// not return an error. To update an existing share, you must use the
// UpdatePortfolioShare API instead. When you associate a principal with portfolio,
// a potential privilege escalation path may occur when that portfolio is then
// shared with other accounts. For a user in a recipient account who is not an
// Service Catalog Admin, but still has the ability to create Principals
// (Users/Groups/Roles), that user could create a role that matches a principal
// name association for the portfolio. Although this user may not know which
// principal names are associated through Service Catalog, they may be able to
// guess the user. If this potential escalation path is a concern, then Service
// Catalog recommends using PrincipalType as IAM . With this configuration, the
// PrincipalARN must already exist in the recipient account before it can be
// associated.
func (c *Client) CreatePortfolioShare(ctx context.Context, params *CreatePortfolioShareInput, optFns ...func(*Options)) (*CreatePortfolioShareOutput, error) {
	if params == nil {
		params = &CreatePortfolioShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePortfolioShare", params, optFns, c.addOperationCreatePortfolioShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePortfolioShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePortfolioShareInput struct {

	// The portfolio identifier.
	//
	// This member is required.
	PortfolioId *string

	// The language code.
	//   - jp - Japanese
	//   - zh - Chinese
	AcceptLanguage *string

	// The Amazon Web Services account ID. For example, 123456789012 .
	AccountId *string

	// The organization node to whom you are going to share. When you pass
	// OrganizationNode , it creates PortfolioShare for all of the Amazon Web Services
	// accounts that are associated to the OrganizationNode . The output returns a
	// PortfolioShareToken , which enables the administrator to monitor the status of
	// the PortfolioShare creation process.
	OrganizationNode *types.OrganizationNode

	// This parameter is only supported for portfolios with an OrganizationalNode Type
	// of ORGANIZATION or ORGANIZATIONAL_UNIT . Enables or disables Principal sharing
	// when creating the portfolio share. If you do not provide this flag, principal
	// sharing is disabled. When you enable Principal Name Sharing for a portfolio
	// share, the share recipient account end users with a principal that matches any
	// of the associated IAM patterns can provision products from the portfolio. Once
	// shared, the share recipient can view associations of PrincipalType : IAM_PATTERN
	// on their portfolio. You can create the principals in the recipient account
	// before or after creating the share.
	SharePrincipals bool

	// Enables or disables TagOptions  sharing when creating the portfolio share. If
	// this flag is not provided, TagOptions sharing is disabled.
	ShareTagOptions bool

	noSmithyDocumentSerde
}

type CreatePortfolioShareOutput struct {

	// The portfolio shares a unique identifier that only returns if the portfolio is
	// shared to an organization node.
	PortfolioShareToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePortfolioShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePortfolioShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePortfolioShare{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePortfolioShare"); err != nil {
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
	if err = addOpCreatePortfolioShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePortfolioShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePortfolioShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePortfolioShare",
	}
}
