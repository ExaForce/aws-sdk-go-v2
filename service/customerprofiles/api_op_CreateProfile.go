// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a standard profile. A standard profile represents the following
// attributes for a customer profile in a domain.
func (c *Client) CreateProfile(ctx context.Context, params *CreateProfileInput, optFns ...func(*Options)) (*CreateProfileOutput, error) {
	if params == nil {
		params = &CreateProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProfile", params, optFns, c.addOperationCreateProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProfileInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// A unique account number that you have given to the customer.
	AccountNumber *string

	// Any additional information relevant to the customer’s profile.
	AdditionalInformation *string

	// A generic address associated with the customer that is not mailing, shipping,
	// or billing.
	Address *types.Address

	// A key value pair of attributes of a customer profile.
	Attributes map[string]string

	// The customer’s billing address.
	BillingAddress *types.Address

	// The customer’s birth date.
	BirthDate *string

	// The customer’s business email address.
	BusinessEmailAddress *string

	// The name of the customer’s business.
	BusinessName *string

	// The customer’s business phone number.
	BusinessPhoneNumber *string

	// The customer’s email address, which has not been specified as a personal or
	// business address.
	EmailAddress *string

	// The customer’s first name.
	FirstName *string

	// The gender with which the customer identifies.
	//
	// Deprecated: This member has been deprecated.
	Gender types.Gender

	// An alternative to Gender which accepts any string as input.
	GenderString *string

	// The customer’s home phone number.
	HomePhoneNumber *string

	// The customer’s last name.
	LastName *string

	// The customer’s mailing address.
	MailingAddress *types.Address

	// The customer’s middle name.
	MiddleName *string

	// The customer’s mobile phone number.
	MobilePhoneNumber *string

	// The type of profile used to describe the customer.
	//
	// Deprecated: This member has been deprecated.
	PartyType types.PartyType

	// An alternative to PartyType which accepts any string as input.
	PartyTypeString *string

	// The customer’s personal email address.
	PersonalEmailAddress *string

	// The customer’s phone number, which has not been specified as a mobile, home, or
	// business number.
	PhoneNumber *string

	// The customer’s shipping address.
	ShippingAddress *types.Address

	noSmithyDocumentSerde
}

type CreateProfileOutput struct {

	// The unique identifier of a customer profile.
	//
	// This member is required.
	ProfileId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateProfile"); err != nil {
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
	if err = addOpCreateProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateProfile",
	}
}
