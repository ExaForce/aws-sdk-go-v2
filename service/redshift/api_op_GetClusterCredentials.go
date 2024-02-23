// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns a database user name and temporary password with temporary
// authorization to log on to an Amazon Redshift database. The action returns the
// database user name prefixed with IAM: if AutoCreate is False or IAMA: if
// AutoCreate is True . You can optionally specify one or more database user groups
// that the user will join at log on. By default, the temporary credentials expire
// in 900 seconds. You can optionally specify a duration between 900 seconds (15
// minutes) and 3600 seconds (60 minutes). For more information, see Using IAM
// Authentication to Generate Database User Credentials (https://docs.aws.amazon.com/redshift/latest/mgmt/generating-user-credentials.html)
// in the Amazon Redshift Cluster Management Guide. The Identity and Access
// Management (IAM) user or role that runs GetClusterCredentials must have an IAM
// policy attached that allows access to all necessary actions and resources. For
// more information about permissions, see Resource Policies for
// GetClusterCredentials (https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-identity-based.html#redshift-policy-resources.getclustercredentials-resources)
// in the Amazon Redshift Cluster Management Guide. If the DbGroups parameter is
// specified, the IAM policy must allow the redshift:JoinGroup action with access
// to the listed dbgroups . In addition, if the AutoCreate parameter is set to True
// , then the policy must include the redshift:CreateClusterUser permission. If
// the DbName parameter is specified, the IAM policy must allow access to the
// resource dbname for the specified database name.
func (c *Client) GetClusterCredentials(ctx context.Context, params *GetClusterCredentialsInput, optFns ...func(*Options)) (*GetClusterCredentialsOutput, error) {
	if params == nil {
		params = &GetClusterCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetClusterCredentials", params, optFns, c.addOperationGetClusterCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetClusterCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request parameters to get cluster credentials.
type GetClusterCredentialsInput struct {

	// The name of a database user. If a user name matching DbUser exists in the
	// database, the temporary user credentials have the same permissions as the
	// existing user. If DbUser doesn't exist in the database and Autocreate is True ,
	// a new user is created using the value for DbUser with PUBLIC permissions. If a
	// database user matching the value for DbUser doesn't exist and Autocreate is
	// False , then the command succeeds but the connection attempt will fail because
	// the user doesn't exist in the database. For more information, see CREATE USER (https://docs.aws.amazon.com/redshift/latest/dg/r_CREATE_USER.html)
	// in the Amazon Redshift Database Developer Guide. Constraints:
	//   - Must be 1 to 64 alphanumeric characters or hyphens. The user name can't be
	//   PUBLIC .
	//   - Must contain uppercase or lowercase letters, numbers, underscore, plus
	//   sign, period (dot), at symbol (@), or hyphen.
	//   - First character must be a letter.
	//   - Must not contain a colon ( : ) or slash ( / ).
	//   - Cannot be a reserved word. A list of reserved words can be found in
	//   Reserved Words (http://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html)
	//   in the Amazon Redshift Database Developer Guide.
	//
	// This member is required.
	DbUser *string

	// Create a database user with the name specified for the user named in DbUser if
	// one does not exist.
	AutoCreate *bool

	// The unique identifier of the cluster that contains the database for which you
	// are requesting credentials. This parameter is case sensitive.
	ClusterIdentifier *string

	// The custom domain name for the cluster credentials.
	CustomDomainName *string

	// A list of the names of existing database groups that the user named in DbUser
	// will join for the current session, in addition to any group memberships for an
	// existing user. If not specified, a new user is added only to PUBLIC. Database
	// group name constraints
	//   - Must be 1 to 64 alphanumeric characters or hyphens
	//   - Must contain only lowercase letters, numbers, underscore, plus sign, period
	//   (dot), at symbol (@), or hyphen.
	//   - First character must be a letter.
	//   - Must not contain a colon ( : ) or slash ( / ).
	//   - Cannot be a reserved word. A list of reserved words can be found in
	//   Reserved Words (http://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html)
	//   in the Amazon Redshift Database Developer Guide.
	DbGroups []string

	// The name of a database that DbUser is authorized to log on to. If DbName is not
	// specified, DbUser can log on to any existing database. Constraints:
	//   - Must be 1 to 64 alphanumeric characters or hyphens
	//   - Must contain uppercase or lowercase letters, numbers, underscore, plus
	//   sign, period (dot), at symbol (@), or hyphen.
	//   - First character must be a letter.
	//   - Must not contain a colon ( : ) or slash ( / ).
	//   - Cannot be a reserved word. A list of reserved words can be found in
	//   Reserved Words (http://docs.aws.amazon.com/redshift/latest/dg/r_pg_keywords.html)
	//   in the Amazon Redshift Database Developer Guide.
	DbName *string

	// The number of seconds until the returned temporary password expires.
	// Constraint: minimum 900, maximum 3600. Default: 900
	DurationSeconds *int32

	noSmithyDocumentSerde
}

// Temporary credentials with authorization to log on to an Amazon Redshift
// database.
type GetClusterCredentialsOutput struct {

	// A temporary password that authorizes the user name returned by DbUser to log on
	// to the database DbName .
	DbPassword *string

	// A database user name that is authorized to log on to the database DbName using
	// the password DbPassword . If the specified DbUser exists in the database, the
	// new user name has the same database permissions as the the user named in DbUser.
	// By default, the user is added to PUBLIC. If the DbGroups parameter is specifed,
	// DbUser is added to the listed groups for any sessions created using these
	// credentials.
	DbUser *string

	// The date and time the password in DbPassword expires.
	Expiration *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetClusterCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetClusterCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetClusterCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetClusterCredentials"); err != nil {
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
	if err = addOpGetClusterCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetClusterCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetClusterCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetClusterCredentials",
	}
}
