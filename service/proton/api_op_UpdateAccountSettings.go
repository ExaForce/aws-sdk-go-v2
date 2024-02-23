// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update Proton settings that are used for multiple services in the Amazon Web
// Services account.
func (c *Client) UpdateAccountSettings(ctx context.Context, params *UpdateAccountSettingsInput, optFns ...func(*Options)) (*UpdateAccountSettingsOutput, error) {
	if params == nil {
		params = &UpdateAccountSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccountSettings", params, optFns, c.addOperationUpdateAccountSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccountSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAccountSettingsInput struct {

	// Set to true to remove a configured pipeline repository from the account
	// settings. Don't set this field if you are updating the configured pipeline
	// repository.
	DeletePipelineProvisioningRepository *bool

	// The Amazon Resource Name (ARN) of the service role you want to use for
	// provisioning pipelines. Proton assumes this role for CodeBuild-based
	// provisioning.
	PipelineCodebuildRoleArn *string

	// A linked repository for pipeline provisioning. Specify it if you have
	// environments configured for self-managed provisioning with services that include
	// pipelines. A linked repository is a repository that has been registered with
	// Proton. For more information, see CreateRepository . To remove a previously
	// configured repository, set deletePipelineProvisioningRepository to true , and
	// don't set pipelineProvisioningRepository .
	PipelineProvisioningRepository *types.RepositoryBranchInput

	// The Amazon Resource Name (ARN) of the service role you want to use for
	// provisioning pipelines. Assumed by Proton for Amazon Web Services-managed
	// provisioning, and by customer-owned automation for self-managed provisioning. To
	// remove a previously configured ARN, specify an empty string.
	PipelineServiceRoleArn *string

	noSmithyDocumentSerde
}

type UpdateAccountSettingsOutput struct {

	// The Proton pipeline service role and repository data shared across the Amazon
	// Web Services account.
	//
	// This member is required.
	AccountSettings *types.AccountSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAccountSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateAccountSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateAccountSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAccountSettings"); err != nil {
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
	if err = addOpUpdateAccountSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccountSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAccountSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAccountSettings",
	}
}
