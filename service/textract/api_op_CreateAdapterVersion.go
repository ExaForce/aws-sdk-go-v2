// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new version of an adapter. Operates on a provided AdapterId and a
// specified dataset provided via the DatasetConfig argument. Requires that you
// specify an Amazon S3 bucket with the OutputConfig argument. You can provide an
// optional KMSKeyId, an optional ClientRequestToken, and optional tags.
func (c *Client) CreateAdapterVersion(ctx context.Context, params *CreateAdapterVersionInput, optFns ...func(*Options)) (*CreateAdapterVersionOutput, error) {
	if params == nil {
		params = &CreateAdapterVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAdapterVersion", params, optFns, c.addOperationCreateAdapterVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAdapterVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAdapterVersionInput struct {

	// A string containing a unique ID for the adapter that will receive a new version.
	//
	// This member is required.
	AdapterId *string

	// Specifies a dataset used to train a new adapter version. Takes a
	// ManifestS3Object as the value.
	//
	// This member is required.
	DatasetConfig *types.AdapterVersionDatasetConfig

	// Sets whether or not your output will go to a user created bucket. Used to set
	// the name of the bucket, and the prefix on the output file. OutputConfig is an
	// optional parameter which lets you adjust where your output will be placed. By
	// default, Amazon Textract will store the results internally and can only be
	// accessed by the Get API operations. With OutputConfig enabled, you can set the
	// name of the bucket the output will be sent to the file prefix of the results
	// where you can download your results. Additionally, you can set the KMSKeyID
	// parameter to a customer master key (CMK) to encrypt your output. Without this
	// parameter set Amazon Textract will encrypt server-side using the AWS managed CMK
	// for Amazon S3. Decryption of Customer Content is necessary for processing of the
	// documents by Amazon Textract. If your account is opted out under an AI services
	// opt out policy then all unencrypted Customer Content is immediately and
	// permanently deleted after the Customer Content has been processed by the
	// service. No copy of of the output is retained by Amazon Textract. For
	// information about how to opt out, see Managing AI services opt-out policy.  (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html)
	// For more information on data privacy, see the Data Privacy FAQ (https://aws.amazon.com/compliance/data-privacy-faq/)
	// .
	//
	// This member is required.
	OutputConfig *types.OutputConfig

	// Idempotent token is used to recognize the request. If the same token is used
	// with multiple CreateAdapterVersion requests, the same session is returned. This
	// token is employed to avoid unintentionally creating the same session multiple
	// times.
	ClientRequestToken *string

	// The identifier for your AWS Key Management Service key (AWS KMS key). Used to
	// encrypt your documents.
	KMSKeyId *string

	// A set of tags (key-value pairs) that you want to attach to the adapter version.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAdapterVersionOutput struct {

	// A string containing the unique ID for the adapter that has received a new
	// version.
	AdapterId *string

	// A string describing the new version of the adapter.
	AdapterVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAdapterVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAdapterVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAdapterVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAdapterVersion"); err != nil {
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
	if err = addIdempotencyToken_opCreateAdapterVersionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAdapterVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAdapterVersion(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAdapterVersion struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAdapterVersion) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAdapterVersion) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAdapterVersionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAdapterVersionInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAdapterVersionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAdapterVersion{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAdapterVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAdapterVersion",
	}
}
