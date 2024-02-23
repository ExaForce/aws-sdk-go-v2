// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a signing profile. A signing profile is a code-signing template that
// can be used to carry out a pre-defined signing job.
func (c *Client) PutSigningProfile(ctx context.Context, params *PutSigningProfileInput, optFns ...func(*Options)) (*PutSigningProfileOutput, error) {
	if params == nil {
		params = &PutSigningProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutSigningProfile", params, optFns, c.addOperationPutSigningProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutSigningProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutSigningProfileInput struct {

	// The ID of the signing platform to be created.
	//
	// This member is required.
	PlatformId *string

	// The name of the signing profile to be created.
	//
	// This member is required.
	ProfileName *string

	// A subfield of platform . This specifies any different configuration options that
	// you want to apply to the chosen platform (such as a different hash-algorithm or
	// signing-algorithm ).
	Overrides *types.SigningPlatformOverrides

	// The default validity period override for any signature generated using this
	// signing profile. If unspecified, the default is 135 months.
	SignatureValidityPeriod *types.SignatureValidityPeriod

	// The AWS Certificate Manager certificate that will be used to sign code with the
	// new signing profile.
	SigningMaterial *types.SigningMaterial

	// Map of key-value pairs for signing. These can include any information that you
	// want to use during signing.
	SigningParameters map[string]string

	// Tags to be associated with the signing profile that is being created.
	Tags map[string]string

	noSmithyDocumentSerde
}

type PutSigningProfileOutput struct {

	// The Amazon Resource Name (ARN) of the signing profile created.
	Arn *string

	// The version of the signing profile being created.
	ProfileVersion *string

	// The signing profile ARN, including the profile version.
	ProfileVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutSigningProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutSigningProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutSigningProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutSigningProfile"); err != nil {
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
	if err = addOpPutSigningProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutSigningProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutSigningProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutSigningProfile",
	}
}
