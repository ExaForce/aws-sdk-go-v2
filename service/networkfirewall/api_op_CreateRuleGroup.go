// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the specified stateless or stateful rule group, which includes the
// rules for network traffic inspection, a capacity setting, and tags. You provide
// your rule group specification in your request using either RuleGroup or Rules .
func (c *Client) CreateRuleGroup(ctx context.Context, params *CreateRuleGroupInput, optFns ...func(*Options)) (*CreateRuleGroupOutput, error) {
	if params == nil {
		params = &CreateRuleGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRuleGroup", params, optFns, c.addOperationCreateRuleGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRuleGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRuleGroupInput struct {

	// The maximum operating resources that this rule group can use. Rule group
	// capacity is fixed at creation. When you update a rule group, you are limited to
	// this capacity. When you reference a rule group from a firewall policy, Network
	// Firewall reserves this capacity for the rule group. You can retrieve the
	// capacity that would be required for a rule group before you create the rule
	// group by calling CreateRuleGroup with DryRun set to TRUE . You can't change or
	// exceed this capacity when you update the rule group, so leave room for your rule
	// group to grow. Capacity for a stateless rule group For a stateless rule group,
	// the capacity required is the sum of the capacity requirements of the individual
	// rules that you expect to have in the rule group. To calculate the capacity
	// requirement of a single rule, multiply the capacity requirement values of each
	// of the rule's match settings:
	//   - A match setting with no criteria specified has a value of 1.
	//   - A match setting with Any specified has a value of 1.
	//   - All other match settings have a value equal to the number of elements
	//   provided in the setting. For example, a protocol setting ["UDP"] and a source
	//   setting ["10.0.0.0/24"] each have a value of 1. A protocol setting ["UDP","TCP"]
	//   has a value of 2. A source setting ["10.0.0.0/24","10.0.0.1/24","10.0.0.2/24"]
	//   has a value of 3.
	// A rule with no criteria specified in any of its match settings has a capacity
	// requirement of 1. A rule with protocol setting ["UDP","TCP"], source setting
	// ["10.0.0.0/24","10.0.0.1/24","10.0.0.2/24"], and a single specification or no
	// specification for each of the other match settings has a capacity requirement of
	// 6. Capacity for a stateful rule group For a stateful rule group, the minimum
	// capacity required is the number of individual rules that you expect to have in
	// the rule group.
	//
	// This member is required.
	Capacity *int32

	// The descriptive name of the rule group. You can't change the name of a rule
	// group after you create it.
	//
	// This member is required.
	RuleGroupName *string

	// Indicates whether the rule group is stateless or stateful. If the rule group is
	// stateless, it contains stateless rules. If it is stateful, it contains stateful
	// rules.
	//
	// This member is required.
	Type types.RuleGroupType

	// Indicates whether you want Network Firewall to analyze the stateless rules in
	// the rule group for rule behavior such as asymmetric routing. If set to TRUE ,
	// Network Firewall runs the analysis and then creates the rule group for you. To
	// run the stateless rule group analyzer without creating the rule group, set
	// DryRun to TRUE .
	AnalyzeRuleGroup bool

	// A description of the rule group.
	Description *string

	// Indicates whether you want Network Firewall to just check the validity of the
	// request, rather than run the request. If set to TRUE , Network Firewall checks
	// whether the request can run successfully, but doesn't actually make the
	// requested changes. The call returns the value that the request would return if
	// you ran it with dry run set to FALSE , but doesn't make additions or changes to
	// your resources. This option allows you to make sure that you have the required
	// permissions to run the request and that your request parameters are valid. If
	// set to FALSE , Network Firewall makes the requested changes to your resources.
	DryRun bool

	// A complex type that contains settings for encryption of your rule group
	// resources.
	EncryptionConfiguration *types.EncryptionConfiguration

	// An object that defines the rule group rules. You must provide either this rule
	// group setting or a Rules setting, but not both.
	RuleGroup *types.RuleGroup

	// A string containing stateful rule group rules specifications in Suricata flat
	// format, with one rule per line. Use this to import your existing Suricata
	// compatible rule groups. You must provide either this rules setting or a
	// populated RuleGroup setting, but not both. You can provide your rule group
	// specification in Suricata flat format through this setting when you create or
	// update your rule group. The call response returns a RuleGroup object that
	// Network Firewall has populated from your string.
	Rules *string

	// A complex type that contains metadata about the rule group that your own rule
	// group is copied from. You can use the metadata to keep track of updates made to
	// the originating rule group.
	SourceMetadata *types.SourceMetadata

	// The key:value pairs to associate with the resource.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateRuleGroupOutput struct {

	// The high-level properties of a rule group. This, along with the RuleGroup ,
	// define the rule group. You can retrieve all objects for a rule group by calling
	// DescribeRuleGroup .
	//
	// This member is required.
	RuleGroupResponse *types.RuleGroupResponse

	// A token used for optimistic locking. Network Firewall returns a token to your
	// requests that access the rule group. The token marks the state of the rule group
	// resource at the time of the request. To make changes to the rule group, you
	// provide the token in your request. Network Firewall uses the token to ensure
	// that the rule group hasn't changed since you last retrieved it. If it has
	// changed, the operation fails with an InvalidTokenException . If this happens,
	// retrieve the rule group again to get a current copy of it with a current token.
	// Reapply your changes as needed, then try the operation again using the new
	// token.
	//
	// This member is required.
	UpdateToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRuleGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRuleGroup"); err != nil {
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
	if err = addOpCreateRuleGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRuleGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRuleGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRuleGroup",
	}
}
