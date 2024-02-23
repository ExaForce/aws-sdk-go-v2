// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcases

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/connectcases/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// If you provide a value for PerformedBy.UserArn you must also have
// connect:DescribeUser (https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeUser.html)
// permission on the User ARN resource that you provide Updates the values of
// fields on a case. Fields to be updated are received as an array of id/value
// pairs identical to the CreateCase input . If the action is successful, the
// service sends back an HTTP 200 response with an empty HTTP body.
func (c *Client) UpdateCase(ctx context.Context, params *UpdateCaseInput, optFns ...func(*Options)) (*UpdateCaseOutput, error) {
	if params == nil {
		params = &UpdateCaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCase", params, optFns, c.addOperationUpdateCaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCaseInput struct {

	// A unique identifier of the case.
	//
	// This member is required.
	CaseId *string

	// The unique identifier of the Cases domain.
	//
	// This member is required.
	DomainId *string

	// An array of objects with fieldId (matching ListFields/DescribeField) and value
	// union data, structured identical to CreateCase .
	//
	// This member is required.
	Fields []types.FieldValue

	// Represents the identity of the person who performed the action.
	PerformedBy types.UserUnion

	noSmithyDocumentSerde
}

type UpdateCaseOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateCaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateCase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateCase{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateCase"); err != nil {
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
	if err = addOpUpdateCaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCase(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateCase",
	}
}
