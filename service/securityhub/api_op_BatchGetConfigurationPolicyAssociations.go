// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns associations between an Security Hub configuration and a batch of
// target accounts, organizational units, or the root. Only the Security Hub
// delegated administrator can invoke this operation from the home Region. A
// configuration can refer to a configuration policy or to a self-managed
// configuration.
func (c *Client) BatchGetConfigurationPolicyAssociations(ctx context.Context, params *BatchGetConfigurationPolicyAssociationsInput, optFns ...func(*Options)) (*BatchGetConfigurationPolicyAssociationsOutput, error) {
	if params == nil {
		params = &BatchGetConfigurationPolicyAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetConfigurationPolicyAssociations", params, optFns, c.addOperationBatchGetConfigurationPolicyAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetConfigurationPolicyAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetConfigurationPolicyAssociationsInput struct {

	// Specifies one or more target account IDs, organizational unit (OU) IDs, or the
	// root ID to retrieve associations for.
	//
	// This member is required.
	ConfigurationPolicyAssociationIdentifiers []types.ConfigurationPolicyAssociation

	noSmithyDocumentSerde
}

type BatchGetConfigurationPolicyAssociationsOutput struct {

	// Describes associations for the target accounts, OUs, or the root.
	ConfigurationPolicyAssociations []types.ConfigurationPolicyAssociationSummary

	// An array of configuration policy associations, one for each configuration
	// policy association identifier, that was specified in the request but couldn’t be
	// processed due to an error.
	UnprocessedConfigurationPolicyAssociations []types.UnprocessedConfigurationPolicyAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetConfigurationPolicyAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetConfigurationPolicyAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetConfigurationPolicyAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetConfigurationPolicyAssociations"); err != nil {
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
	if err = addOpBatchGetConfigurationPolicyAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetConfigurationPolicyAssociations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetConfigurationPolicyAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetConfigurationPolicyAssociations",
	}
}
