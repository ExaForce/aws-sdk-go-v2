// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets an Amazon DataZone data source.
func (c *Client) GetDataSource(ctx context.Context, params *GetDataSourceInput, optFns ...func(*Options)) (*GetDataSourceOutput, error) {
	if params == nil {
		params = &GetDataSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataSource", params, optFns, c.addOperationGetDataSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataSourceInput struct {

	// The ID of the Amazon DataZone domain in which the data source exists.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the Amazon DataZone data source.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetDataSourceOutput struct {

	// The ID of the Amazon DataZone domain in which the data source exists.
	//
	// This member is required.
	DomainId *string

	// The ID of the environment where this data source creates and publishes assets,
	//
	// This member is required.
	EnvironmentId *string

	// The ID of the data source.
	//
	// This member is required.
	Id *string

	// The name of the data source.
	//
	// This member is required.
	Name *string

	// The ID of the project where the data source creates and publishes assets.
	//
	// This member is required.
	ProjectId *string

	// The metadata forms attached to the assets created by this data source.
	AssetFormsOutput []types.FormOutput

	// The configuration of the data source.
	Configuration types.DataSourceConfigurationOutput

	// The timestamp of when the data source was created.
	CreatedAt *time.Time

	// The description of the data source.
	Description *string

	// Specifies whether this data source is enabled or not.
	EnableSetting types.EnableSetting

	// Specifies the error message that is returned if the operation cannot be
	// successfully completed.
	ErrorMessage *types.DataSourceErrorMessage

	// The number of assets created by the data source during its last run.
	LastRunAssetCount *int32

	// The timestamp of the last run of the data source.
	LastRunAt *time.Time

	// Specifies the error message that is returned if the operation cannot be
	// successfully completed.
	LastRunErrorMessage *types.DataSourceErrorMessage

	// The status of the last run of the data source.
	LastRunStatus types.DataSourceRunStatus

	// Specifies whether the assets that this data source creates in the inventory are
	// to be also automatically published to the catalog.
	PublishOnImport *bool

	//
	Recommendation *types.RecommendationConfiguration

	// The schedule of the data source runs.
	Schedule *types.ScheduleConfiguration

	// The status of the data source.
	Status types.DataSourceStatus

	// The type of the data source.
	Type *string

	// The timestamp of when the data source was updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataSource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDataSource"); err != nil {
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
	if err = addOpGetDataSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDataSource",
	}
}
