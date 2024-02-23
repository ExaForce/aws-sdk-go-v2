// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an intent. To define the interaction between the user and your bot, you
// define one or more intents. For example, for a pizza ordering bot you would
// create an OrderPizza intent. When you create an intent, you must provide a
// name. You can optionally provide the following:
//   - Sample utterances. For example, "I want to order a pizza" and "Can I order
//     a pizza." You can't provide utterances for built-in intents.
//   - Information to be gathered. You specify slots for the information that you
//     bot requests from the user. You can specify standard slot types, such as date
//     and time, or custom slot types for your application.
//   - How the intent is fulfilled. You can provide a Lambda function or configure
//     the intent to return the intent information to your client application. If you
//     use a Lambda function, Amazon Lex invokes the function when all of the intent
//     information is available.
//   - A confirmation prompt to send to the user to confirm an intent. For
//     example, "Shall I order your pizza?"
//   - A conclusion statement to send to the user after the intent is fulfilled.
//     For example, "I ordered your pizza."
//   - A follow-up prompt that asks the user for additional activity. For example,
//     "Do you want a drink with your pizza?"
func (c *Client) CreateIntent(ctx context.Context, params *CreateIntentInput, optFns ...func(*Options)) (*CreateIntentOutput, error) {
	if params == nil {
		params = &CreateIntentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIntent", params, optFns, c.addOperationCreateIntentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIntentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIntentInput struct {

	// The identifier of the bot associated with this intent.
	//
	// This member is required.
	BotId *string

	// The version of the bot associated with this intent.
	//
	// This member is required.
	BotVersion *string

	// The name of the intent. Intent names must be unique in the locale that contains
	// the intent and cannot match the name of any built-in intent.
	//
	// This member is required.
	IntentName *string

	// The identifier of the language and locale where this intent is used. All of the
	// bots, slot types, and slots used by the intent must have the same locale. For
	// more information, see Supported languages (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html)
	// .
	//
	// This member is required.
	LocaleId *string

	// A description of the intent. Use the description to help identify the intent in
	// lists.
	Description *string

	// Specifies that Amazon Lex invokes the alias Lambda function for each user
	// input. You can invoke this Lambda function to personalize user interaction. For
	// example, suppose that your bot determines that the user's name is John. You
	// Lambda function might retrieve John's information from a backend database and
	// prepopulate some of the values. For example, if you find that John is gluten
	// intolerant, you might set the corresponding intent slot, glutenIntolerant to
	// true . You might find John's phone number and set the corresponding session
	// attribute.
	DialogCodeHook *types.DialogCodeHookSettings

	// Specifies that Amazon Lex invokes the alias Lambda function when the intent is
	// ready for fulfillment. You can invoke this function to complete the bot's
	// transaction with the user. For example, in a pizza ordering bot, the Lambda
	// function can look up the closest pizza restaurant to the customer's location and
	// then place an order on the customer's behalf.
	FulfillmentCodeHook *types.FulfillmentCodeHookSettings

	// Configuration settings for the response that is sent to the user at the
	// beginning of a conversation, before eliciting slot values.
	InitialResponseSetting *types.InitialResponseSetting

	// A list of contexts that must be active for this intent to be considered by
	// Amazon Lex. When an intent has an input context list, Amazon Lex only considers
	// using the intent in an interaction with the user when the specified contexts are
	// included in the active context list for the session. If the contexts are not
	// active, then Amazon Lex will not use the intent. A context can be automatically
	// activated using the outputContexts property or it can be set at runtime. For
	// example, if there are two intents with different input contexts that respond to
	// the same utterances, only the intent with the active context will respond. An
	// intent may have up to 5 input contexts. If an intent has multiple input
	// contexts, all of the contexts must be active to consider the intent.
	InputContexts []types.InputContext

	// Sets the response that Amazon Lex sends to the user when the intent is closed.
	IntentClosingSetting *types.IntentClosingSetting

	// Provides prompts that Amazon Lex sends to the user to confirm the completion of
	// an intent. If the user answers "no," the settings contain a statement that is
	// sent to the user to end the intent.
	IntentConfirmationSetting *types.IntentConfirmationSetting

	// Configuration information required to use the AMAZON.KendraSearchIntent intent
	// to connect to an Amazon Kendra index. The AMAZON.KendraSearchIntent intent is
	// called when Amazon Lex can't determine another intent to invoke.
	KendraConfiguration *types.KendraConfiguration

	// A lists of contexts that the intent activates when it is fulfilled. You can use
	// an output context to indicate the intents that Amazon Lex should consider for
	// the next turn of the conversation with a customer. When you use the
	// outputContextsList property, all of the contexts specified in the list are
	// activated when the intent is fulfilled. You can set up to 10 output contexts.
	// You can also set the number of conversation turns that the context should be
	// active, or the length of time that the context should be active.
	OutputContexts []types.OutputContext

	// A unique identifier for the built-in intent to base this intent on.
	ParentIntentSignature *string

	// An array of strings that a user might say to signal the intent. For example, "I
	// want a pizza", or "I want a {PizzaSize} pizza". In an utterance, slot names are
	// enclosed in curly braces ("{", "}") to indicate where they should be displayed
	// in the utterance shown to the user..
	SampleUtterances []types.SampleUtterance

	noSmithyDocumentSerde
}

type CreateIntentOutput struct {

	// The identifier of the bot associated with the intent.
	BotId *string

	// The version of the bot associated with the intent.
	BotVersion *string

	// A timestamp of the date and time that the intent was created.
	CreationDateTime *time.Time

	// The description specified for the intent.
	Description *string

	// The dialog Lambda function specified for the intent.
	DialogCodeHook *types.DialogCodeHookSettings

	// The fulfillment Lambda function specified for the intent.
	FulfillmentCodeHook *types.FulfillmentCodeHookSettings

	// Configuration settings for the response that is sent to the user at the
	// beginning of a conversation, before eliciting slot values.
	InitialResponseSetting *types.InitialResponseSetting

	// The list of input contexts specified for the intent.
	InputContexts []types.InputContext

	// The closing setting specified for the intent.
	IntentClosingSetting *types.IntentClosingSetting

	// The confirmation setting specified for the intent.
	IntentConfirmationSetting *types.IntentConfirmationSetting

	// A unique identifier for the intent.
	IntentId *string

	// The name specified for the intent.
	IntentName *string

	// Configuration for searching a Amazon Kendra index specified for the intent.
	KendraConfiguration *types.KendraConfiguration

	// The locale that the intent is specified to use.
	LocaleId *string

	// The list of output contexts specified for the intent.
	OutputContexts []types.OutputContext

	// The signature of the parent intent specified for the intent.
	ParentIntentSignature *string

	// The sample utterances specified for the intent.
	SampleUtterances []types.SampleUtterance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIntentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateIntent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIntent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateIntent"); err != nil {
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
	if err = addOpCreateIntentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIntent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateIntent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateIntent",
	}
}
