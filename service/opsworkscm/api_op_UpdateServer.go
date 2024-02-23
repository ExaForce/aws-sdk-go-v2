// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates settings for a server. This operation is synchronous.
func (c *Client) UpdateServer(ctx context.Context, params *UpdateServerInput, optFns ...func(*Options)) (*UpdateServerOutput, error) {
	if params == nil {
		params = &UpdateServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateServer", params, optFns, c.addOperationUpdateServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServerInput struct {

	// The name of the server to update.
	//
	// This member is required.
	ServerName *string

	// Sets the number of automated backups that you want to keep.
	BackupRetentionCount *int32

	// Setting DisableAutomatedBackup to true disables automated or scheduled backups.
	// Automated backups are enabled by default.
	DisableAutomatedBackup *bool

	// DDD:HH:MM (weekly start time) or HH:MM (daily start time). Time windows always
	// use coordinated universal time (UTC). Valid strings for day of week ( DDD ) are:
	// Mon , Tue , Wed , Thr , Fri , Sat , or Sun .
	PreferredBackupWindow *string

	// DDD:HH:MM (weekly start time) or HH:MM (daily start time). Time windows always
	// use coordinated universal time (UTC). Valid strings for day of week ( DDD ) are:
	// Mon , Tue , Wed , Thr , Fri , Sat , or Sun .
	PreferredMaintenanceWindow *string

	noSmithyDocumentSerde
}

type UpdateServerOutput struct {

	// Contains the response to a UpdateServer request.
	Server *types.Server

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateServer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateServer"); err != nil {
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
	if err = addOpUpdateServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateServer",
	}
}
