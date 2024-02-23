// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Geocodes free-form text, such as an address, name, city, or region to allow you
// to search for Places or points of interest. Optional parameters let you narrow
// your search results by bounding box or country, or bias your search toward a
// specific position on the globe. You can search for places near a given position
// using BiasPosition , or filter results within a bounding box using FilterBBox .
// Providing both parameters simultaneously returns an error. Search results are
// returned in order of highest to lowest relevance.
func (c *Client) SearchPlaceIndexForText(ctx context.Context, params *SearchPlaceIndexForTextInput, optFns ...func(*Options)) (*SearchPlaceIndexForTextOutput, error) {
	if params == nil {
		params = &SearchPlaceIndexForTextInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchPlaceIndexForText", params, optFns, c.addOperationSearchPlaceIndexForTextMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchPlaceIndexForTextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchPlaceIndexForTextInput struct {

	// The name of the place index resource you want to use for the search.
	//
	// This member is required.
	IndexName *string

	// The address, name, city, or region to be used in the search in free-form text
	// format. For example, 123 Any Street .
	//
	// This member is required.
	Text *string

	// An optional parameter that indicates a preference for places that are closer to
	// a specified position. If provided, this parameter must contain a pair of
	// numbers. The first number represents the X coordinate, or longitude; the second
	// number represents the Y coordinate, or latitude. For example, [-123.1174,
	// 49.2847] represents the position with longitude -123.1174 and latitude 49.2847 .
	// BiasPosition and FilterBBox are mutually exclusive. Specifying both options
	// results in an error.
	BiasPosition []float64

	// An optional parameter that limits the search results by returning only places
	// that are within the provided bounding box. If provided, this parameter must
	// contain a total of four consecutive numbers in two pairs. The first pair of
	// numbers represents the X and Y coordinates (longitude and latitude,
	// respectively) of the southwest corner of the bounding box; the second pair of
	// numbers represents the X and Y coordinates (longitude and latitude,
	// respectively) of the northeast corner of the bounding box. For example,
	// [-12.7935, -37.4835, -12.0684, -36.9542] represents a bounding box where the
	// southwest corner has longitude -12.7935 and latitude -37.4835 , and the
	// northeast corner has longitude -12.0684 and latitude -36.9542 . FilterBBox and
	// BiasPosition are mutually exclusive. Specifying both options results in an error.
	FilterBBox []float64

	// A list of one or more Amazon Location categories to filter the returned places.
	// If you include more than one category, the results will include results that
	// match any of the categories listed. For more information about using categories,
	// including a list of Amazon Location categories, see Categories and filtering (https://docs.aws.amazon.com/location/latest/developerguide/category-filtering.html)
	// , in the Amazon Location Service Developer Guide.
	FilterCategories []string

	// An optional parameter that limits the search results by returning only places
	// that are in a specified list of countries.
	//   - Valid values include ISO 3166 (https://www.iso.org/iso-3166-country-codes.html)
	//   3-digit country codes. For example, Australia uses three upper-case characters:
	//   AUS .
	FilterCountries []string

	// The optional API key (https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html)
	// to authorize the request.
	Key *string

	// The preferred language used to return results. The value must be a valid BCP 47 (https://tools.ietf.org/search/bcp47)
	// language tag, for example, en for English. This setting affects the languages
	// used in the results, but not the results themselves. If no language is
	// specified, or not supported for a particular result, the partner automatically
	// chooses a language for the result. For an example, we'll use the Greek language.
	// You search for Athens, Greece , with the language parameter set to en . The
	// result found will most likely be returned as Athens . If you set the language
	// parameter to el , for Greek, then the result found will more likely be returned
	// as Αθήνα . If the data provider does not have a value for Greek, the result will
	// be in a language that the provider does support.
	Language *string

	// An optional parameter. The maximum number of results returned per request. The
	// default: 50
	MaxResults *int32

	noSmithyDocumentSerde
}

type SearchPlaceIndexForTextOutput struct {

	// A list of Places matching the input text. Each result contains additional
	// information about the specific point of interest. Not all response properties
	// are included with all responses. Some properties may only be returned by
	// specific data partners.
	//
	// This member is required.
	Results []types.SearchForTextResult

	// Contains a summary of the request. Echoes the input values for BiasPosition ,
	// FilterBBox , FilterCountries , Language , MaxResults , and Text . Also includes
	// the DataSource of the place index and the bounding box, ResultBBox , which
	// surrounds the search results.
	//
	// This member is required.
	Summary *types.SearchPlaceIndexForTextSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchPlaceIndexForTextMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchPlaceIndexForText{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchPlaceIndexForText{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchPlaceIndexForText"); err != nil {
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
	if err = addEndpointPrefix_opSearchPlaceIndexForTextMiddleware(stack); err != nil {
		return err
	}
	if err = addOpSearchPlaceIndexForTextValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchPlaceIndexForText(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opSearchPlaceIndexForTextMiddleware struct {
}

func (*endpointPrefix_opSearchPlaceIndexForTextMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opSearchPlaceIndexForTextMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "places." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opSearchPlaceIndexForTextMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opSearchPlaceIndexForTextMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opSearchPlaceIndexForText(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchPlaceIndexForText",
	}
}
