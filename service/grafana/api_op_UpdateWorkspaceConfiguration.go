// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the configuration string for the given workspace
func (c *Client) UpdateWorkspaceConfiguration(ctx context.Context, params *UpdateWorkspaceConfigurationInput, optFns ...func(*Options)) (*UpdateWorkspaceConfigurationOutput, error) {
	if params == nil {
		params = &UpdateWorkspaceConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWorkspaceConfiguration", params, optFns, c.addOperationUpdateWorkspaceConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWorkspaceConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWorkspaceConfigurationInput struct {

	// The new configuration string for the workspace. For more information about the
	// format and configuration options available, see Working in your Grafana
	// workspace (https://docs.aws.amazon.com/grafana/latest/userguide/AMG-configure-workspace.html)
	// .
	//
	// This value conforms to the media type: application/json
	//
	// This member is required.
	Configuration *string

	// The ID of the workspace to update.
	//
	// This member is required.
	WorkspaceId *string

	// Specifies the version of Grafana to support in the new workspace. Can only be
	// used to upgrade (for example, from 8.4 to 9.4), not downgrade (for example, from
	// 9.4 to 8.4). To know what versions are available to upgrade to for a specific
	// workspace, see the ListVersions operation.
	GrafanaVersion *string

	noSmithyDocumentSerde
}

type UpdateWorkspaceConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWorkspaceConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateWorkspaceConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateWorkspaceConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateWorkspaceConfiguration"); err != nil {
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
	if err = addOpUpdateWorkspaceConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWorkspaceConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWorkspaceConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateWorkspaceConfiguration",
	}
}
