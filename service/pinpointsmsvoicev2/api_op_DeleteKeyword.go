// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an existing keyword from an origination phone number or pool. A keyword
// is a word that you can search for on a particular phone number or pool. It is
// also a specific word or phrase that an end user can send to your number to
// elicit a response, such as an informational message or a special offer. When
// your number receives a message that begins with a keyword, Amazon Pinpoint
// responds with a customizable message. Keywords "HELP" and "STOP" can't be
// deleted or modified.
func (c *Client) DeleteKeyword(ctx context.Context, params *DeleteKeywordInput, optFns ...func(*Options)) (*DeleteKeywordOutput, error) {
	if params == nil {
		params = &DeleteKeywordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteKeyword", params, optFns, c.addOperationDeleteKeywordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteKeywordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteKeywordInput struct {

	// The keyword to delete.
	//
	// This member is required.
	Keyword *string

	// The origination identity to use such as a PhoneNumberId, PhoneNumberArn, PoolId
	// or PoolArn. You can use DescribePhoneNumbers to find the values for
	// PhoneNumberId and PhoneNumberArn and DescribePools to find the values of PoolId
	// and PoolArn.
	//
	// This member is required.
	OriginationIdentity *string

	noSmithyDocumentSerde
}

type DeleteKeywordOutput struct {

	// The keyword that was deleted.
	Keyword *string

	// The action that was associated with the deleted keyword.
	KeywordAction types.KeywordAction

	// The message that was associated with the deleted keyword.
	KeywordMessage *string

	// The PhoneNumberId or PoolId that the keyword was associated with.
	OriginationIdentity *string

	// The PhoneNumberArn or PoolArn that the keyword was associated with.
	OriginationIdentityArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteKeywordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteKeyword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteKeyword{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteKeyword"); err != nil {
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
	if err = addOpDeleteKeywordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteKeyword(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteKeyword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteKeyword",
	}
}
