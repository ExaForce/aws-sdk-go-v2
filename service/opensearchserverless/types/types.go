// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless/document"
	smithydocument "github.com/aws/smithy-go/document"
)

// Details about an OpenSearch Serverless access policy.
type AccessPolicyDetail struct {

	// The date the policy was created.
	CreatedDate *int64

	// The description of the policy.
	Description *string

	// The timestamp of when the policy was last modified.
	LastModifiedDate *int64

	// The name of the policy.
	Name *string

	// The JSON policy document without any whitespaces.
	Policy document.Interface

	// The version of the policy.
	PolicyVersion *string

	// The type of access policy.
	Type AccessPolicyType

	noSmithyDocumentSerde
}

// Statistics for an OpenSearch Serverless access policy.
type AccessPolicyStats struct {

	// The number of data access policies in the current account.
	DataPolicyCount *int64

	noSmithyDocumentSerde
}

// A summary of the data access policy.
type AccessPolicySummary struct {

	// The Epoch time when the access policy was created.
	CreatedDate *int64

	// The description of the access policy.
	Description *string

	// The date and time when the collection was last modified.
	LastModifiedDate *int64

	// The name of the access policy.
	Name *string

	// The version of the policy.
	PolicyVersion *string

	// The type of access policy. Currently, the only available type is data .
	Type AccessPolicyType

	noSmithyDocumentSerde
}

// OpenSearch Serverless-related information for the current account.
type AccountSettingsDetail struct {

	// The maximum capacity limits for all OpenSearch Serverless collections, in
	// OpenSearch Compute Units (OCUs). These limits are used to scale your collections
	// based on the current workload. For more information, see [Managing capacity limits for Amazon OpenSearch Serverless].
	//
	// [Managing capacity limits for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-scaling.html
	CapacityLimits *CapacityLimits

	noSmithyDocumentSerde
}

// The maximum capacity limits for all OpenSearch Serverless collections, in
// OpenSearch Compute Units (OCUs). These limits are used to scale your collections
// based on the current workload. For more information, see [Managing capacity limits for Amazon OpenSearch Serverless].
//
// [Managing capacity limits for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-scaling.html
type CapacityLimits struct {

	// The maximum indexing capacity for collections.
	MaxIndexingCapacityInOCU *int32

	// The maximum search capacity for collections.
	MaxSearchCapacityInOCU *int32

	noSmithyDocumentSerde
}

// Details about each OpenSearch Serverless collection, including the collection
// endpoint and the OpenSearch Dashboards endpoint.
type CollectionDetail struct {

	// The Amazon Resource Name (ARN) of the collection.
	Arn *string

	// Collection-specific endpoint used to submit index, search, and data upload
	// requests to an OpenSearch Serverless collection.
	CollectionEndpoint *string

	// The Epoch time when the collection was created.
	CreatedDate *int64

	// Collection-specific endpoint used to access OpenSearch Dashboards.
	DashboardEndpoint *string

	// A description of the collection.
	Description *string

	// A failure code associated with the request.
	FailureCode *string

	// A message associated with the failure code.
	FailureMessage *string

	// A unique identifier for the collection.
	Id *string

	// The ARN of the Amazon Web Services KMS key used to encrypt the collection.
	KmsKeyArn *string

	// The date and time when the collection was last modified.
	LastModifiedDate *int64

	// The name of the collection.
	Name *string

	// Details about an OpenSearch Serverless collection.
	StandbyReplicas StandbyReplicas

	// The current status of the collection.
	Status CollectionStatus

	// The type of collection.
	Type CollectionType

	noSmithyDocumentSerde
}

// Error information for an OpenSearch Serverless request.
type CollectionErrorDetail struct {

	// The error code for the request. For example, NOT_FOUND .
	ErrorCode *string

	// A description of the error. For example, The specified Collection is not found.
	ErrorMessage *string

	// If the request contains collection IDs, the response includes the IDs provided
	// in the request.
	Id *string

	// If the request contains collection names, the response includes the names
	// provided in the request.
	Name *string

	noSmithyDocumentSerde
}

