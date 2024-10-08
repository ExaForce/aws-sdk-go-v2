// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DeploymentFilterKey string

// Enum values for DeploymentFilterKey
const (
	DeploymentFilterKeyWorkloadName     DeploymentFilterKey = "WORKLOAD_NAME"
	DeploymentFilterKeyDeploymentStatus DeploymentFilterKey = "DEPLOYMENT_STATUS"
)

// Values returns all known values for DeploymentFilterKey. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentFilterKey) Values() []DeploymentFilterKey {
	return []DeploymentFilterKey{
		"WORKLOAD_NAME",
		"DEPLOYMENT_STATUS",
	}
}

type DeploymentStatus string

// Enum values for DeploymentStatus
const (
	DeploymentStatusCompleted        DeploymentStatus = "COMPLETED"
	DeploymentStatusCreating         DeploymentStatus = "CREATING"
	DeploymentStatusDeleteInProgress DeploymentStatus = "DELETE_IN_PROGRESS"
	DeploymentStatusDeleteInitiating DeploymentStatus = "DELETE_INITIATING"
	DeploymentStatusDeleteFailed     DeploymentStatus = "DELETE_FAILED"
	DeploymentStatusDeleted          DeploymentStatus = "DELETED"
	DeploymentStatusFailed           DeploymentStatus = "FAILED"
	DeploymentStatusInProgress       DeploymentStatus = "IN_PROGRESS"
	DeploymentStatusValidating       DeploymentStatus = "VALIDATING"
)

// Values returns all known values for DeploymentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentStatus) Values() []DeploymentStatus {
	return []DeploymentStatus{
		"COMPLETED",
		"CREATING",
		"DELETE_IN_PROGRESS",
		"DELETE_INITIATING",
		"DELETE_FAILED",
		"DELETED",
		"FAILED",
		"IN_PROGRESS",
		"VALIDATING",
	}
}

type EventStatus string

// Enum values for EventStatus
const (
	EventStatusCanceled   EventStatus = "CANCELED"
	EventStatusCanceling  EventStatus = "CANCELING"
	EventStatusCompleted  EventStatus = "COMPLETED"
	EventStatusCreated    EventStatus = "CREATED"
	EventStatusFailed     EventStatus = "FAILED"
	EventStatusInProgress EventStatus = "IN_PROGRESS"
	EventStatusPending    EventStatus = "PENDING"
	EventStatusTimedOut   EventStatus = "TIMED_OUT"
)

// Values returns all known values for EventStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (EventStatus) Values() []EventStatus {
	return []EventStatus{
		"CANCELED",
		"CANCELING",
		"COMPLETED",
		"CREATED",
		"FAILED",
		"IN_PROGRESS",
		"PENDING",
		"TIMED_OUT",
	}
}

type WorkloadDeploymentPatternStatus string

// Enum values for WorkloadDeploymentPatternStatus
const (
	WorkloadDeploymentPatternStatusActive   WorkloadDeploymentPatternStatus = "ACTIVE"
	WorkloadDeploymentPatternStatusInactive WorkloadDeploymentPatternStatus = "INACTIVE"
	WorkloadDeploymentPatternStatusDisabled WorkloadDeploymentPatternStatus = "DISABLED"
	WorkloadDeploymentPatternStatusDeleted  WorkloadDeploymentPatternStatus = "DELETED"
)

// Values returns all known values for WorkloadDeploymentPatternStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WorkloadDeploymentPatternStatus) Values() []WorkloadDeploymentPatternStatus {
	return []WorkloadDeploymentPatternStatus{
		"ACTIVE",
		"INACTIVE",
		"DISABLED",
		"DELETED",
	}
}

type WorkloadStatus string

// Enum values for WorkloadStatus
const (
	WorkloadStatusActive   WorkloadStatus = "ACTIVE"
	WorkloadStatusInactive WorkloadStatus = "INACTIVE"
	WorkloadStatusDisabled WorkloadStatus = "DISABLED"
	WorkloadStatusDeleted  WorkloadStatus = "DELETED"
)

// Values returns all known values for WorkloadStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (WorkloadStatus) Values() []WorkloadStatus {
	return []WorkloadStatus{
		"ACTIVE",
		"INACTIVE",
		"DISABLED",
		"DELETED",
	}
}
