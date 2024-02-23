// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a resource share. You can provide a list of the Amazon Resource Names
// (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
// for the resources that you want to share, a list of principals you want to share
// the resources with, and the permissions to grant those principals. Sharing a
// resource makes it available for use by principals outside of the Amazon Web
// Services account that created the resource. Sharing doesn't change any
// permissions or quotas that apply to the resource in the account that created it.
func (c *Client) CreateResourceShare(ctx context.Context, params *CreateResourceShareInput, optFns ...func(*Options)) (*CreateResourceShareOutput, error) {
	if params == nil {
		params = &CreateResourceShareInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResourceShare", params, optFns, c.addOperationCreateResourceShareMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResourceShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResourceShareInput struct {

	// Specifies the name of the resource share.
	//
	// This member is required.
	Name *string

	// Specifies whether principals outside your organization in Organizations can be
	// associated with a resource share. A value of true lets you share with
	// individual Amazon Web Services accounts that are not in your organization. A
	// value of false only has meaning if your account is a member of an Amazon Web
	// Services Organization. The default value is true .
	AllowExternalPrincipals *bool

	// Specifies a unique, case-sensitive identifier that you provide to ensure the
	// idempotency of the request. This lets you safely retry the request without
	// accidentally performing the same operation a second time. Passing the same value
	// to a later call to an operation requires that you also pass the same value for
	// all other parameters. We recommend that you use a UUID type of value. (https://wikipedia.org/wiki/Universally_unique_identifier)
	// . If you don't provide this value, then Amazon Web Services generates a random
	// one for you. If you retry the operation with the same ClientToken , but with
	// different parameters, the retry fails with an IdempotentParameterMismatch error.
	ClientToken *string

	// Specifies the Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the RAM permission to associate with the resource share. If you do not
	// specify an ARN for the permission, RAM automatically attaches the default
	// version of the permission for each resource type. You can associate only one
	// permission with each resource type included in the resource share.
	PermissionArns []string

	// Specifies a list of one or more principals to associate with the resource
	// share. You can include the following values:
	//   - An Amazon Web Services account ID, for example: 123456789012
	//   - An Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	//   of an organization in Organizations, for example:
	//   organizations::123456789012:organization/o-exampleorgid
	//   - An ARN of an organizational unit (OU) in Organizations, for example:
	//   organizations::123456789012:ou/o-exampleorgid/ou-examplerootid-exampleouid123
	//   - An ARN of an IAM role, for example: iam::123456789012:role/rolename
	//   - An ARN of an IAM user, for example: iam::123456789012user/username
	// Not all resource types can be shared with IAM roles and users. For more
	// information, see Sharing with IAM roles and users (https://docs.aws.amazon.com/ram/latest/userguide/permissions.html#permissions-rbp-supported-resource-types)
	// in the Resource Access Manager User Guide.
	Principals []string

	// Specifies a list of one or more ARNs of the resources to associate with the
	// resource share.
	ResourceArns []string

	// Specifies from which source accounts the service principal has access to the
	// resources in this resource share.
	Sources []string

	// Specifies one or more tags to attach to the resource share itself. It doesn't
	// attach the tags to the resources associated with the resource share.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateResourceShareOutput struct {

	// The idempotency identifier associated with this request. If you want to repeat
	// the same operation in an idempotent manner then you must include this value in
	// the clientToken request parameter of that later call. All other parameters must
	// also have the same values that you used in the first call.
	ClientToken *string

	// An object with information about the new resource share.
	ResourceShare *types.ResourceShare

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResourceShareMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateResourceShare{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateResourceShare{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateResourceShare"); err != nil {
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
	if err = addOpCreateResourceShareValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResourceShare(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateResourceShare(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateResourceShare",
	}
}
