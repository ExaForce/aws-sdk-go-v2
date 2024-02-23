// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves table metadata from the Data Catalog that contains unfiltered
// metadata. For IAM authorization, the public IAM action associated with this API
// is glue:GetTable .
func (c *Client) GetUnfilteredTableMetadata(ctx context.Context, params *GetUnfilteredTableMetadataInput, optFns ...func(*Options)) (*GetUnfilteredTableMetadataOutput, error) {
	if params == nil {
		params = &GetUnfilteredTableMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUnfilteredTableMetadata", params, optFns, c.addOperationGetUnfilteredTableMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUnfilteredTableMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUnfilteredTableMetadataInput struct {

	// The catalog ID where the table resides.
	//
	// This member is required.
	CatalogId *string

	// (Required) Specifies the name of a database that contains the table.
	//
	// This member is required.
	DatabaseName *string

	// (Required) Specifies the name of a table for which you are requesting metadata.
	//
	// This member is required.
	Name *string

	// (Required) A list of supported permission types.
	//
	// This member is required.
	SupportedPermissionTypes []types.PermissionType

	// A structure containing Lake Formation audit context information.
	AuditContext *types.AuditContext

	// The Lake Formation data permissions of the caller on the table. Used to
	// authorize the call when no view context is found.
	Permissions []types.Permission

	// A structure used as a protocol between query engines and Lake Formation or
	// Glue. Contains both a Lake Formation generated authorization identifier and
	// information from the request's authorization context.
	QuerySessionContext *types.QuerySessionContext

	// Specified only if the base tables belong to a different Amazon Web Services
	// Region.
	Region *string

	// A structure specifying the dialect and dialect version used by the query engine.
	SupportedDialect *types.SupportedDialect

	noSmithyDocumentSerde
}

type GetUnfilteredTableMetadataOutput struct {

	// A list of column names that the user has been granted access to.
	AuthorizedColumns []string

	// A list of column row filters.
	CellFilters []types.ColumnRowFilter

	// A Boolean value that indicates whether the partition location is registered
	// with Lake Formation.
	IsRegisteredWithLakeFormation bool

	// The Lake Formation data permissions of the caller on the table. Used to
	// authorize the call when no view context is found.
	Permissions []types.Permission

	// A cryptographically generated query identifier generated by Glue or Lake
	// Formation.
	QueryAuthorizationId *string

	// The resource ARN of the parent resource extracted from the request.
	ResourceArn *string

	// A Table object containing the table metadata.
	Table *types.Table

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUnfilteredTableMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetUnfilteredTableMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUnfilteredTableMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetUnfilteredTableMetadata"); err != nil {
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
	if err = addOpGetUnfilteredTableMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUnfilteredTableMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUnfilteredTableMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetUnfilteredTableMetadata",
	}
}
