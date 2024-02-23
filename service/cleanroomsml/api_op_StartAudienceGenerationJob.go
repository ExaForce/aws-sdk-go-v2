// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanroomsml

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cleanroomsml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Information necessary to start the audience generation job.
func (c *Client) StartAudienceGenerationJob(ctx context.Context, params *StartAudienceGenerationJobInput, optFns ...func(*Options)) (*StartAudienceGenerationJobOutput, error) {
	if params == nil {
		params = &StartAudienceGenerationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartAudienceGenerationJob", params, optFns, c.addOperationStartAudienceGenerationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartAudienceGenerationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartAudienceGenerationJobInput struct {

	// The Amazon Resource Name (ARN) of the configured audience model that is used
	// for this audience generation job.
	//
	// This member is required.
	ConfiguredAudienceModelArn *string

	// The name of the audience generation job.
	//
	// This member is required.
	Name *string

	// The seed audience that is used to generate the audience.
	//
	// This member is required.
	SeedAudience *types.AudienceGenerationJobDataSource

	// The identifier of the collaboration that contains the audience generation job.
	CollaborationId *string

	// The description of the audience generation job.
	Description *string

	// Whether the seed audience is included in the audience generation output.
	IncludeSeedInOutput bool

	// The optional metadata that you apply to the resource to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of which
	// you define. The following basic restrictions apply to tags:
	//   - Maximum number of tags per resource - 50.
	//   - For each resource, each tag key must be unique, and each tag key can have
	//   only one value.
	//   - Maximum key length - 128 Unicode characters in UTF-8.
	//   - Maximum value length - 256 Unicode characters in UTF-8.
	//   - If your tagging schema is used across multiple services and resources,
	//   remember that other services may have restrictions on allowed characters.
	//   Generally allowed characters are: letters, numbers, and spaces representable in
	//   UTF-8, and the following characters: + - = . _ : / @.
	//   - Tag keys and values are case sensitive.
	//   - Do not use aws:, AWS:, or any upper or lowercase combination of such as a
	//   prefix for keys as it is reserved for AWS use. You cannot edit or delete tag
	//   keys with this prefix. Values can have this prefix. If a tag value has aws as
	//   its prefix but the key does not, then Forecast considers it to be a user tag and
	//   will count against the limit of 50 tags. Tags with only the key prefix of aws do
	//   not count against your tags per resource limit.
	Tags map[string]string

	noSmithyDocumentSerde
}

type StartAudienceGenerationJobOutput struct {

	// The Amazon Resource Name (ARN) of the audience generation job.
	//
	// This member is required.
	AudienceGenerationJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartAudienceGenerationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartAudienceGenerationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartAudienceGenerationJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartAudienceGenerationJob"); err != nil {
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
	if err = addOpStartAudienceGenerationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartAudienceGenerationJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartAudienceGenerationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartAudienceGenerationJob",
	}
}
