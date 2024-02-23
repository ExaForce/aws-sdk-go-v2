// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new group in the specified user pool. Amazon Cognito evaluates
// Identity and Access Management (IAM) policies in requests for this API
// operation. For this operation, you must use IAM credentials to authorize
// requests, and you must grant yourself the corresponding IAM permission in a
// policy. Learn more
//   - Signing Amazon Web Services API Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html)
//   - Using the Amazon Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
func (c *Client) CreateGroup(ctx context.Context, params *CreateGroupInput, optFns ...func(*Options)) (*CreateGroupOutput, error) {
	if params == nil {
		params = &CreateGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGroup", params, optFns, c.addOperationCreateGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGroupInput struct {

	// The name of the group. Must be unique.
	//
	// This member is required.
	GroupName *string

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// A string containing the description of the group.
	Description *string

	// A non-negative integer value that specifies the precedence of this group
	// relative to the other groups that a user can belong to in the user pool. Zero is
	// the highest precedence value. Groups with lower Precedence values take
	// precedence over groups with higher or null Precedence values. If a user belongs
	// to two or more groups, it is the group with the lowest precedence value whose
	// role ARN is given in the user's tokens for the cognito:roles and
	// cognito:preferred_role claims. Two groups can have the same Precedence value.
	// If this happens, neither group takes precedence over the other. If two groups
	// with the same Precedence have the same role ARN, that role is used in the
	// cognito:preferred_role claim in tokens for users in each group. If the two
	// groups have different role ARNs, the cognito:preferred_role claim isn't set in
	// users' tokens. The default Precedence value is null. The maximum Precedence
	// value is 2^31-1 .
	Precedence *int32

	// The role Amazon Resource Name (ARN) for the group.
	RoleArn *string

	noSmithyDocumentSerde
}

type CreateGroupOutput struct {

	// The group object for the group.
	Group *types.GroupType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateGroup"); err != nil {
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
	if err = addOpCreateGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateGroup",
	}
}
