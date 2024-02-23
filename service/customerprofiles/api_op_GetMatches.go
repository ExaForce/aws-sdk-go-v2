// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Before calling this API, use CreateDomain (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_CreateDomain.html)
// or UpdateDomain (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UpdateDomain.html)
// to enable identity resolution: set Matching to true. GetMatches returns
// potentially matching profiles, based on the results of the latest run of a
// machine learning process. The process of matching duplicate profiles. If
// Matching = true , Amazon Connect Customer Profiles starts a weekly batch process
// called Identity Resolution Job. If you do not specify a date and time for
// Identity Resolution Job to run, by default it runs every Saturday at 12AM UTC to
// detect duplicate profiles in your domains. After the Identity Resolution Job
// completes, use the GetMatches (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html)
// API to return and review the results. Or, if you have configured ExportingConfig
// in the MatchingRequest , you can download the results from S3. Amazon Connect
// uses the following profile attributes to identify matches:
//   - PhoneNumber
//   - HomePhoneNumber
//   - BusinessPhoneNumber
//   - MobilePhoneNumber
//   - EmailAddress
//   - PersonalEmailAddress
//   - BusinessEmailAddress
//   - FullName
//
// For example, two or more profiles—with spelling mistakes such as John Doe and
// Jhn Doe, or different casing email addresses such as JOHN_DOE@ANYCOMPANY.COM and
// johndoe@anycompany.com, or different phone number formats such as 555-010-0000
// and +1-555-010-0000—can be detected as belonging to the same customer John Doe
// and merged into a unified profile.
func (c *Client) GetMatches(ctx context.Context, params *GetMatchesInput, optFns ...func(*Options)) (*GetMatchesOutput, error) {
	if params == nil {
		params = &GetMatchesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMatches", params, optFns, c.addOperationGetMatchesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMatchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMatchesInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetMatchesOutput struct {

	// The timestamp this version of Match Result generated.
	MatchGenerationDate *time.Time

	// The list of matched profiles for this instance.
	Matches []types.MatchItem

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// The number of potential matches found.
	PotentialMatches *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMatchesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMatches{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMatches{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMatches"); err != nil {
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
	if err = addOpGetMatchesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMatches(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMatches(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMatches",
	}
}
