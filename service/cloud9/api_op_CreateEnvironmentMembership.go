// Code generated by smithy-go-codegen DO NOT EDIT.

package cloud9

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloud9/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds an environment member to an Cloud9 development environment.
func (c *Client) CreateEnvironmentMembership(ctx context.Context, params *CreateEnvironmentMembershipInput, optFns ...func(*Options)) (*CreateEnvironmentMembershipOutput, error) {
	if params == nil {
		params = &CreateEnvironmentMembershipInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironmentMembership", params, optFns, c.addOperationCreateEnvironmentMembershipMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentMembershipInput struct {

	// The ID of the environment that contains the environment member you want to add.
	//
	// This member is required.
	EnvironmentId *string

	// The type of environment member permissions you want to associate with this
	// environment member. Available values include:
	//   - read-only : Has read-only access to the environment.
	//   - read-write : Has read-write access to the environment.
	//
	// This member is required.
	Permissions types.MemberPermissions

	// The Amazon Resource Name (ARN) of the environment member you want to add.
	//
	// This member is required.
	UserArn *string

	noSmithyDocumentSerde
}

type CreateEnvironmentMembershipOutput struct {

	// Information about the environment member that was added.
	//
	// This member is required.
	Membership *types.EnvironmentMember

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentMembershipMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEnvironmentMembership{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEnvironmentMembership{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEnvironmentMembership"); err != nil {
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
	if err = addOpCreateEnvironmentMembershipValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironmentMembership(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEnvironmentMembership(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEnvironmentMembership",
	}
}
