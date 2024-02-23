// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Begins the export of a discovered data report to an Amazon S3 bucket managed by
// Amazon Web Services. Exports might provide an estimate of fees and savings based
// on certain information that you provide. Fee estimates do not include any taxes
// that might apply. Your actual fees and savings depend on a variety of factors,
// including your actual usage of Amazon Web Services services, which might vary
// from the estimates provided in this report. If you do not specify preferences
// or agentIds in the filter, a summary of all servers, applications, tags, and
// performance is generated. This data is an aggregation of all server data
// collected through on-premises tooling, file import, application grouping and
// applying tags. If you specify agentIds in a filter, the task exports up to 72
// hours of detailed data collected by the identified Application Discovery Agent,
// including network, process, and performance details. A time range for exported
// agent data may be set by using startTime and endTime . Export of detailed agent
// data is limited to five concurrently running exports. Export of detailed agent
// data is limited to two exports per day. If you enable
// ec2RecommendationsPreferences in preferences , an Amazon EC2 instance matching
// the characteristics of each server in Application Discovery Service is
// generated. Changing the attributes of the ec2RecommendationsPreferences changes
// the criteria of the recommendation.
func (c *Client) StartExportTask(ctx context.Context, params *StartExportTaskInput, optFns ...func(*Options)) (*StartExportTaskOutput, error) {
	if params == nil {
		params = &StartExportTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartExportTask", params, optFns, c.addOperationStartExportTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartExportTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartExportTaskInput struct {

	// The end timestamp for exported data from the single Application Discovery Agent
	// selected in the filters. If no value is specified, exported data includes the
	// most recent data collected by the agent.
	EndTime *time.Time

	// The file format for the returned export data. Default value is CSV . Note: The
	// GRAPHML option has been deprecated.
	ExportDataFormat []types.ExportDataFormat

	// If a filter is present, it selects the single agentId of the Application
	// Discovery Agent for which data is exported. The agentId can be found in the
	// results of the DescribeAgents API or CLI. If no filter is present, startTime
	// and endTime are ignored and exported data includes both Amazon Web Services
	// Application Discovery Service Agentless Collector collectors data and summary
	// data from Application Discovery Agent agents.
	Filters []types.ExportFilter

	// Indicates the type of data that needs to be exported. Only one ExportPreferences (https://docs.aws.amazon.com/application-discovery/latest/APIReference/API_ExportPreferences.html)
	// can be enabled at any time.
	Preferences types.ExportPreferences

	// The start timestamp for exported data from the single Application Discovery
	// Agent selected in the filters. If no value is specified, data is exported
	// starting from the first data collected by the agent.
	StartTime *time.Time

	noSmithyDocumentSerde
}

type StartExportTaskOutput struct {

	// A unique identifier used to query the status of an export request.
	ExportId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartExportTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartExportTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartExportTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartExportTask"); err != nil {
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
	if err = addOpStartExportTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartExportTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartExportTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartExportTask",
	}
}
