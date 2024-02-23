// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Generates batch recommendations based on a list of items or users stored in
// Amazon S3 and exports the recommendations to an Amazon S3 bucket. To generate
// batch recommendations, specify the ARN of a solution version and an Amazon S3
// URI for the input and output data. For user personalization, popular items, and
// personalized ranking solutions, the batch inference job generates a list of
// recommended items for each user ID in the input file. For related items
// solutions, the job generates a list of recommended items for each item ID in the
// input file. For more information, see Creating a batch inference job  (https://docs.aws.amazon.com/personalize/latest/dg/getting-batch-recommendations.html)
// . If you use the Similar-Items recipe, Amazon Personalize can add descriptive
// themes to batch recommendations. To generate themes, set the job's mode to
// THEME_GENERATION and specify the name of the field that contains item names in
// the input data. For more information about generating themes, see Batch
// recommendations with themes from Content Generator  (https://docs.aws.amazon.com/personalize/latest/dg/themed-batch-recommendations.html)
// . You can't get batch recommendations with the Trending-Now or Next-Best-Action
// recipes.
func (c *Client) CreateBatchInferenceJob(ctx context.Context, params *CreateBatchInferenceJobInput, optFns ...func(*Options)) (*CreateBatchInferenceJobOutput, error) {
	if params == nil {
		params = &CreateBatchInferenceJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBatchInferenceJob", params, optFns, c.addOperationCreateBatchInferenceJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBatchInferenceJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBatchInferenceJobInput struct {

	// The Amazon S3 path that leads to the input file to base your recommendations
	// on. The input material must be in JSON format.
	//
	// This member is required.
	JobInput *types.BatchInferenceJobInput

	// The name of the batch inference job to create.
	//
	// This member is required.
	JobName *string

	// The path to the Amazon S3 bucket where the job's output will be stored.
	//
	// This member is required.
	JobOutput *types.BatchInferenceJobOutput

	// The ARN of the Amazon Identity and Access Management role that has permissions
	// to read and write to your input and output Amazon S3 buckets respectively.
	//
	// This member is required.
	RoleArn *string

	// The Amazon Resource Name (ARN) of the solution version that will be used to
	// generate the batch inference recommendations.
	//
	// This member is required.
	SolutionVersionArn *string

	// The configuration details of a batch inference job.
	BatchInferenceJobConfig *types.BatchInferenceJobConfig

	// The mode of the batch inference job. To generate descriptive themes for groups
	// of similar items, set the job mode to THEME_GENERATION . If you don't want to
	// generate themes, use the default BATCH_INFERENCE . When you get batch
	// recommendations with themes, you will incur additional costs. For more
	// information, see Amazon Personalize pricing (https://aws.amazon.com/personalize/pricing/)
	// .
	BatchInferenceJobMode types.BatchInferenceJobMode

	// The ARN of the filter to apply to the batch inference job. For more information
	// on using filters, see Filtering batch recommendations (https://docs.aws.amazon.com/personalize/latest/dg/filter-batch.html)
	// .
	FilterArn *string

	// The number of recommendations to retrieve.
	NumResults *int32

	// A list of tags (https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html)
	// to apply to the batch inference job.
	Tags []types.Tag

	// For theme generation jobs, specify the name of the column in your Items dataset
	// that contains each item's name.
	ThemeGenerationConfig *types.ThemeGenerationConfig

	noSmithyDocumentSerde
}

type CreateBatchInferenceJobOutput struct {

	// The ARN of the batch inference job.
	BatchInferenceJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBatchInferenceJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBatchInferenceJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBatchInferenceJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBatchInferenceJob"); err != nil {
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
	if err = addOpCreateBatchInferenceJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBatchInferenceJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBatchInferenceJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBatchInferenceJob",
	}
}
