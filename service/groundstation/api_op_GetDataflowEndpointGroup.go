// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/groundstation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the dataflow endpoint group.
func (c *Client) GetDataflowEndpointGroup(ctx context.Context, params *GetDataflowEndpointGroupInput, optFns ...func(*Options)) (*GetDataflowEndpointGroupOutput, error) {
	if params == nil {
		params = &GetDataflowEndpointGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataflowEndpointGroup", params, optFns, c.addOperationGetDataflowEndpointGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataflowEndpointGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataflowEndpointGroupInput struct {

	// UUID of a dataflow endpoint group.
	//
	// This member is required.
	DataflowEndpointGroupId *string

	noSmithyDocumentSerde
}

type GetDataflowEndpointGroupOutput struct {

	// Amount of time, in seconds, after a contact ends that the Ground Station
	// Dataflow Endpoint Group will be in a POSTPASS state. A Ground Station Dataflow
	// Endpoint Group State Change event will be emitted when the Dataflow Endpoint
	// Group enters and exits the POSTPASS state.
	ContactPostPassDurationSeconds *int32

	// Amount of time, in seconds, before a contact starts that the Ground Station
	// Dataflow Endpoint Group will be in a PREPASS state. A Ground Station Dataflow
	// Endpoint Group State Change event will be emitted when the Dataflow Endpoint
	// Group enters and exits the PREPASS state.
	ContactPrePassDurationSeconds *int32

	// ARN of a dataflow endpoint group.
	DataflowEndpointGroupArn *string

	// UUID of a dataflow endpoint group.
	DataflowEndpointGroupId *string

	// Details of a dataflow endpoint.
	EndpointsDetails []types.EndpointDetails

	// Tags assigned to a dataflow endpoint group.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataflowEndpointGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataflowEndpointGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataflowEndpointGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDataflowEndpointGroup"); err != nil {
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
	if err = addOpGetDataflowEndpointGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataflowEndpointGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataflowEndpointGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDataflowEndpointGroup",
	}
}