// A list of filter keys that you can use for LIST, UPDATE, and DELETE requests to
// OpenSearch Serverless collections.
type CollectionFilters struct {

	// The name of the collection.
	Name *string

	// The current status of the collection.
	Status CollectionStatus

	noSmithyDocumentSerde
}

// Details about each OpenSearch Serverless collection.
type CollectionSummary struct {

	// The Amazon Resource Name (ARN) of the collection.
	Arn *string

	// The unique identifier of the collection.
	Id *string

	// The name of the collection.
	Name *string

	// The current status of the collection.
	Status CollectionStatus

	noSmithyDocumentSerde
}

// Details about the created OpenSearch Serverless collection.
type CreateCollectionDetail struct {

	// The Amazon Resource Name (ARN) of the collection.
	Arn *string

	// The Epoch time when the collection was created.
	CreatedDate *int64

	// A description of the collection.
	Description *string

	// The unique identifier of the collection.
	Id *string

	// The Amazon Resource Name (ARN) of the KMS key with which to encrypt the
	// collection.
	KmsKeyArn *string

	// The date and time when the collection was last modified.
	LastModifiedDate *int64

	// The name of the collection.
	Name *string

	// Creates details about an OpenSearch Serverless collection.
	StandbyReplicas StandbyReplicas

	// The current status of the collection.
	Status CollectionStatus

	// The type of collection.
	Type CollectionType

	noSmithyDocumentSerde
}

// Creation details for an OpenSearch Serverless-managed interface endpoint. For
// more information, see [Access Amazon OpenSearch Serverless using an interface endpoint].
//
// [Access Amazon OpenSearch Serverless using an interface endpoint]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-vpc.html
type CreateVpcEndpointDetail struct {

	// The unique identifier of the endpoint.
	Id *string

	// The name of the endpoint.
	Name *string

	// The current status in the endpoint creation process.
	Status VpcEndpointStatus

	noSmithyDocumentSerde
}

// Details about a deleted OpenSearch Serverless collection.
type DeleteCollectionDetail struct {

	// The unique identifier of the collection.
	Id *string

	// The name of the collection.
	Name *string

	// The current status of the collection.
	Status CollectionStatus

	noSmithyDocumentSerde
}

// Deletion details for an OpenSearch Serverless-managed interface endpoint.
type DeleteVpcEndpointDetail struct {

	// The unique identifier of the endpoint.
	Id *string

	// The name of the endpoint.
	Name *string

	// The current status of the endpoint deletion process.
	Status VpcEndpointStatus

	noSmithyDocumentSerde
}

// Error information for an OpenSearch Serverless request.
type EffectiveLifecyclePolicyDetail struct {

	// The minimum number of index retention days set. That is an optional param that
	// will return as true if the minimum number of days or hours is not set to a
	// index resource.
	NoMinRetentionPeriod *bool

	// The name of the lifecycle policy.
	PolicyName *string

	// The name of the OpenSearch Serverless index resource.
	Resource *string

	// The type of OpenSearch Serverless resource. Currently, the only supported
	// resource is index .
	ResourceType ResourceType

	// The minimum number of index retention in days or hours. This is an optional
	// parameter that will return only if it’s set.
	RetentionPeriod *string

	// The type of lifecycle policy.
	Type LifecyclePolicyType

	noSmithyDocumentSerde
}

// Error information for an OpenSearch Serverless request.
type EffectiveLifecyclePolicyErrorDetail struct {

	// The error code for the request.
	ErrorCode *string

	// A description of the error. For example, The specified Index resource is not
	// found .
	ErrorMessage *string

	// The name of OpenSearch Serverless index resource.
	Resource *string

	// The type of lifecycle policy.
	Type LifecyclePolicyType

	noSmithyDocumentSerde
}

// Details about an OpenSearch Serverless lifecycle policy.
type LifecyclePolicyDetail struct {

	// The date the lifecycle policy was created.
	CreatedDate *int64

	// The description of the lifecycle policy.
	Description *string

	// The timestamp of when the lifecycle policy was last modified.
	LastModifiedDate *int64

	// The name of the lifecycle policy.
	Name *string

	// The JSON policy document without any whitespaces.
	Policy document.Interface

	// The version of the lifecycle policy.
	PolicyVersion *string

	// The type of lifecycle policy.
	Type LifecyclePolicyType

	noSmithyDocumentSerde
}

