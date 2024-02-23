// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// After you submit a UpdateTrafficPolicyInstance request, there's a brief delay
// while Route 53 creates the resource record sets that are specified in the
// traffic policy definition. Use GetTrafficPolicyInstance with the id of updated
// traffic policy instance confirm that the UpdateTrafficPolicyInstance request
// completed successfully. For more information, see the State response element.
// Updates the resource record sets in a specified hosted zone that were created
// based on the settings in a specified traffic policy version. When you update a
// traffic policy instance, Amazon Route 53 continues to respond to DNS queries for
// the root resource record set name (such as example.com) while it replaces one
// group of resource record sets with another. Route 53 performs the following
// operations:
//   - Route 53 creates a new group of resource record sets based on the specified
//     traffic policy. This is true regardless of how significant the differences are
//     between the existing resource record sets and the new resource record sets.
//   - When all of the new resource record sets have been created, Route 53 starts
//     to respond to DNS queries for the root resource record set name (such as
//     example.com) by using the new resource record sets.
//   - Route 53 deletes the old group of resource record sets that are associated
//     with the root resource record set name.
func (c *Client) UpdateTrafficPolicyInstance(ctx context.Context, params *UpdateTrafficPolicyInstanceInput, optFns ...func(*Options)) (*UpdateTrafficPolicyInstanceOutput, error) {
	if params == nil {
		params = &UpdateTrafficPolicyInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTrafficPolicyInstance", params, optFns, c.addOperationUpdateTrafficPolicyInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTrafficPolicyInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the resource record sets that
// you want to update based on a specified traffic policy instance.
type UpdateTrafficPolicyInstanceInput struct {

	// The ID of the traffic policy instance that you want to update.
	//
	// This member is required.
	Id *string

	// The TTL that you want Amazon Route 53 to assign to all of the updated resource
	// record sets.
	//
	// This member is required.
	TTL *int64

	// The ID of the traffic policy that you want Amazon Route 53 to use to update
	// resource record sets for the specified traffic policy instance.
	//
	// This member is required.
	TrafficPolicyId *string

	// The version of the traffic policy that you want Amazon Route 53 to use to
	// update resource record sets for the specified traffic policy instance.
	//
	// This member is required.
	TrafficPolicyVersion *int32

	noSmithyDocumentSerde
}

// A complex type that contains information about the resource record sets that
// Amazon Route 53 created based on a specified traffic policy.
type UpdateTrafficPolicyInstanceOutput struct {

	// A complex type that contains settings for the updated traffic policy instance.
	//
	// This member is required.
	TrafficPolicyInstance *types.TrafficPolicyInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTrafficPolicyInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateTrafficPolicyInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateTrafficPolicyInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateTrafficPolicyInstance"); err != nil {
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
	if err = addOpUpdateTrafficPolicyInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTrafficPolicyInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTrafficPolicyInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateTrafficPolicyInstance",
	}
}
