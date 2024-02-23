// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a storage virtual machine (SVM) for an Amazon FSx for ONTAP file system.
func (c *Client) CreateStorageVirtualMachine(ctx context.Context, params *CreateStorageVirtualMachineInput, optFns ...func(*Options)) (*CreateStorageVirtualMachineOutput, error) {
	if params == nil {
		params = &CreateStorageVirtualMachineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStorageVirtualMachine", params, optFns, c.addOperationCreateStorageVirtualMachineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStorageVirtualMachineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStorageVirtualMachineInput struct {

	// The globally unique ID of the file system, assigned by Amazon FSx.
	//
	// This member is required.
	FileSystemId *string

	// The name of the SVM.
	//
	// This member is required.
	Name *string

	// Describes the self-managed Microsoft Active Directory to which you want to join
	// the SVM. Joining an Active Directory provides user authentication and access
	// control for SMB clients, including Microsoft Windows and macOS client accessing
	// the file system.
	ActiveDirectoryConfiguration *types.CreateSvmActiveDirectoryConfiguration

	// (Optional) An idempotency token for resource creation, in a string of up to 63
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// The security style of the root volume of the SVM. Specify one of the following
	// values:
	//   - UNIX if the file system is managed by a UNIX administrator, the majority of
	//   users are NFS clients, and an application accessing the data uses a UNIX user as
	//   the service account.
	//   - NTFS if the file system is managed by a Windows administrator, the majority
	//   of users are SMB clients, and an application accessing the data uses a Windows
	//   user as the service account.
	//   - MIXED if the file system is managed by both UNIX and Windows administrators
	//   and users consist of both NFS and SMB clients.
	RootVolumeSecurityStyle types.StorageVirtualMachineRootVolumeSecurityStyle

	// The password to use when managing the SVM using the NetApp ONTAP CLI or REST
	// API. If you do not specify a password, you can still use the file system's
	// fsxadmin user to manage the SVM.
	SvmAdminPassword *string

	// A list of Tag values, with a maximum of 50 elements.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateStorageVirtualMachineOutput struct {

	// Returned after a successful CreateStorageVirtualMachine operation; describes
	// the SVM just created.
	StorageVirtualMachine *types.StorageVirtualMachine

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStorageVirtualMachineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStorageVirtualMachine{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStorageVirtualMachine{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateStorageVirtualMachine"); err != nil {
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
	if err = addIdempotencyToken_opCreateStorageVirtualMachineMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateStorageVirtualMachineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStorageVirtualMachine(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateStorageVirtualMachine struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateStorageVirtualMachine) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateStorageVirtualMachine) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateStorageVirtualMachineInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateStorageVirtualMachineInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateStorageVirtualMachineMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateStorageVirtualMachine{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateStorageVirtualMachine(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateStorageVirtualMachine",
	}
}
