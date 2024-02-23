// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of the available releases for the mobile SDK and the specified
// device platform. The mobile SDK is not generally available. Customers who have
// access to the mobile SDK can use it to establish and manage WAF tokens for use
// in HTTP(S) requests from a mobile device to WAF. For more information, see WAF
// client application integration (https://docs.aws.amazon.com/waf/latest/developerguide/waf-application-integration.html)
// in the WAF Developer Guide.
func (c *Client) ListMobileSdkReleases(ctx context.Context, params *ListMobileSdkReleasesInput, optFns ...func(*Options)) (*ListMobileSdkReleasesOutput, error) {
	if params == nil {
		params = &ListMobileSdkReleasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMobileSdkReleases", params, optFns, c.addOperationListMobileSdkReleasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMobileSdkReleasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMobileSdkReleasesInput struct {

	// The device platform to retrieve the list for.
	//
	// This member is required.
	Platform types.Platform

	// The maximum number of objects that you want WAF to return for this request. If
	// more objects are available, in the response, WAF provides a NextMarker value
	// that you can use in a subsequent call to get the next batch of objects.
	Limit *int32

	// When you request a list of objects with a Limit setting, if the number of
	// objects that are still available for retrieval exceeds the limit, WAF returns a
	// NextMarker value in the response. To retrieve the next batch of objects, provide
	// the marker from the prior call in your next request.
	NextMarker *string

	noSmithyDocumentSerde
}

type ListMobileSdkReleasesOutput struct {

	// When you request a list of objects with a Limit setting, if the number of
	// objects that are still available for retrieval exceeds the limit, WAF returns a
	// NextMarker value in the response. To retrieve the next batch of objects, provide
	// the marker from the prior call in your next request.
	NextMarker *string

	// The high level information for the available SDK releases. If you specified a
	// Limit in your request, this might not be the full list.
	ReleaseSummaries []types.ReleaseSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMobileSdkReleasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMobileSdkReleases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMobileSdkReleases{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMobileSdkReleases"); err != nil {
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
	if err = addOpListMobileSdkReleasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMobileSdkReleases(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListMobileSdkReleases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMobileSdkReleases",
	}
}
