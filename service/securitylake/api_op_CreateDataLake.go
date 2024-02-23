// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initializes an Amazon Security Lake instance with the provided (or default)
// configuration. You can enable Security Lake in Amazon Web Services Regions with
// customized settings before enabling log collection in Regions. To specify
// particular Regions, configure these Regions using the configurations parameter.
// If you have already enabled Security Lake in a Region when you call this
// command, the command will update the Region if you provide new configuration
// parameters. If you have not already enabled Security Lake in the Region when you
// call this API, it will set up the data lake in the Region with the specified
// configurations. When you enable Security Lake, it starts ingesting security data
// after the CreateAwsLogSource call. This includes ingesting security data from
// sources, storing data, and making data accessible to subscribers. Security Lake
// also enables all the existing settings and resources that it stores or maintains
// for your Amazon Web Services account in the current Region, including security
// log and event data. For more information, see the Amazon Security Lake User
// Guide (https://docs.aws.amazon.com/security-lake/latest/userguide/what-is-security-lake.html)
// .
func (c *Client) CreateDataLake(ctx context.Context, params *CreateDataLakeInput, optFns ...func(*Options)) (*CreateDataLakeOutput, error) {
	if params == nil {
		params = &CreateDataLakeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataLake", params, optFns, c.addOperationCreateDataLakeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataLakeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataLakeInput struct {

	// Specify the Region or Regions that will contribute data to the rollup region.
	//
	// This member is required.
	Configurations []types.DataLakeConfiguration

	// The Amazon Resource Name (ARN) used to create and update the Glue table. This
	// table contains partitions generated by the ingestion and normalization of Amazon
	// Web Services log sources and custom sources.
	//
	// This member is required.
	MetaStoreManagerRoleArn *string

	// An array of objects, one for each tag to associate with the data lake
	// configuration. For each tag, you must specify both a tag key and a tag value. A
	// tag value cannot be null, but it can be an empty string.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDataLakeOutput struct {

	// The created Security Lake configuration object.
	DataLakes []types.DataLakeResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataLakeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataLake{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataLake{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDataLake"); err != nil {
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
	if err = addOpCreateDataLakeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataLake(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDataLake(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDataLake",
	}
}
