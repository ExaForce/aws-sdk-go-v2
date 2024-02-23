// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation applies only to Amazon Rekognition Custom Labels. Attaches a
// project policy to a Amazon Rekognition Custom Labels project in a trusting AWS
// account. A project policy specifies that a trusted AWS account can copy a model
// version from a trusting AWS account to a project in the trusted AWS account. To
// copy a model version you use the CopyProjectVersion operation. Only applies to
// Custom Labels projects. For more information about the format of a project
// policy document, see Attaching a project policy (SDK) in the Amazon Rekognition
// Custom Labels Developer Guide. The response from PutProjectPolicy is a revision
// ID for the project policy. You can attach multiple project policies to a
// project. You can also update an existing project policy by specifying the policy
// revision ID of the existing policy. To remove a project policy from a project,
// call DeleteProjectPolicy . To get a list of project policies attached to a
// project, call ListProjectPolicies . You copy a model version by calling
// CopyProjectVersion . This operation requires permissions to perform the
// rekognition:PutProjectPolicy action.
func (c *Client) PutProjectPolicy(ctx context.Context, params *PutProjectPolicyInput, optFns ...func(*Options)) (*PutProjectPolicyOutput, error) {
	if params == nil {
		params = &PutProjectPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutProjectPolicy", params, optFns, c.addOperationPutProjectPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutProjectPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutProjectPolicyInput struct {

	// A resource policy to add to the model. The policy is a JSON structure that
	// contains one or more statements that define the policy. The policy must follow
	// the IAM syntax. For more information about the contents of a JSON policy
	// document, see IAM JSON policy reference (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html)
	// .
	//
	// This member is required.
	PolicyDocument *string

	// A name for the policy.
	//
	// This member is required.
	PolicyName *string

	// The Amazon Resource Name (ARN) of the project that the project policy is
	// attached to.
	//
	// This member is required.
	ProjectArn *string

	// The revision ID for the Project Policy. Each time you modify a policy, Amazon
	// Rekognition Custom Labels generates and assigns a new PolicyRevisionId and then
	// deletes the previous version of the policy.
	PolicyRevisionId *string

	noSmithyDocumentSerde
}

type PutProjectPolicyOutput struct {

	// The ID of the project policy.
	PolicyRevisionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutProjectPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutProjectPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutProjectPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutProjectPolicy"); err != nil {
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
	if err = addOpPutProjectPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutProjectPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutProjectPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutProjectPolicy",
	}
}
