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

// Gets an Amazon DataZone data source run.
func (c *Client) GetDataSourceRun(ctx context.Context, params *GetDataSourceRunInput, optFns ...func(*Options)) (*GetDataSourceRunOutput, error) {
	if params == nil {
		params = &GetDataSourceRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataSourceRun", params, optFns, c.addOperationGetDataSourceRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataSourceRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataSourceRunInput struct {

	// The ID of the domain in which this data source run was performed.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the data source run.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetDataSourceRunOutput struct {

	// The timestamp of when the data source run was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The ID of the data source for this data source run.
	//
	// This member is required.
	DataSourceId *string

	// The ID of the domain in which this data source run was performed.
	//
	// This member is required.
	DomainId *string

	// The ID of the data source run.
	//
	// This member is required.
	Id *string

	// The ID of the project in which this data source run occured.
	//
	// This member is required.
	ProjectId *string

	// The status of this data source run.
	//
	// This member is required.
	Status types.DataSourceRunStatus

	// The type of this data source run.
	//
	// This member is required.
	Type types.DataSourceRunType

	// The timestamp of when this data source run was updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// The configuration snapshot of the data source run.
	DataSourceConfigurationSnapshot *string

	// Specifies the error message that is returned if the operation cannot be
	// successfully completed.
	ErrorMessage *types.DataSourceErrorMessage

	// The asset statistics from this data source run.
	RunStatisticsForAssets *types.RunStatisticsForAssets

	// The timestamp of when this data source run started.
	StartedAt *time.Time

	// The timestamp of when this data source run stopped.
	StoppedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataSourceRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataSourceRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataSourceRun{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDataSourceRun"); err != nil {
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
	if err = addOpGetDataSourceRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataSourceRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataSourceRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDataSourceRun",
	}
}
