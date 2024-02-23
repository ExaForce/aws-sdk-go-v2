// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the criteria and other settings for a findings filter.
func (c *Client) GetFindingsFilter(ctx context.Context, params *GetFindingsFilterInput, optFns ...func(*Options)) (*GetFindingsFilterOutput, error) {
	if params == nil {
		params = &GetFindingsFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFindingsFilter", params, optFns, c.addOperationGetFindingsFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFindingsFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFindingsFilterInput struct {

	// The unique identifier for the Amazon Macie resource that the request applies to.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetFindingsFilterOutput struct {

	// The action that's performed on findings that match the filter criteria
	// (findingCriteria). Possible values are: ARCHIVE, suppress (automatically
	// archive) the findings; and, NOOP, don't perform any action on the findings.
	Action types.FindingsFilterAction

	// The Amazon Resource Name (ARN) of the filter.
	Arn *string

	// The custom description of the filter.
	Description *string

	// The criteria that's used to filter findings.
	FindingCriteria *types.FindingCriteria

	// The unique identifier for the filter.
	Id *string

	// The custom name of the filter.
	Name *string

	// The position of the filter in the list of saved filters on the Amazon Macie
	// console. This value also determines the order in which the filter is applied to
	// findings, relative to other filters that are also applied to the findings.
	Position *int32

	// A map of key-value pairs that specifies which tags (keys and values) are
	// associated with the filter.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFindingsFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFindingsFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFindingsFilter{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetFindingsFilter"); err != nil {
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
	if err = addOpGetFindingsFilterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFindingsFilter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFindingsFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetFindingsFilter",
	}
}
