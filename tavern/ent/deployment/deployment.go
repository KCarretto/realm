// Code generated by entc, DO NOT EDIT.

package deployment

import (
	"time"
)

const (
	// Label holds the string label denoting the deployment type in the database.
	Label = "deployment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"
	// FieldQueuedAt holds the string denoting the queuedat field in the database.
	FieldQueuedAt = "queued_at"
	// FieldLastModifiedAt holds the string denoting the lastmodifiedat field in the database.
	FieldLastModifiedAt = "last_modified_at"
	// FieldStartedAt holds the string denoting the startedat field in the database.
	FieldStartedAt = "started_at"
	// FieldFinishedAt holds the string denoting the finishedat field in the database.
	FieldFinishedAt = "finished_at"
	// EdgeConfig holds the string denoting the config edge name in mutations.
	EdgeConfig = "config"
	// EdgeTarget holds the string denoting the target edge name in mutations.
	EdgeTarget = "target"
	// Table holds the table name of the deployment in the database.
	Table = "deployments"
	// ConfigTable is the table that holds the config relation/edge.
	ConfigTable = "deployments"
	// ConfigInverseTable is the table name for the DeploymentConfig entity.
	// It exists in this package in order to avoid circular dependency with the "deploymentconfig" package.
	ConfigInverseTable = "deployment_configs"
	// ConfigColumn is the table column denoting the config relation/edge.
	ConfigColumn = "deployment_config"
	// TargetTable is the table that holds the target relation/edge.
	TargetTable = "deployments"
	// TargetInverseTable is the table name for the Target entity.
	// It exists in this package in order to avoid circular dependency with the "target" package.
	TargetInverseTable = "targets"
	// TargetColumn is the table column denoting the target relation/edge.
	TargetColumn = "deployment_target"
)

// Columns holds all SQL columns for deployment fields.
var Columns = []string{
	FieldID,
	FieldOutput,
	FieldError,
	FieldQueuedAt,
	FieldLastModifiedAt,
	FieldStartedAt,
	FieldFinishedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "deployments"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"deployment_config",
	"deployment_target",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultOutput holds the default value on creation for the "output" field.
	DefaultOutput string
	// DefaultError holds the default value on creation for the "error" field.
	DefaultError string
	// DefaultQueuedAt holds the default value on creation for the "queuedAt" field.
	DefaultQueuedAt func() time.Time
	// DefaultLastModifiedAt holds the default value on creation for the "lastModifiedAt" field.
	DefaultLastModifiedAt func() time.Time
)
