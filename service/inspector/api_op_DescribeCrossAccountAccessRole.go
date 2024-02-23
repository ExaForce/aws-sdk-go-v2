// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the IAM role that enables Amazon Inspector to access your AWS account.
func (c *Client) DescribeCrossAccountAccessRole(ctx context.Context, params *DescribeCrossAccountAccessRoleInput, optFns ...func(*Options)) (*DescribeCrossAccountAccessRoleOutput, error) {
	if params == nil {
		params = &DescribeCrossAccountAccessRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCrossAccountAccessRole", params, optFns, c.addOperationDescribeCrossAccountAccessRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCrossAccountAccessRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCrossAccountAccessRoleInput struct {
	noSmithyDocumentSerde
}

type DescribeCrossAccountAccessRoleOutput struct {

	// The date when the cross-account access role was registered.
	//
	// This member is required.
	RegisteredAt *time.Time

	// The ARN that specifies the IAM role that Amazon Inspector uses to access your
	// AWS account.
	//
	// This member is required.
	RoleArn *string

	// A Boolean value that specifies whether the IAM role has the necessary policies
	// attached to enable Amazon Inspector to access your AWS account.
	//
	// This member is required.
	Valid *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCrossAccountAccessRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCrossAccountAccessRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCrossAccountAccessRole{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeCrossAccountAccessRole"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCrossAccountAccessRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCrossAccountAccessRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeCrossAccountAccessRole",
	}
}
