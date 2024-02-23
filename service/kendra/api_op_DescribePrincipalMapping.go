// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the processing of PUT and DELETE actions for mapping users to their
// groups. This includes information on the status of actions currently processing
// or yet to be processed, when actions were last updated, when actions were
// received by Amazon Kendra, the latest action that should process and apply after
// other actions, and useful error messages if an action could not be processed.
// DescribePrincipalMapping is currently not supported in the Amazon Web Services
// GovCloud (US-West) region.
func (c *Client) DescribePrincipalMapping(ctx context.Context, params *DescribePrincipalMappingInput, optFns ...func(*Options)) (*DescribePrincipalMappingOutput, error) {
	if params == nil {
		params = &DescribePrincipalMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePrincipalMapping", params, optFns, c.addOperationDescribePrincipalMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePrincipalMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePrincipalMappingInput struct {

	// The identifier of the group required to check the processing of PUT and DELETE
	// actions for mapping users to their groups.
	//
	// This member is required.
	GroupId *string

	// The identifier of the index required to check the processing of PUT and DELETE
	// actions for mapping users to their groups.
	//
	// This member is required.
	IndexId *string

	// The identifier of the data source to check the processing of PUT and DELETE
	// actions for mapping users to their groups.
	DataSourceId *string

	noSmithyDocumentSerde
}

type DescribePrincipalMappingOutput struct {

	// Shows the identifier of the data source to see information on the processing of
	// PUT and DELETE actions for mapping users to their groups.
	DataSourceId *string

	// Shows the identifier of the group to see information on the processing of PUT
	// and DELETE actions for mapping users to their groups.
	GroupId *string

	// Shows the following information on the processing of PUT and DELETE actions for
	// mapping users to their groups:
	//   - Status—the status can be either PROCESSING , SUCCEEDED , DELETING , DELETED
	//   , or FAILED .
	//   - Last updated—the last date-time an action was updated.
	//   - Received—the last date-time an action was received or submitted.
	//   - Ordering ID—the latest action that should process and apply after other
	//   actions.
	//   - Failure reason—the reason an action could not be processed.
	GroupOrderingIdSummaries []types.GroupOrderingIdSummary

	// Shows the identifier of the index to see information on the processing of PUT
	// and DELETE actions for mapping users to their groups.
	IndexId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePrincipalMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePrincipalMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePrincipalMapping{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribePrincipalMapping"); err != nil {
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
	if err = addOpDescribePrincipalMappingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePrincipalMapping(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePrincipalMapping(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribePrincipalMapping",
	}
}
