// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A query filter used by SearchUsers . This filter object provides the attribute
// name and attribute value to search users or groups.
type Filter struct {

	// The attribute path that is used to specify which attribute name to search.
	// Length limit is 255 characters. For example, UserName is a valid attribute path
	// for the SearchUsers API.
	//
	// This member is required.
	AttributePath *string

	// Represents the data for an attribute. Each attribute value is described as a
	// name-value pair.
	//
	// This member is required.
	AttributeValue *string

	noSmithyDocumentSerde
}

// A user object that contains the metadata and attributes for a specified user.
type User struct {

	// The identifier for a user in the identity store.
	//
	// This member is required.
	UserId *string

	// A Boolean value representing whether the user is active.
	Active bool

	// Metadata information for a user.
	Meta *UserInfo

	// A unique string used to identify the user. The length limit is 128 characters.
	// This value can consist of letters, accented characters, symbols, numbers, and
	// punctuation. This value is specified at the time the user is created and stored
	// as an attribute of the user object in the identity store.
	UserName *string

	noSmithyDocumentSerde
}

// Metadata information about a user.
type UserInfo struct {

	// The time when the user was created.
	CreatedAt *time.Time

	// Entity that created the user.
	CreatedBy *string

	// The time when the user was updated.
	UpdatedAt *time.Time

	// Entity that updated the user.
	UpdatedBy *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
