// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
	"time"
)

// Updates the resource policy of the S3 Access Grants instance. Permissions You
// must have the s3:PutAccessGrantsInstanceResourcePolicy permission to use this
// operation.
func (c *Client) PutAccessGrantsInstanceResourcePolicy(ctx context.Context, params *PutAccessGrantsInstanceResourcePolicyInput, optFns ...func(*Options)) (*PutAccessGrantsInstanceResourcePolicyOutput, error) {
	if params == nil {
		params = &PutAccessGrantsInstanceResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAccessGrantsInstanceResourcePolicy", params, optFns, c.addOperationPutAccessGrantsInstanceResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAccessGrantsInstanceResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAccessGrantsInstanceResourcePolicyInput struct {

	// The ID of the Amazon Web Services account that is making this request.
	//
	// This member is required.
	AccountId *string

	// The resource policy of the S3 Access Grants instance that you are updating.
	//
	// This member is required.
	Policy *string

	// The Organization of the resource policy of the S3 Access Grants instance.
	Organization *string

	noSmithyDocumentSerde
}

func (in *PutAccessGrantsInstanceResourcePolicyInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.RequiresAccountId = ptr.Bool(true)
}

type PutAccessGrantsInstanceResourcePolicyOutput struct {

	// The date and time when you created the S3 Access Grants instance resource
	// policy.
	CreatedAt *time.Time

	// The Organization of the resource policy of the S3 Access Grants instance.
	Organization *string

	// The updated resource policy of the S3 Access Grants instance.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAccessGrantsInstanceResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutAccessGrantsInstanceResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutAccessGrantsInstanceResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutAccessGrantsInstanceResourcePolicy"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opPutAccessGrantsInstanceResourcePolicyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutAccessGrantsInstanceResourcePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccessGrantsInstanceResourcePolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutAccessGrantsInstanceResourcePolicyUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opPutAccessGrantsInstanceResourcePolicyMiddleware struct {
}

func (*endpointPrefix_opPutAccessGrantsInstanceResourcePolicyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutAccessGrantsInstanceResourcePolicyMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*PutAccessGrantsInstanceResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opPutAccessGrantsInstanceResourcePolicyMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opPutAccessGrantsInstanceResourcePolicyMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opPutAccessGrantsInstanceResourcePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutAccessGrantsInstanceResourcePolicy",
	}
}

func copyPutAccessGrantsInstanceResourcePolicyInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*PutAccessGrantsInstanceResourcePolicyInput)
	if !ok {
		return nil, fmt.Errorf("expect *PutAccessGrantsInstanceResourcePolicyInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *PutAccessGrantsInstanceResourcePolicyInput) copy() interface{} {
	v := *in
	return &v
}
func backFillPutAccessGrantsInstanceResourcePolicyAccountID(input interface{}, v string) error {
	in := input.(*PutAccessGrantsInstanceResourcePolicyInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addPutAccessGrantsInstanceResourcePolicyUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyPutAccessGrantsInstanceResourcePolicyInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
