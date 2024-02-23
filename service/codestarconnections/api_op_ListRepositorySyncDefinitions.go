// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarconnections

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the repository sync definitions for repository links in your account.
func (c *Client) ListRepositorySyncDefinitions(ctx context.Context, params *ListRepositorySyncDefinitionsInput, optFns ...func(*Options)) (*ListRepositorySyncDefinitionsOutput, error) {
	if params == nil {
		params = &ListRepositorySyncDefinitionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRepositorySyncDefinitions", params, optFns, c.addOperationListRepositorySyncDefinitionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRepositorySyncDefinitionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRepositorySyncDefinitionsInput struct {

	// The ID of the repository link for the sync definition for which you want to
	// retrieve information.
	//
	// This member is required.
	RepositoryLinkId *string

	// The sync type of the repository link for the the sync definition for which you
	// want to retrieve information.
	//
	// This member is required.
	SyncType types.SyncConfigurationType

	noSmithyDocumentSerde
}

type ListRepositorySyncDefinitionsOutput struct {

	// The list of repository sync definitions returned by the request. A
	// RepositorySyncDefinition is a mapping from a repository branch to all the Amazon
	// Web Services resources that are being synced from that branch.
	//
	// This member is required.
	RepositorySyncDefinitions []types.RepositorySyncDefinition

	// An enumeration token that, when provided in a request, returns the next batch
	// of the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRepositorySyncDefinitionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListRepositorySyncDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListRepositorySyncDefinitions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRepositorySyncDefinitions"); err != nil {
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
	if err = addOpListRepositorySyncDefinitionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRepositorySyncDefinitions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListRepositorySyncDefinitions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRepositorySyncDefinitions",
	}
}
