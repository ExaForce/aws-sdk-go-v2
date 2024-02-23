// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the settings for the sensitivity inspection template for an account.
func (c *Client) GetSensitivityInspectionTemplate(ctx context.Context, params *GetSensitivityInspectionTemplateInput, optFns ...func(*Options)) (*GetSensitivityInspectionTemplateOutput, error) {
	if params == nil {
		params = &GetSensitivityInspectionTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSensitivityInspectionTemplate", params, optFns, c.addOperationGetSensitivityInspectionTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSensitivityInspectionTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSensitivityInspectionTemplateInput struct {

	// The unique identifier for the Amazon Macie resource that the request applies to.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetSensitivityInspectionTemplateOutput struct {

	// The custom description of the template.
	Description *string

	// The managed data identifiers that are explicitly excluded (not used) when
	// analyzing data.
	Excludes *types.SensitivityInspectionTemplateExcludes

	// The allow lists, custom data identifiers, and managed data identifiers that are
	// explicitly included (used) when analyzing data.
	Includes *types.SensitivityInspectionTemplateIncludes

	// The name of the template: automated-sensitive-data-discovery.
	Name *string

	// The unique identifier for the template.
	SensitivityInspectionTemplateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSensitivityInspectionTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSensitivityInspectionTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSensitivityInspectionTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSensitivityInspectionTemplate"); err != nil {
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
	if err = addOpGetSensitivityInspectionTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSensitivityInspectionTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSensitivityInspectionTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSensitivityInspectionTemplate",
	}
}
