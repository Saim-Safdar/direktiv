// Code generated by entc, DO NOT EDIT.

package workflowevents

const (
	// Label holds the string label denoting the workflowevents type in the database.
	Label = "workflow_events"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEvents holds the string denoting the events field in the database.
	FieldEvents = "events"
	// FieldCorrelations holds the string denoting the correlations field in the database.
	FieldCorrelations = "correlations"
	// FieldSignature holds the string denoting the signature field in the database.
	FieldSignature = "signature"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// EdgeWorkflow holds the string denoting the workflow edge name in mutations.
	EdgeWorkflow = "workflow"
	// EdgeWfeventswait holds the string denoting the wfeventswait edge name in mutations.
	EdgeWfeventswait = "wfeventswait"
	// EdgeWorkflowinstance holds the string denoting the workflowinstance edge name in mutations.
	EdgeWorkflowinstance = "workflowinstance"
	// Table holds the table name of the workflowevents in the database.
	Table = "workflow_events"
	// WorkflowTable is the table that holds the workflow relation/edge.
	WorkflowTable = "workflow_events"
	// WorkflowInverseTable is the table name for the Workflow entity.
	// It exists in this package in order to avoid circular dependency with the "workflow" package.
	WorkflowInverseTable = "workflows"
	// WorkflowColumn is the table column denoting the workflow relation/edge.
	WorkflowColumn = "workflow_wfevents"
	// WfeventswaitTable is the table that holds the wfeventswait relation/edge.
	WfeventswaitTable = "workflow_events_waits"
	// WfeventswaitInverseTable is the table name for the WorkflowEventsWait entity.
	// It exists in this package in order to avoid circular dependency with the "workfloweventswait" package.
	WfeventswaitInverseTable = "workflow_events_waits"
	// WfeventswaitColumn is the table column denoting the wfeventswait relation/edge.
	WfeventswaitColumn = "workflow_events_wfeventswait"
	// WorkflowinstanceTable is the table that holds the workflowinstance relation/edge.
	WorkflowinstanceTable = "workflow_events"
	// WorkflowinstanceInverseTable is the table name for the WorkflowInstance entity.
	// It exists in this package in order to avoid circular dependency with the "workflowinstance" package.
	WorkflowinstanceInverseTable = "workflow_instances"
	// WorkflowinstanceColumn is the table column denoting the workflowinstance relation/edge.
	WorkflowinstanceColumn = "workflow_instance_instance"
)

// Columns holds all SQL columns for workflowevents fields.
var Columns = []string{
	FieldID,
	FieldEvents,
	FieldCorrelations,
	FieldSignature,
	FieldCount,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "workflow_events"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"workflow_wfevents",
	"workflow_instance_instance",
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
