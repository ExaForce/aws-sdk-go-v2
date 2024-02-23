// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a group of permissions for various actions that a user can perform in
// FinSpace.
//
// Deprecated: This method will be discontinued.
func (c *Client) CreatePermissionGroup(ctx context.Context, params *CreatePermissionGroupInput, optFns ...func(*Options)) (*CreatePermissionGroupOutput, error) {
	if params == nil {
		params = &CreatePermissionGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePermissionGroup", params, optFns, c.addOperationCreatePermissionGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePermissionGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePermissionGroupInput struct {

	// The option to indicate FinSpace application permissions that are granted to a
	// specific group. When assigning application permissions, be aware that the
	// permission ManageUsersAndGroups allows users to grant themselves or others
	// access to any functionality in their FinSpace environment's application. It
	// should only be granted to trusted users.
	//   - CreateDataset – Group members can create new datasets.
	//   - ManageClusters – Group members can manage Apache Spark clusters from
	//   FinSpace notebooks.
	//   - ManageUsersAndGroups – Group members can manage users and permission groups.
	//   This is a privileged permission that allows users to grant themselves or others
	//   access to any functionality in the application. It should only be granted to
	//   trusted users.
	//   - ManageAttributeSets – Group members can manage attribute sets.
	//   - ViewAuditData – Group members can view audit data.
	//   - AccessNotebooks – Group members will have access to FinSpace notebooks.
	//   - GetTemporaryCredentials – Group members can get temporary API credentials.
	//
	// This member is required.
	ApplicationPermissions []types.ApplicationPermission

	// The name of the permission group.
	//
	// This member is required.
	Name *string

	// A token that ensures idempotency. This token expires in 10 minutes.
	ClientToken *string

	// A brief description for the permission group.
	Description *string

	noSmithyDocumentSerde
}

type CreatePermissionGroupOutput struct {

	// The unique identifier for the permission group.
	PermissionGroupId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePermissionGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreatePermissionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreatePermissionGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePermissionGroup"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addIdempotencyToken_opCreatePermissionGroupMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreatePermissionGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePermissionGroup(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreatePermissionGroup struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreatePermissionGroup) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreatePermissionGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreatePermissionGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreatePermissionGroupInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreatePermissionGroupMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreatePermissionGroup{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreatePermissionGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePermissionGroup",
	}
}