// Error information for an OpenSearch Serverless request.
type LifecyclePolicyErrorDetail struct {

	// The error code for the request. For example, NOT_FOUND .
	ErrorCode *string

	// A description of the error. For example, The specified Lifecycle Policy is not
	// found .
	ErrorMessage *string

	// The name of the lifecycle policy.
	Name *string

	// The type of lifecycle policy.
	Type LifecyclePolicyType

	noSmithyDocumentSerde
}

// The unique identifiers of policy types and policy names.
type LifecyclePolicyIdentifier struct {

	// The name of the lifecycle policy.
	//
	// This member is required.
	Name *string

	// The type of lifecycle policy.
	//
	// This member is required.
	Type LifecyclePolicyType

	noSmithyDocumentSerde
}

// The unique identifiers of policy types and resource names.
type LifecyclePolicyResourceIdentifier struct {

	// The name of the OpenSearch Serverless ilndex resource.
	//
	// This member is required.
	Resource *string

	// The type of lifecycle policy.
	//
	// This member is required.
	Type LifecyclePolicyType

	noSmithyDocumentSerde
}

// Statistics for an OpenSearch Serverless lifecycle policy.
type LifecyclePolicyStats struct {

	// The number of retention lifecycle policies in the current account.
	RetentionPolicyCount *int64

	noSmithyDocumentSerde
}

// A summary of the lifecycle policy.
type LifecyclePolicySummary struct {

	// The Epoch time when the lifecycle policy was created.
	CreatedDate *int64

	// The description of the lifecycle policy.
	Description *string

	// The date and time when the lifecycle policy was last modified.
	LastModifiedDate *int64

	// The name of the lifecycle policy.
	Name *string

	// The version of the lifecycle policy.
	PolicyVersion *string

	// The type of lifecycle policy.
	Type LifecyclePolicyType

	noSmithyDocumentSerde
}

// Describes SAML options for an OpenSearch Serverless security configuration in
// the form of a key-value map.
type SamlConfigOptions struct {

	// The XML IdP metadata file generated from your identity provider.
	//
	// This member is required.
	Metadata *string

	// The group attribute for this SAML integration.
	GroupAttribute *string

	// The session timeout, in minutes. Default is 60 minutes (12 hours).
	SessionTimeout *int32

	// A user attribute for this SAML integration.
	UserAttribute *string

	noSmithyDocumentSerde
}

// Details about a security configuration for OpenSearch Serverless.
type SecurityConfigDetail struct {

	// The version of the security configuration.
	ConfigVersion *string

	// The date the configuration was created.
	CreatedDate *int64

	// The description of the security configuration.
	Description *string

	// The unique identifier of the security configuration.
	Id *string

	// The timestamp of when the configuration was last modified.
	LastModifiedDate *int64

	// SAML options for the security configuration in the form of a key-value map.
	SamlOptions *SamlConfigOptions

	// The type of security configuration.
	Type SecurityConfigType

	noSmithyDocumentSerde
}

// Statistics for an OpenSearch Serverless security configuration.
type SecurityConfigStats struct {

	// The number of security configurations in the current account.
	SamlConfigCount *int64

	noSmithyDocumentSerde
}

// A summary of a security configuration for OpenSearch Serverless.
type SecurityConfigSummary struct {

	// The version of the security configuration.
	ConfigVersion *string

	// The Epoch time when the security configuration was created.
	CreatedDate *int64

	// The description of the security configuration.
	Description *string

	// The unique identifier of the security configuration.
	Id *string

	// The timestamp of when the configuration was last modified.
	LastModifiedDate *int64

	// The type of security configuration.
	Type SecurityConfigType

	noSmithyDocumentSerde
}

