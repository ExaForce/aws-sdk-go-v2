// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides high-level information for a managed rule group, including
// descriptions of the rules.
func (c *Client) DescribeManagedRuleGroup(ctx context.Context, params *DescribeManagedRuleGroupInput, optFns ...func(*Options)) (*DescribeManagedRuleGroupOutput, error) {
	if params == nil {
		params = &DescribeManagedRuleGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeManagedRuleGroup", params, optFns, c.addOperationDescribeManagedRuleGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeManagedRuleGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeManagedRuleGroupInput struct {

	// The name of the managed rule group. You use this, along with the vendor name,
	// to identify the rule group.
	//
	// This member is required.
	Name *string

	// Specifies whether this is for an Amazon CloudFront distribution or for a
	// regional application. A regional application can be an Application Load Balancer
	// (ALB), an Amazon API Gateway REST API, an AppSync GraphQL API, an Amazon Cognito
	// user pool, an App Runner service, or an Amazon Web Services Verified Access
	// instance. To work with CloudFront, you must also specify the Region US East (N.
	// Virginia) as follows:
	//   - CLI - Specify the Region when you use the CloudFront scope:
	//   --scope=CLOUDFRONT --region=us-east-1 .
	//   - API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// The name of the managed rule group vendor. You use this, along with the rule
	// group name, to identify a rule group.
	//
	// This member is required.
	VendorName *string

	// The version of the rule group. You can only use a version that is not scheduled
	// for expiration. If you don't provide this, WAF uses the vendor's default
	// version.
	VersionName *string

	noSmithyDocumentSerde
}

type DescribeManagedRuleGroupOutput struct {

	// The labels that one or more rules in this rule group add to matching web
	// requests. These labels are defined in the RuleLabels for a Rule .
	AvailableLabels []types.LabelSummary

	// The web ACL capacity units (WCUs) required for this rule group. WAF uses WCUs
	// to calculate and control the operating resources that are used to run your
	// rules, rule groups, and web ACLs. WAF calculates capacity differently for each
	// rule type, to reflect the relative cost of each rule. Simple rules that cost
	// little to run use fewer WCUs than more complex rules that use more processing
	// power. Rule group capacity is fixed at creation, which helps users plan their
	// web ACL WCU usage when they use a rule group. For more information, see WAF web
	// ACL capacity units (WCU) (https://docs.aws.amazon.com/waf/latest/developerguide/aws-waf-capacity-units.html)
	// in the WAF Developer Guide.
	Capacity *int64

	// The labels that one or more rules in this rule group match against in label
	// match statements. These labels are defined in a LabelMatchStatement
	// specification, in the Statement definition of a rule.
	ConsumedLabels []types.LabelSummary

	// The label namespace prefix for this rule group. All labels added by rules in
	// this rule group have this prefix.
	//   - The syntax for the label namespace prefix for a managed rule group is the
	//   following: awswaf:managed:: :
	//   - When a rule with a label matches a web request, WAF adds the fully
	//   qualified label to the request. A fully qualified label is made up of the label
	//   namespace from the rule group or web ACL where the rule is defined and the label
	//   from the rule, separated by a colon: :
	LabelNamespace *string

	//
	Rules []types.RuleSummary

	// The Amazon resource name (ARN) of the Amazon Simple Notification Service SNS
	// topic that's used to provide notification of changes to the managed rule group.
	// You can subscribe to the SNS topic to receive notifications when the managed
	// rule group is modified, such as for new versions and for version expiration. For
	// more information, see the Amazon Simple Notification Service Developer Guide (https://docs.aws.amazon.com/sns/latest/dg/welcome.html)
	// .
	SnsTopicArn *string

	// The managed rule group's version.
	VersionName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeManagedRuleGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeManagedRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeManagedRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeManagedRuleGroup"); err != nil {
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
	if err = addOpDescribeManagedRuleGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeManagedRuleGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeManagedRuleGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeManagedRuleGroup",
	}
}
