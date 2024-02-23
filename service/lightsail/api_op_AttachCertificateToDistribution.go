// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches an SSL/TLS certificate to your Amazon Lightsail content delivery
// network (CDN) distribution. After the certificate is attached, your distribution
// accepts HTTPS traffic for all of the domains that are associated with the
// certificate. Use the CreateCertificate action to create a certificate that you
// can attach to your distribution. Only certificates created in the us-east-1
// Amazon Web Services Region can be attached to Lightsail distributions. Lightsail
// distributions are global resources that can reference an origin in any Amazon
// Web Services Region, and distribute its content globally. However, all
// distributions are located in the us-east-1 Region.
func (c *Client) AttachCertificateToDistribution(ctx context.Context, params *AttachCertificateToDistributionInput, optFns ...func(*Options)) (*AttachCertificateToDistributionOutput, error) {
	if params == nil {
		params = &AttachCertificateToDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachCertificateToDistribution", params, optFns, c.addOperationAttachCertificateToDistributionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachCertificateToDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachCertificateToDistributionInput struct {

	// The name of the certificate to attach to a distribution. Only certificates with
	// a status of ISSUED can be attached to a distribution. Use the GetCertificates
	// action to get a list of certificate names that you can specify. This is the name
	// of the certificate resource type and is used only to reference the certificate
	// in other API actions. It can be different than the domain name of the
	// certificate. For example, your certificate name might be
	// WordPress-Blog-Certificate and the domain name of the certificate might be
	// example.com .
	//
	// This member is required.
	CertificateName *string

	// The name of the distribution that the certificate will be attached to. Use the
	// GetDistributions action to get a list of distribution names that you can specify.
	//
	// This member is required.
	DistributionName *string

	noSmithyDocumentSerde
}

type AttachCertificateToDistributionOutput struct {

	// An object that describes the result of the action, such as the status of the
	// request, the timestamp of the request, and the resources affected by the
	// request.
	Operation *types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachCertificateToDistributionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAttachCertificateToDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachCertificateToDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AttachCertificateToDistribution"); err != nil {
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
	if err = addOpAttachCertificateToDistributionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachCertificateToDistribution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachCertificateToDistribution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AttachCertificateToDistribution",
	}
}
