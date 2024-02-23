// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of protected health information (PHI) detection jobs you have
// submitted.
func (c *Client) ListPHIDetectionJobs(ctx context.Context, params *ListPHIDetectionJobsInput, optFns ...func(*Options)) (*ListPHIDetectionJobsOutput, error) {
	if params == nil {
		params = &ListPHIDetectionJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPHIDetectionJobs", params, optFns, c.addOperationListPHIDetectionJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPHIDetectionJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPHIDetectionJobsInput struct {

	// Filters the jobs that are returned. You can filter jobs based on their names,
	// status, or the date and time that they were submitted. You can only set one
	// filter at a time.
	Filter *types.ComprehendMedicalAsyncJobFilter

	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPHIDetectionJobsOutput struct {

	// A list containing the properties of each job returned.
	ComprehendMedicalAsyncJobPropertiesList []types.ComprehendMedicalAsyncJobProperties

	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPHIDetectionJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPHIDetectionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPHIDetectionJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPHIDetectionJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPHIDetectionJobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListPHIDetectionJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPHIDetectionJobs",
	}
}
