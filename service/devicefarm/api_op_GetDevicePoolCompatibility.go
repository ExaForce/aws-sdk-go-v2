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

// Gets information about compatibility with a device pool.
func (c *Client) GetDevicePoolCompatibility(ctx context.Context, params *GetDevicePoolCompatibilityInput, optFns ...func(*Options)) (*GetDevicePoolCompatibilityOutput, error) {
	if params == nil {
		params = &GetDevicePoolCompatibilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDevicePoolCompatibility", params, optFns, c.addOperationGetDevicePoolCompatibilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDevicePoolCompatibilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to the get device pool compatibility operation.
type GetDevicePoolCompatibilityInput struct {

	// The device pool's ARN.
	//
	// This member is required.
	DevicePoolArn *string

	// The ARN of the app that is associated with the specified device pool.
	AppArn *string

	// An object that contains information about the settings for a run.
	Configuration *types.ScheduleRunConfiguration

	// Information about the uploaded test to be run against the device pool.
	Test *types.ScheduleRunTest

	// The test type for the specified device pool. Allowed values include the
	// following:
	//   - BUILTIN_FUZZ.
	//   - BUILTIN_EXPLORER. For Android, an app explorer that traverses an Android
	//   app, interacting with it and capturing screenshots at the same time.
	//   - APPIUM_JAVA_JUNIT.
	//   - APPIUM_JAVA_TESTNG.
	//   - APPIUM_PYTHON.
	//   - APPIUM_NODE.
	//   - APPIUM_RUBY.
	//   - APPIUM_WEB_JAVA_JUNIT.
	//   - APPIUM_WEB_JAVA_TESTNG.
	//   - APPIUM_WEB_PYTHON.
	//   - APPIUM_WEB_NODE.
	//   - APPIUM_WEB_RUBY.
	//   - CALABASH.
	//   - INSTRUMENTATION.
	//   - UIAUTOMATION.
	//   - UIAUTOMATOR.
	//   - XCTEST.
	//   - XCTEST_UI.
	TestType types.TestType

	noSmithyDocumentSerde
}

// Represents the result of describe device pool compatibility request.
type GetDevicePoolCompatibilityOutput struct {

	// Information about compatible devices.
	CompatibleDevices []types.DevicePoolCompatibilityResult

	// Information about incompatible devices.
	IncompatibleDevices []types.DevicePoolCompatibilityResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDevicePoolCompatibilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDevicePoolCompatibility{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDevicePoolCompatibility{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDevicePoolCompatibility"); err != nil {
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
	if err = addOpGetDevicePoolCompatibilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDevicePoolCompatibility(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDevicePoolCompatibility(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDevicePoolCompatibility",
	}
}
