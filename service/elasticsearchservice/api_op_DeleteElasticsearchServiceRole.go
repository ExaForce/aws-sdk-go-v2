// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the service-linked role that Elasticsearch Service uses to manage and
// maintain VPC domains. Role deletion will fail if any existing VPC domains use
// the role. You must delete any such Elasticsearch domains before deleting the
// role. See Deleting Elasticsearch Service Role (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-enabling-slr)
// in VPC Endpoints for Amazon Elasticsearch Service Domains.
func (c *Client) DeleteElasticsearchServiceRole(ctx context.Context, params *DeleteElasticsearchServiceRoleInput, optFns ...func(*Options)) (*DeleteElasticsearchServiceRoleOutput, error) {
	if params == nil {
		params = &DeleteElasticsearchServiceRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteElasticsearchServiceRole", params, optFns, c.addOperationDeleteElasticsearchServiceRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteElasticsearchServiceRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteElasticsearchServiceRoleInput struct {
	noSmithyDocumentSerde
}

type DeleteElasticsearchServiceRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteElasticsearchServiceRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteElasticsearchServiceRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteElasticsearchServiceRole{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteElasticsearchServiceRole"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteElasticsearchServiceRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteElasticsearchServiceRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteElasticsearchServiceRole",
	}
}
