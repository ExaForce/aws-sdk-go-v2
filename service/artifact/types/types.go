// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Account settings for the customer.
type AccountSettings struct {

	// Notification subscription status of the customer.
	NotificationSubscriptionStatus NotificationSubscriptionStatus

	noSmithyDocumentSerde
}

// Full detail for report resource metadata.
type ReportDetail struct {

	// Acceptance type for report.
	AcceptanceType AcceptanceType

	// ARN for the report resource.
	Arn *string

	// Category for the report resource.
	Category *string

	// Associated company name for the report resource.
	CompanyName *string

	// Timestamp indicating when the report resource was created.
	CreatedAt *time.Time

	// Timestamp indicating when the report resource was deleted.
	DeletedAt *time.Time

	// Description for the report resource.
	Description *string

	// Unique resource ID for the report resource.
	Id *string

	// Timestamp indicating when the report resource was last modified.
	LastModifiedAt *time.Time

	// Name for the report resource.
	Name *string

	// Timestamp indicating the report resource effective end.
	PeriodEnd *time.Time

	// Timestamp indicating the report resource effective start.
	PeriodStart *time.Time

	// Associated product name for the report resource.
	ProductName *string

	// Sequence number to enforce optimistic locking.
	SequenceNumber *int64

	// Series for the report resource.
	Series *string

	// Current state of the report resource
	State PublishedState

	// The message associated with the current upload state.
	StatusMessage *string

	// Unique resource ARN for term resource.
	TermArn *string

	// The current state of the document upload.
	UploadState UploadState

	// Version for the report resource.
	Version *int64

	noSmithyDocumentSerde
}

// Summary for report resource.
type ReportSummary struct {

	// ARN for the report resource.
	Arn *string

	// Category for the report resource.
	Category *string

	// Associated company name for the report resource.
	CompanyName *string

	// Description for the report resource.
	Description *string

	// Unique resource ID for the report resource.
	Id *string

	// Name for the report resource.
	Name *string

	// Timestamp indicating the report resource effective end.
	PeriodEnd *time.Time

	// Timestamp indicating the report resource effective start.
	PeriodStart *time.Time

	// Associated product name for the report resource.
	ProductName *string

	// Series for the report resource.
	Series *string

	// Current state of the report resource.
	State PublishedState

	// The message associated with the current upload state.
	StatusMessage *string

	// The current state of the document upload.
	UploadState UploadState

	// Version for the report resource.
	Version *int64

	noSmithyDocumentSerde
}

// Validation exception message and name.
type ValidationExceptionField struct {

	// Message describing why the field failed validation.
	//
	// This member is required.
	Message *string

	// Name of validation exception.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
