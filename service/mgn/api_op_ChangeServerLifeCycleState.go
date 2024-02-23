// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows the user to set the SourceServer.LifeCycle.state property for specific
// Source Server IDs to one of the following: READY_FOR_TEST or READY_FOR_CUTOVER.
// This command only works if the Source Server is already launchable
// (dataReplicationInfo.lagDuration is not null.)
func (c *Client) ChangeServerLifeCycleState(ctx context.Context, params *ChangeServerLifeCycleStateInput, optFns ...func(*Options)) (*ChangeServerLifeCycleStateOutput, error) {
	if params == nil {
		params = &ChangeServerLifeCycleStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ChangeServerLifeCycleState", params, optFns, c.addOperationChangeServerLifeCycleStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ChangeServerLifeCycleStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ChangeServerLifeCycleStateInput struct {

	// The request to change the source server migration lifecycle state.
	//
	// This member is required.
	LifeCycle *types.ChangeServerLifeCycleStateSourceServerLifecycle

	// The request to change the source server migration lifecycle state by source
	// server ID.
	//
	// This member is required.
	SourceServerID *string

	// The request to change the source server migration account ID.
	AccountID *string

	noSmithyDocumentSerde
}

type ChangeServerLifeCycleStateOutput struct {

	// Source server application ID.
	ApplicationID *string

	// Source server ARN.
	Arn *string

	// Source Server connector action.
	ConnectorAction *types.SourceServerConnectorAction

	// Source server data replication info.
	DataReplicationInfo *types.DataReplicationInfo

	// Source server fqdn for action framework.
	FqdnForActionFramework *string

	// Source server archived status.
	IsArchived *bool

	// Source server launched instance.
	LaunchedInstance *types.LaunchedInstance

	// Source server lifecycle state.
	LifeCycle *types.LifeCycle

	// Source server replication type.
	ReplicationType types.ReplicationType

	// Source server properties.
	SourceProperties *types.SourceProperties

	// Source server ID.
	SourceServerID *string

	// Source server Tags.
	Tags map[string]string

	// Source server user provided ID.
	UserProvidedID *string

	// Source server vCenter client id.
	VcenterClientID *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationChangeServerLifeCycleStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpChangeServerLifeCycleState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpChangeServerLifeCycleState{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ChangeServerLifeCycleState"); err != nil {
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
	if err = addOpChangeServerLifeCycleStateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opChangeServerLifeCycleState(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opChangeServerLifeCycleState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ChangeServerLifeCycleState",
	}
}
