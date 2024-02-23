// Code generated by smithy-go-codegen DO NOT EDIT.

package chatbot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the Microsoft Teams team authorization allowing for channels to be
// configured in that Microsoft Teams team. Note that the Microsoft Teams team must
// have no channels configured to remove it.
func (c *Client) DeleteMicrosoftTeamsConfiguredTeam(ctx context.Context, params *DeleteMicrosoftTeamsConfiguredTeamInput, optFns ...func(*Options)) (*DeleteMicrosoftTeamsConfiguredTeamOutput, error) {
	if params == nil {
		params = &DeleteMicrosoftTeamsConfiguredTeamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMicrosoftTeamsConfiguredTeam", params, optFns, c.addOperationDeleteMicrosoftTeamsConfiguredTeamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMicrosoftTeamsConfiguredTeamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteMicrosoftTeamsConfiguredTeamInput struct {

	// The ID of the Microsoft Team authorized with AWS Chatbot. To get the team ID,
	// you must perform the initial authorization flow with Microsoft Teams in the AWS
	// Chatbot console. Then you can copy and paste the team ID from the console. For
	// more details, see steps 1-4 in Get started with Microsoft Teams in the AWS
	// Chatbot Administrator Guide.
	//
	// This member is required.
	TeamId *string

	noSmithyDocumentSerde
}

type DeleteMicrosoftTeamsConfiguredTeamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteMicrosoftTeamsConfiguredTeamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteMicrosoftTeamsConfiguredTeam{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteMicrosoftTeamsConfiguredTeam{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteMicrosoftTeamsConfiguredTeam"); err != nil {
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
	if err = addOpDeleteMicrosoftTeamsConfiguredTeamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMicrosoftTeamsConfiguredTeam(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMicrosoftTeamsConfiguredTeam(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteMicrosoftTeamsConfiguredTeam",
	}
}
