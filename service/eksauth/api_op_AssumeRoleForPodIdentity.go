// Code generated by smithy-go-codegen DO NOT EDIT.

package eksauth

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/eksauth/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The Amazon EKS Auth API and the AssumeRoleForPodIdentity action are only used
// by the EKS Pod Identity Agent. We recommend that applications use the Amazon Web
// Services SDKs to connect to Amazon Web Services services; if credentials from an
// EKS Pod Identity association are available in the pod, the latest versions of
// the SDKs use them automatically.
func (c *Client) AssumeRoleForPodIdentity(ctx context.Context, params *AssumeRoleForPodIdentityInput, optFns ...func(*Options)) (*AssumeRoleForPodIdentityOutput, error) {
	if params == nil {
		params = &AssumeRoleForPodIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssumeRoleForPodIdentity", params, optFns, c.addOperationAssumeRoleForPodIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssumeRoleForPodIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssumeRoleForPodIdentityInput struct {

	// The name of the cluster for the request.
	//
	// This member is required.
	ClusterName *string

	// The token of the Kubernetes service account for the pod.
	//
	// This member is required.
	Token *string

	noSmithyDocumentSerde
}

type AssumeRoleForPodIdentityOutput struct {

	// An object with the permanent IAM role identity and the temporary session name.
	// The ARN of the IAM role that the temporary credentials authenticate to. The
	// session name of the temporary session requested to STS. The value is a unique
	// identifier that contains the role ID, a colon ( : ), and the role session name
	// of the role that is being assumed. The role ID is generated by IAM when the role
	// is created. The role session name part of the value follows this format:
	// eks-clustername-podname-random UUID
	//
	// This member is required.
	AssumedRoleUser *types.AssumedRoleUser

	// The identity that is allowed to use the credentials. This value is always
	// pods.eks.amazonaws.com .
	//
	// This member is required.
	Audience *string

	// The Amazon Web Services Signature Version 4 type of temporary credentials.
	//
	// This member is required.
	Credentials *types.Credentials

	// The Amazon Resource Name (ARN) and ID of the EKS Pod Identity association.
	//
	// This member is required.
	PodIdentityAssociation *types.PodIdentityAssociation

	// The name of the Kubernetes service account inside the cluster to associate the
	// IAM credentials with.
	//
	// This member is required.
	Subject *types.Subject

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssumeRoleForPodIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssumeRoleForPodIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssumeRoleForPodIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssumeRoleForPodIdentity"); err != nil {
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
	if err = addOpAssumeRoleForPodIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssumeRoleForPodIdentity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssumeRoleForPodIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssumeRoleForPodIdentity",
	}
}
