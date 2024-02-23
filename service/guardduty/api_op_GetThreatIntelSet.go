// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the ThreatIntelSet that is specified by the ThreatIntelSet ID.
func (c *Client) GetThreatIntelSet(ctx context.Context, params *GetThreatIntelSetInput, optFns ...func(*Options)) (*GetThreatIntelSetOutput, error) {
	if params == nil {
		params = &GetThreatIntelSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetThreatIntelSet", params, optFns, c.addOperationGetThreatIntelSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetThreatIntelSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetThreatIntelSetInput struct {

	// The unique ID of the detector that the threatIntelSet is associated with.
	//
	// This member is required.
	DetectorId *string

	// The unique ID of the threatIntelSet that you want to get.
	//
	// This member is required.
	ThreatIntelSetId *string

	noSmithyDocumentSerde
}

type GetThreatIntelSetOutput struct {

	// The format of the threatIntelSet.
	//
	// This member is required.
	Format types.ThreatIntelSetFormat

	// The URI of the file that contains the ThreatIntelSet.
	//
	// This member is required.
	Location *string

	// A user-friendly ThreatIntelSet name displayed in all findings that are
	// generated by activity that involves IP addresses included in this
	// ThreatIntelSet.
	//
	// This member is required.
	Name *string

	// The status of threatIntelSet file uploaded.
	//
	// This member is required.
	Status types.ThreatIntelSetStatus

	// The tags of the threat list resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetThreatIntelSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetThreatIntelSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetThreatIntelSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetThreatIntelSet"); err != nil {
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
	if err = addOpGetThreatIntelSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetThreatIntelSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetThreatIntelSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetThreatIntelSet",
	}
}
