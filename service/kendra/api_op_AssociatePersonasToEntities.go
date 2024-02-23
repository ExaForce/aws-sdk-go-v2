// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Defines the specific permissions of users or groups in your IAM Identity Center
// identity source with access to your Amazon Kendra experience. You can create an
// Amazon Kendra experience such as a search application. For more information on
// creating a search application experience, see Building a search experience with
// no code (https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html)
// .
func (c *Client) AssociatePersonasToEntities(ctx context.Context, params *AssociatePersonasToEntitiesInput, optFns ...func(*Options)) (*AssociatePersonasToEntitiesOutput, error) {
	if params == nil {
		params = &AssociatePersonasToEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociatePersonasToEntities", params, optFns, c.addOperationAssociatePersonasToEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociatePersonasToEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociatePersonasToEntitiesInput struct {

	// The identifier of your Amazon Kendra experience.
	//
	// This member is required.
	Id *string

	// The identifier of the index for your Amazon Kendra experience.
	//
	// This member is required.
	IndexId *string

	// The personas that define the specific permissions of users or groups in your
	// IAM Identity Center identity source. The available personas or access roles are
	// Owner and Viewer . For more information on these personas, see Providing access
	// to your search page (https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html#access-search-experience)
	// .
	//
	// This member is required.
	Personas []types.EntityPersonaConfiguration

	noSmithyDocumentSerde
}

type AssociatePersonasToEntitiesOutput struct {

	// Lists the users or groups in your IAM Identity Center identity source that
	// failed to properly configure with your Amazon Kendra experience.
	FailedEntityList []types.FailedEntity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociatePersonasToEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociatePersonasToEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociatePersonasToEntities{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociatePersonasToEntities"); err != nil {
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
	if err = addOpAssociatePersonasToEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociatePersonasToEntities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociatePersonasToEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociatePersonasToEntities",
	}
}
