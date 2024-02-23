// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the structure of an approval rule created specifically for a pull
// request. For example, you can change the number of required approvers and the
// approval pool for approvers.
func (c *Client) UpdatePullRequestApprovalRuleContent(ctx context.Context, params *UpdatePullRequestApprovalRuleContentInput, optFns ...func(*Options)) (*UpdatePullRequestApprovalRuleContentOutput, error) {
	if params == nil {
		params = &UpdatePullRequestApprovalRuleContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePullRequestApprovalRuleContent", params, optFns, c.addOperationUpdatePullRequestApprovalRuleContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePullRequestApprovalRuleContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePullRequestApprovalRuleContentInput struct {

	// The name of the approval rule you want to update.
	//
	// This member is required.
	ApprovalRuleName *string

	// The updated content for the approval rule. When you update the content of the
	// approval rule, you can specify approvers in an approval pool in one of two ways:
	//
	//   - CodeCommitApprovers: This option only requires an Amazon Web Services
	//   account and a resource. It can be used for both IAM users and federated access
	//   users whose name matches the provided resource name. This is a very powerful
	//   option that offers a great deal of flexibility. For example, if you specify the
	//   Amazon Web Services account 123456789012 and Mary_Major, all of the following
	//   are counted as approvals coming from that user:
	//   - An IAM user in the account (arn:aws:iam::123456789012:user/Mary_Major)
	//   - A federated user identified in IAM as Mary_Major
	//   (arn:aws:sts::123456789012:federated-user/Mary_Major) This option does not
	//   recognize an active session of someone assuming the role of CodeCommitReview
	//   with a role session name of Mary_Major
	//   (arn:aws:sts::123456789012:assumed-role/CodeCommitReview/Mary_Major) unless you
	//   include a wildcard (*Mary_Major).
	//   - Fully qualified ARN: This option allows you to specify the fully qualified
	//   Amazon Resource Name (ARN) of the IAM user or role.
	// For more information about IAM ARNs, wildcards, and formats, see IAM Identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html)
	// in the IAM User Guide.
	//
	// This member is required.
	NewRuleContent *string

	// The system-generated ID of the pull request.
	//
	// This member is required.
	PullRequestId *string

	// The SHA-256 hash signature for the content of the approval rule. You can
	// retrieve this information by using GetPullRequest .
	ExistingRuleContentSha256 *string

	noSmithyDocumentSerde
}

type UpdatePullRequestApprovalRuleContentOutput struct {

	// Information about the updated approval rule.
	//
	// This member is required.
	ApprovalRule *types.ApprovalRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePullRequestApprovalRuleContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePullRequestApprovalRuleContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePullRequestApprovalRuleContent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdatePullRequestApprovalRuleContent"); err != nil {
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
	if err = addOpUpdatePullRequestApprovalRuleContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePullRequestApprovalRuleContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePullRequestApprovalRuleContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdatePullRequestApprovalRuleContent",
	}
}
