// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the number of resources that are compliant and the number that are
// noncompliant. You can specify one or more resource types to get these numbers
// for each resource type. The maximum number returned is 100.
func (c *Client) GetComplianceSummaryByResourceType(ctx context.Context, params *GetComplianceSummaryByResourceTypeInput, optFns ...func(*Options)) (*GetComplianceSummaryByResourceTypeOutput, error) {
	if params == nil {
		params = &GetComplianceSummaryByResourceTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetComplianceSummaryByResourceType", params, optFns, c.addOperationGetComplianceSummaryByResourceTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetComplianceSummaryByResourceTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetComplianceSummaryByResourceTypeInput struct {

	// Specify one or more resource types to get the number of resources that are
	// compliant and the number that are noncompliant for each resource type. For this
	// request, you can specify an Amazon Web Services resource type such as
	// AWS::EC2::Instance . You can specify that the resource type is an Amazon Web
	// Services account by specifying AWS::::Account .
	ResourceTypes []string

	noSmithyDocumentSerde
}

type GetComplianceSummaryByResourceTypeOutput struct {

	// The number of resources that are compliant and the number that are
	// noncompliant. If one or more resource types were provided with the request, the
	// numbers are returned for each resource type. The maximum number returned is 100.
	ComplianceSummariesByResourceType []types.ComplianceSummaryByResourceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetComplianceSummaryByResourceTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetComplianceSummaryByResourceType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetComplianceSummaryByResourceType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetComplianceSummaryByResourceType"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetComplianceSummaryByResourceType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetComplianceSummaryByResourceType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetComplianceSummaryByResourceType",
	}
}
