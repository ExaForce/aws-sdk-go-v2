// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// This operation downloads the output of the job you initiated using InitiateJob .
// Depending on the job type you specified when you initiated the job, the output
// will be either the content of an archive or a vault inventory. You can download
// all the job output or download a portion of the output by specifying a byte
// range. In the case of an archive retrieval job, depending on the byte range you
// specify, Amazon S3 Glacier (Glacier) returns the checksum for the portion of the
// data. You can compute the checksum on the client and verify that the values
// match to ensure the portion you downloaded is the correct data. A job ID will
// not expire for at least 24 hours after Glacier completes the job. That a byte
// range. For both archive and inventory retrieval jobs, you should verify the
// downloaded size against the size returned in the headers from the Get Job Output
// response. For archive retrieval jobs, you should also verify that the size is
// what you expected. If you download a portion of the output, the expected size is
// based on the range of bytes you specified. For example, if you specify a range
// of bytes=0-1048575 , you should verify your download size is 1,048,576 bytes. If
// you download an entire archive, the expected size is the size of the archive
// when you uploaded it to Amazon S3 Glacier The expected size is also returned in
// the headers from the Get Job Output response. In the case of an archive
// retrieval job, depending on the byte range you specify, Glacier returns the
// checksum for the portion of the data. To ensure the portion you downloaded is
// the correct data, compute the checksum on the client, verify that the values
// match, and verify that the size is what you expected. A job ID does not expire
// for at least 24 hours after Glacier completes the job. That is, you can download
// the job output within the 24 hours period after Amazon Glacier completes the
// job. An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see Access Control Using AWS Identity
// and Access Management (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html)
// . For conceptual information and the underlying REST API, see Downloading a
// Vault Inventory (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-inventory.html)
// , Downloading an Archive (https://docs.aws.amazon.com/amazonglacier/latest/dev/downloading-an-archive.html)
// , and Get Job Output  (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-job-output-get.html)
func (c *Client) GetJobOutput(ctx context.Context, params *GetJobOutputInput, optFns ...func(*Options)) (*GetJobOutputOutput, error) {
	if params == nil {
		params = &GetJobOutputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetJobOutput", params, optFns, c.addOperationGetJobOutputMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetJobOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for downloading output of an Amazon S3 Glacier job.
type GetJobOutputInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single ' - ' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The job ID whose data is downloaded.
	//
	// This member is required.
	JobId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// The range of bytes to retrieve from the output. For example, if you want to
	// download the first 1,048,576 bytes, specify the range as bytes=0-1048575 . By
	// default, this operation downloads the entire output. If the job output is large,
	// then you can use a range to retrieve a portion of the output. This allows you to
	// download the entire output in smaller chunks of bytes. For example, suppose you
	// have 1 GB of job output you want to download and you decide to download 128 MB
	// chunks of data at a time, which is a total of eight Get Job Output requests. You
	// use the following process to download the job output:
	//   - Download a 128 MB chunk of output by specifying the appropriate byte range.
	//   Verify that all 128 MB of data was received.
	//   - Along with the data, the response includes a SHA256 tree hash of the
	//   payload. You compute the checksum of the payload on the client and compare it
	//   with the checksum you received in the response to ensure you received all the
	//   expected data.
	//   - Repeat steps 1 and 2 for all the eight 128 MB chunks of output data, each
	//   time specifying the appropriate byte range.
	//   - After downloading all the parts of the job output, you have a list of eight
	//   checksum values. Compute the tree hash of these values to find the checksum of
	//   the entire output. Using the DescribeJob API, obtain job information of the
	//   job that provided you the output. The response includes the checksum of the
	//   entire archive stored in Amazon S3 Glacier. You compare this value with the
	//   checksum you computed to ensure you have downloaded the entire archive content
	//   with no errors.
	Range *string

	noSmithyDocumentSerde
}

// Contains the Amazon S3 Glacier response to your request.
type GetJobOutputOutput struct {

	// Indicates the range units accepted. For more information, see RFC2616 (http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html)
	// .
	AcceptRanges *string

	// The description of an archive.
	ArchiveDescription *string

	// The job data, either archive data or inventory data.
	Body io.ReadCloser

	// The checksum of the data in the response. This header is returned only when
	// retrieving the output for an archive retrieval job. Furthermore, this header
	// appears only under the following conditions:
	//   - You get the entire range of the archive.
	//   - You request a range to return of the archive that starts and ends on a
	//   multiple of 1 MB. For example, if you have an 3.1 MB archive and you specify a
	//   range to return that starts at 1 MB and ends at 2 MB, then the
	//   x-amz-sha256-tree-hash is returned as a response header.
	//   - You request a range of the archive to return that starts on a multiple of 1
	//   MB and goes to the end of the archive. For example, if you have a 3.1 MB archive
	//   and you specify a range that starts at 2 MB and ends at 3.1 MB (the end of the
	//   archive), then the x-amz-sha256-tree-hash is returned as a response header.
	Checksum *string

	// The range of bytes returned by Amazon S3 Glacier. If only partial output is
	// downloaded, the response provides the range of bytes Amazon S3 Glacier returned.
	// For example, bytes 0-1048575/8388608 returns the first 1 MB from 8 MB.
	ContentRange *string

	// The Content-Type depends on whether the job output is an archive or a vault
	// inventory. For archive data, the Content-Type is application/octet-stream. For
	// vault inventory, if you requested CSV format when you initiated the job, the
	// Content-Type is text/csv. Otherwise, by default, vault inventory is returned as
	// JSON, and the Content-Type is application/json.
	ContentType *string

	// The HTTP response code for a job output request. The value depends on whether a
	// range was specified in the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetJobOutputMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetJobOutput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetJobOutput{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetJobOutput"); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetJobOutputValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetJobOutput(options.Region), middleware.Before); err != nil {
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
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
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

func newServiceMetadataMiddleware_opGetJobOutput(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetJobOutput",
	}
}