// Details about an OpenSearch Serverless security policy.
type SecurityPolicyDetail struct {

	// The date the policy was created.
	CreatedDate *int64

	// The description of the security policy.
	Description *string

	// The timestamp of when the policy was last modified.
	LastModifiedDate *int64

	// The name of the policy.
	Name *string

	// The JSON policy document without any whitespaces.
	Policy document.Interface

	// The version of the policy.
	PolicyVersion *string

	// The type of security policy.
	Type SecurityPolicyType

	noSmithyDocumentSerde
}

// Statistics for an OpenSearch Serverless security policy.
type SecurityPolicyStats struct {

	// The number of encryption policies in the current account.
	EncryptionPolicyCount *int64

	// The number of network policies in the current account.
	NetworkPolicyCount *int64

	noSmithyDocumentSerde
}

// A summary of a security policy for OpenSearch Serverless.
type SecurityPolicySummary struct {

	// The date the policy was created.
	CreatedDate *int64

	// The description of the security policy.
	Description *string

	// The timestamp of when the policy was last modified.
	LastModifiedDate *int64

	// The name of the policy.
	Name *string

	// The version of the policy.
	PolicyVersion *string

	// The type of security policy.
	Type SecurityPolicyType

	noSmithyDocumentSerde
}

// A map of key-value pairs associated to an OpenSearch Serverless resource.
type Tag struct {

	// The key to use in the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// Details about an updated OpenSearch Serverless collection.
type UpdateCollectionDetail struct {

	// The Amazon Resource Name (ARN) of the collection.
	Arn *string

	// The date and time when the collection was created.
	CreatedDate *int64

	// The description of the collection.
	Description *string

	// The unique identifier of the collection.
	Id *string

	// The date and time when the collection was last modified.
	LastModifiedDate *int64

	// The name of the collection.
	Name *string

	// The current status of the collection.
	Status CollectionStatus

	// The collection type.
	Type CollectionType

	noSmithyDocumentSerde
}

// Update details for an OpenSearch Serverless-managed interface endpoint.
type UpdateVpcEndpointDetail struct {

	// The unique identifier of the endpoint.
	Id *string

	// The timestamp of when the endpoint was last modified.
	LastModifiedDate *int64

	// The name of the endpoint.
	Name *string

	// The unique identifiers of the security groups that define the ports, protocols,
	// and sources for inbound traffic that you are authorizing into your endpoint.
	SecurityGroupIds []string

	// The current status of the endpoint update process.
	Status VpcEndpointStatus

	// The ID of the subnets from which you access OpenSearch Serverless.
	SubnetIds []string

	noSmithyDocumentSerde
}

// Details about an OpenSearch Serverless-managed interface endpoint.
type VpcEndpointDetail struct {

	// The date the endpoint was created.
	CreatedDate *int64

	// A failure code associated with the request.
	FailureCode *string

	// A message associated with the failure code.
	FailureMessage *string

	// The unique identifier of the endpoint.
	Id *string

	// The name of the endpoint.
	Name *string

	// The unique identifiers of the security groups that define the ports, protocols,
	// and sources for inbound traffic that you are authorizing into your endpoint.
	SecurityGroupIds []string

	// The current status of the endpoint.
	Status VpcEndpointStatus

	// The ID of the subnets from which you access OpenSearch Serverless.
	SubnetIds []string

	// The ID of the VPC from which you access OpenSearch Serverless.
	VpcId *string

	noSmithyDocumentSerde
}

// Error information for a failed BatchGetVpcEndpoint request.
type VpcEndpointErrorDetail struct {

	// The error code for the failed request.
	ErrorCode *string

	// An error message describing the reason for the failure.
	ErrorMessage *string

	// The unique identifier of the VPC endpoint.
	Id *string

	noSmithyDocumentSerde
}

// Filter the results of a ListVpcEndpoints request.
type VpcEndpointFilters struct {

	// The current status of the endpoint.
	Status VpcEndpointStatus

	noSmithyDocumentSerde
}

// The VPC endpoint object.
type VpcEndpointSummary struct {

	// The unique identifier of the endpoint.
	Id *string

	// The name of the endpoint.
	Name *string

	// The current status of the endpoint.
	Status VpcEndpointStatus

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
