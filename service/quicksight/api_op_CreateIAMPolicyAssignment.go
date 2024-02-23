// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an assignment with one specified IAM policy, identified by its Amazon
// Resource Name (ARN). This policy assignment is attached to the specified groups
// or users of Amazon QuickSight. Assignment names are unique per Amazon Web
// Services account. To avoid overwriting rules in other namespaces, use assignment
// names that are unique.
func (c *Client) CreateIAMPolicyAssignment(ctx context.Context, params *CreateIAMPolicyAssignmentInput, optFns ...func(*Options)) (*CreateIAMPolicyAssignmentOutput, error) {
	if params == nil {
		params = &CreateIAMPolicyAssignmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIAMPolicyAssignment", params, optFns, c.addOperationCreateIAMPolicyAssignmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIAMPolicyAssignmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIAMPolicyAssignmentInput struct {

	// The name of the assignment, also called a rule. The name must be unique within
	// the Amazon Web Services account.
	//
	// This member is required.
	AssignmentName *string

	// The status of the assignment. Possible values are as follows:
	//   - ENABLED - Anything specified in this assignment is used when creating the
	//   data source.
	//   - DISABLED - This assignment isn't used when creating the data source.
	//   - DRAFT - This assignment is an unfinished draft and isn't used when creating
	//   the data source.
	//
	// This member is required.
	AssignmentStatus types.AssignmentStatus

	// The ID of the Amazon Web Services account where you want to assign an IAM
	// policy to Amazon QuickSight users or groups.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace that contains the assignment.
	//
	// This member is required.
	Namespace *string

	// The Amazon QuickSight users, groups, or both that you want to assign the policy
	// to.
	Identities map[string][]string

	// The ARN for the IAM policy to apply to the Amazon QuickSight users and groups
	// specified in this assignment.
	PolicyArn *string

	noSmithyDocumentSerde
}

type CreateIAMPolicyAssignmentOutput struct {

	// The ID for the assignment.
	AssignmentId *string

	// The name of the assignment. The name must be unique within the Amazon Web
	// Services account.
	AssignmentName *string

	// The status of the assignment. Possible values are as follows:
	//   - ENABLED - Anything specified in this assignment is used when creating the
	//   data source.
	//   - DISABLED - This assignment isn't used when creating the data source.
	//   - DRAFT - This assignment is an unfinished draft and isn't used when creating
	//   the data source.
	AssignmentStatus types.AssignmentStatus

	// The Amazon QuickSight users, groups, or both that the IAM policy is assigned to.
	Identities map[string][]string

	// The ARN for the IAM policy that is applied to the Amazon QuickSight users and
	// groups specified in this assignment.
	PolicyArn *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIAMPolicyAssignmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateIAMPolicyAssignment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIAMPolicyAssignment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateIAMPolicyAssignment"); err != nil {
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
	if err = addOpCreateIAMPolicyAssignmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIAMPolicyAssignment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateIAMPolicyAssignment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateIAMPolicyAssignment",
	}
}
