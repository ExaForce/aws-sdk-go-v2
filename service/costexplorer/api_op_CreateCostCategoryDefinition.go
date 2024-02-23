// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Cost Category with the requested name and rules.
func (c *Client) CreateCostCategoryDefinition(ctx context.Context, params *CreateCostCategoryDefinitionInput, optFns ...func(*Options)) (*CreateCostCategoryDefinitionOutput, error) {
	if params == nil {
		params = &CreateCostCategoryDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCostCategoryDefinition", params, optFns, c.addOperationCreateCostCategoryDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCostCategoryDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCostCategoryDefinitionInput struct {

	// The unique name of the Cost Category.
	//
	// This member is required.
	Name *string

	// The rule schema version in this particular Cost Category.
	//
	// This member is required.
	RuleVersion types.CostCategoryRuleVersion

	// The Cost Category rules used to categorize costs. For more information, see
	// CostCategoryRule (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_CostCategoryRule.html)
	// .
	//
	// This member is required.
	Rules []types.CostCategoryRule

	// The default value for the cost category.
	DefaultValue *string

	// The Cost Category's effective start date. It can only be a billing start date
	// (first day of the month). If the date isn't provided, it's the first day of the
	// current month. Dates can't be before the previous twelve months, or in the
	// future.
	EffectiveStart *string

	// An optional list of tags to associate with the specified CostCategory (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_CostCategory.html)
	// . You can use resource tags to control access to your cost category using IAM
	// policies. Each tag consists of a key and a value, and each key must be unique
	// for the resource. The following restrictions apply to resource tags:
	//   - Although the maximum number of array members is 200, you can assign a
	//   maximum of 50 user-tags to one resource. The remaining are reserved for Amazon
	//   Web Services use
	//   - The maximum length of a key is 128 characters
	//   - The maximum length of a value is 256 characters
	//   - Keys and values can only contain alphanumeric characters, spaces, and any
	//   of the following: _.:/=+@-
	//   - Keys and values are case sensitive
	//   - Keys and values are trimmed for any leading or trailing whitespaces
	//   - Don’t use aws: as a prefix for your keys. This prefix is reserved for Amazon
	//   Web Services use
	ResourceTags []types.ResourceTag

	// The split charge rules used to allocate your charges between your Cost Category
	// values.
	SplitChargeRules []types.CostCategorySplitChargeRule

	noSmithyDocumentSerde
}

type CreateCostCategoryDefinitionOutput struct {

	// The unique identifier for your newly created Cost Category.
	CostCategoryArn *string

	// The Cost Category's effective start date. It can only be a billing start date
	// (first day of the month).
	EffectiveStart *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCostCategoryDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCostCategoryDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCostCategoryDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCostCategoryDefinition"); err != nil {
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
	if err = addOpCreateCostCategoryDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCostCategoryDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCostCategoryDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCostCategoryDefinition",
	}
}
