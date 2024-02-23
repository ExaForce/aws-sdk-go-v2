// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates information about an existing private device instance profile.
func (c *Client) UpdateInstanceProfile(ctx context.Context, params *UpdateInstanceProfileInput, optFns ...func(*Options)) (*UpdateInstanceProfileOutput, error) {
	if params == nil {
		params = &UpdateInstanceProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInstanceProfile", params, optFns, c.addOperationUpdateInstanceProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInstanceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateInstanceProfileInput struct {

	// The Amazon Resource Name (ARN) of the instance profile.
	//
	// This member is required.
	Arn *string

	// The updated description for your instance profile.
	Description *string

	// An array of strings that specifies the list of app packages that should not be
	// cleaned up from the device after a test run is over. The list of packages is
	// only considered if you set packageCleanup to true .
	ExcludeAppPackagesFromCleanup []string

	// The updated name for your instance profile.
	Name *string

	// The updated choice for whether you want to specify package cleanup. The default
	// value is false for private devices.
	PackageCleanup *bool

	// The updated choice for whether you want to reboot the device after use. The
	// default value is true .
	RebootAfterUse *bool

	noSmithyDocumentSerde
}

type UpdateInstanceProfileOutput struct {

	// An object that contains information about your instance profile.
	InstanceProfile *types.InstanceProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateInstanceProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateInstanceProfile"); err != nil {
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
	if err = addOpUpdateInstanceProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInstanceProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateInstanceProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateInstanceProfile",
	}
}
