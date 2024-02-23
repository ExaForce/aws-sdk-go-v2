// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates some of the parameters for an existing connector. Provide the
// ConnectorId for the connector that you want to update, along with the new values
// for the parameters to update.
func (c *Client) UpdateConnector(ctx context.Context, params *UpdateConnectorInput, optFns ...func(*Options)) (*UpdateConnectorOutput, error) {
	if params == nil {
		params = &UpdateConnectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConnector", params, optFns, c.addOperationUpdateConnectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateConnectorInput struct {

	// The unique identifier for the connector.
	//
	// This member is required.
	ConnectorId *string

	// Connectors are used to send files using either the AS2 or SFTP protocol. For
	// the access role, provide the Amazon Resource Name (ARN) of the Identity and
	// Access Management role to use. For AS2 connectors With AS2, you can send files
	// by calling StartFileTransfer and specifying the file paths in the request
	// parameter, SendFilePaths . We use the file’s parent directory (for example, for
	// --send-file-paths /bucket/dir/file.txt , parent directory is /bucket/dir/ ) to
	// temporarily store a processed AS2 message file, store the MDN when we receive
	// them from the partner, and write a final JSON file containing relevant metadata
	// of the transmission. So, the AccessRole needs to provide read and write access
	// to the parent directory of the file location used in the StartFileTransfer
	// request. Additionally, you need to provide read and write access to the parent
	// directory of the files that you intend to send with StartFileTransfer . If you
	// are using Basic authentication for your AS2 connector, the access role requires
	// the secretsmanager:GetSecretValue permission for the secret. If the secret is
	// encrypted using a customer-managed key instead of the Amazon Web Services
	// managed key in Secrets Manager, then the role also needs the kms:Decrypt
	// permission for that key. For SFTP connectors Make sure that the access role
	// provides read and write access to the parent directory of the file location
	// that's used in the StartFileTransfer request. Additionally, make sure that the
	// role provides secretsmanager:GetSecretValue permission to Secrets Manager.
	AccessRole *string

	// A structure that contains the parameters for an AS2 connector object.
	As2Config *types.As2ConnectorConfig

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// that allows a connector to turn on CloudWatch logging for Amazon S3 events. When
	// set, you can view connector activity in your CloudWatch logs.
	LoggingRole *string

	// A structure that contains the parameters for an SFTP connector object.
	SftpConfig *types.SftpConnectorConfig

	// The URL of the partner's AS2 or SFTP endpoint.
	Url *string

	noSmithyDocumentSerde
}

type UpdateConnectorOutput struct {

	// Returns the identifier of the connector object that you are updating.
	//
	// This member is required.
	ConnectorId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConnectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateConnector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateConnector{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateConnector"); err != nil {
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
	if err = addOpUpdateConnectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConnector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConnector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateConnector",
	}
}
