// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the parameters of a DB parameter group. To modify more than one
// parameter, submit a list of the following: ParameterName , ParameterValue , and
// ApplyMethod . A maximum of 20 parameters can be modified in a single request.
// Changes to dynamic parameters are applied immediately. Changes to static
// parameters require a reboot without failover to the DB instance associated with
// the parameter group before the change can take effect. After you modify a DB
// parameter group, you should wait at least 5 minutes before creating your first
// DB instance that uses that DB parameter group as the default parameter group.
// This allows Amazon Neptune to fully complete the modify action before the
// parameter group is used as the default for a new DB instance. This is especially
// important for parameters that are critical when creating the default database
// for a DB instance, such as the character set for the default database defined by
// the character_set_database parameter. You can use the Parameter Groups option
// of the Amazon Neptune console or the DescribeDBParameters command to verify that
// your DB parameter group has been created or modified.
func (c *Client) ModifyDBParameterGroup(ctx context.Context, params *ModifyDBParameterGroupInput, optFns ...func(*Options)) (*ModifyDBParameterGroupOutput, error) {
	if params == nil {
		params = &ModifyDBParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBParameterGroup", params, optFns, c.addOperationModifyDBParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBParameterGroupInput struct {

	// The name of the DB parameter group. Constraints:
	//   - If supplied, must match the name of an existing DBParameterGroup.
	//
	// This member is required.
	DBParameterGroupName *string

	// An array of parameter names, values, and the apply method for the parameter
	// update. At least one parameter name, value, and apply method must be supplied;
	// subsequent arguments are optional. A maximum of 20 parameters can be modified in
	// a single request. Valid Values (for the application method): immediate |
	// pending-reboot You can use the immediate value with dynamic parameters only. You
	// can use the pending-reboot value for both dynamic and static parameters, and
	// changes are applied when you reboot the DB instance without failover.
	//
	// This member is required.
	Parameters []types.Parameter

	noSmithyDocumentSerde
}

type ModifyDBParameterGroupOutput struct {

	// Provides the name of the DB parameter group.
	DBParameterGroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyDBParameterGroup"); err != nil {
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
	if err = addOpModifyDBParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyDBParameterGroup",
	}
}
