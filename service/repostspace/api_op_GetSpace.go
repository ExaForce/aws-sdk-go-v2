// Code generated by smithy-go-codegen DO NOT EDIT.

package repostspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/repostspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Displays information about the AWS re:Post Private private re:Post.
func (c *Client) GetSpace(ctx context.Context, params *GetSpaceInput, optFns ...func(*Options)) (*GetSpaceOutput, error) {
	if params == nil {
		params = &GetSpaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSpace", params, optFns, c.addOperationGetSpaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSpaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSpaceInput struct {

	// The ID of the private re:Post.
	//
	// This member is required.
	SpaceId *string

	noSmithyDocumentSerde
}

type GetSpaceOutput struct {

	// The ARN of the private re:Post.
	//
	// This member is required.
	Arn *string

	// The Identity Center identifier for the Application Instance.
	//
	// This member is required.
	ClientId *string

	// The configuration status of the private re:Post.
	//
	// This member is required.
	ConfigurationStatus types.ConfigurationStatus

	// The date when the private re:Post was created.
	//
	// This member is required.
	CreateDateTime *time.Time

	// The name of the private re:Post.
	//
	// This member is required.
	Name *string

	// The AWS generated subdomain of the private re:Post
	//
	// This member is required.
	RandomDomain *string

	// The unique ID of the private re:Post.
	//
	// This member is required.
	SpaceId *string

	// The creation or deletion status of the private re:Post.
	//
	// This member is required.
	Status *string

	// The storage limit of the private re:Post.
	//
	// This member is required.
	StorageLimit *int64

	// The pricing tier of the private re:Post.
	//
	// This member is required.
	Tier types.TierLevel

	// The custom subdomain that you use to access your private re:Post. All custom
	// subdomains must be approved by AWS before use.
	//
	// This member is required.
	VanityDomain *string

	// The approval status of the custom subdomain.
	//
	// This member is required.
	VanityDomainStatus types.VanityDomainStatus

	// The content size of the private re:Post.
	ContentSize *int64

	// The IAM role that grants permissions to the private re:Post to convert
	// unanswered questions into AWS support tickets.
	CustomerRoleArn *string

	// The date when the private re:Post was deleted.
	DeleteDateTime *time.Time

	// The description of the private re:Post.
	Description *string

	// The list of groups that are administrators of the private re:Post.
	GroupAdmins []string

	// The list of users that are administrators of the private re:Post.
	UserAdmins []string

	// The number of users that have onboarded to the private re:Post.
	UserCount *int32

	// The custom AWS KMS key ARN that’s used for the AWS KMS encryption.
	UserKMSKey *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSpaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSpace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSpace{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSpace"); err != nil {
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
	if err = addOpGetSpaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSpace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSpace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSpace",
	}
}
