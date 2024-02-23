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

// Returns all of the runtime parameters offered by the underlying database
// software, or engine, for a specific database in Amazon Lightsail. In addition to
// the parameter names and values, this operation returns other information about
// each parameter. This information includes whether changes require a reboot,
// whether the parameter is modifiable, the allowed values, and the data types.
func (c *Client) GetRelationalDatabaseParameters(ctx context.Context, params *GetRelationalDatabaseParametersInput, optFns ...func(*Options)) (*GetRelationalDatabaseParametersOutput, error) {
	if params == nil {
		params = &GetRelationalDatabaseParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRelationalDatabaseParameters", params, optFns, c.addOperationGetRelationalDatabaseParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRelationalDatabaseParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabaseParametersInput struct {

	// The name of your database for which to get parameters.
	//
	// This member is required.
	RelationalDatabaseName *string

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetRelationalDatabaseParameters request. If your
	// results are paginated, the response will return a next page token that you can
	// specify as the page token in a subsequent request.
	PageToken *string

	noSmithyDocumentSerde
}

type GetRelationalDatabaseParametersOutput struct {

	// The token to advance to the next page of results from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetRelationalDatabaseParameters request and
	// specify the next page token using the pageToken parameter.
	NextPageToken *string

	// An object describing the result of your get relational database parameters
	// request.
	Parameters []types.RelationalDatabaseParameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRelationalDatabaseParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabaseParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabaseParameters{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRelationalDatabaseParameters"); err != nil {
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
	if err = addOpGetRelationalDatabaseParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabaseParameters(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRelationalDatabaseParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRelationalDatabaseParameters",
	}
}
