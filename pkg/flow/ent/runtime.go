// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/direktiv/direktiv/pkg/flow/ent/cloudevents"
	"github.com/direktiv/direktiv/pkg/flow/ent/events"
	"github.com/direktiv/direktiv/pkg/flow/ent/eventswait"
	"github.com/direktiv/direktiv/pkg/flow/ent/inode"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/instanceruntime"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/ref"
	"github.com/direktiv/direktiv/pkg/flow/ent/revision"
	"github.com/direktiv/direktiv/pkg/flow/ent/route"
	"github.com/direktiv/direktiv/pkg/flow/ent/schema"
	"github.com/direktiv/direktiv/pkg/flow/ent/vardata"
	"github.com/direktiv/direktiv/pkg/flow/ent/varref"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cloudeventsFields := schema.CloudEvents{}.Fields()
	_ = cloudeventsFields
	// cloudeventsDescEventId is the schema descriptor for eventId field.
	cloudeventsDescEventId := cloudeventsFields[1].Descriptor()
	// cloudevents.EventIdValidator is a validator for the "eventId" field. It is called by the builders before save.
	cloudevents.EventIdValidator = cloudeventsDescEventId.Validators[0].(func(string) error)
	// cloudeventsDescFire is the schema descriptor for fire field.
	cloudeventsDescFire := cloudeventsFields[3].Descriptor()
	// cloudevents.DefaultFire holds the default value on creation for the fire field.
	cloudevents.DefaultFire = cloudeventsDescFire.Default.(func() time.Time)
	// cloudeventsDescCreated is the schema descriptor for created field.
	cloudeventsDescCreated := cloudeventsFields[4].Descriptor()
	// cloudevents.DefaultCreated holds the default value on creation for the created field.
	cloudevents.DefaultCreated = cloudeventsDescCreated.Default.(func() time.Time)
	// cloudeventsDescID is the schema descriptor for id field.
	cloudeventsDescID := cloudeventsFields[0].Descriptor()
	// cloudevents.DefaultID holds the default value on creation for the id field.
	cloudevents.DefaultID = cloudeventsDescID.Default.(func() uuid.UUID)
	eventsFields := schema.Events{}.Fields()
	_ = eventsFields
	// eventsDescCreatedAt is the schema descriptor for created_at field.
	eventsDescCreatedAt := eventsFields[5].Descriptor()
	// events.DefaultCreatedAt holds the default value on creation for the created_at field.
	events.DefaultCreatedAt = eventsDescCreatedAt.Default.(func() time.Time)
	// eventsDescUpdatedAt is the schema descriptor for updated_at field.
	eventsDescUpdatedAt := eventsFields[6].Descriptor()
	// events.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	events.DefaultUpdatedAt = eventsDescUpdatedAt.Default.(func() time.Time)
	// events.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	events.UpdateDefaultUpdatedAt = eventsDescUpdatedAt.UpdateDefault.(func() time.Time)
	// eventsDescID is the schema descriptor for id field.
	eventsDescID := eventsFields[0].Descriptor()
	// events.DefaultID holds the default value on creation for the id field.
	events.DefaultID = eventsDescID.Default.(func() uuid.UUID)
	eventswaitFields := schema.EventsWait{}.Fields()
	_ = eventswaitFields
	// eventswaitDescID is the schema descriptor for id field.
	eventswaitDescID := eventswaitFields[0].Descriptor()
	// eventswait.DefaultID holds the default value on creation for the id field.
	eventswait.DefaultID = eventswaitDescID.Default.(func() uuid.UUID)
	inodeFields := schema.Inode{}.Fields()
	_ = inodeFields
	// inodeDescCreatedAt is the schema descriptor for created_at field.
	inodeDescCreatedAt := inodeFields[1].Descriptor()
	// inode.DefaultCreatedAt holds the default value on creation for the created_at field.
	inode.DefaultCreatedAt = inodeDescCreatedAt.Default.(func() time.Time)
	// inodeDescUpdatedAt is the schema descriptor for updated_at field.
	inodeDescUpdatedAt := inodeFields[2].Descriptor()
	// inode.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	inode.DefaultUpdatedAt = inodeDescUpdatedAt.Default.(func() time.Time)
	// inode.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	inode.UpdateDefaultUpdatedAt = inodeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// inodeDescName is the schema descriptor for name field.
	inodeDescName := inodeFields[3].Descriptor()
	// inode.NameValidator is a validator for the "name" field. It is called by the builders before save.
	inode.NameValidator = inodeDescName.Validators[0].(func(string) error)
	// inodeDescID is the schema descriptor for id field.
	inodeDescID := inodeFields[0].Descriptor()
	// inode.DefaultID holds the default value on creation for the id field.
	inode.DefaultID = inodeDescID.Default.(func() uuid.UUID)
	instanceFields := schema.Instance{}.Fields()
	_ = instanceFields
	// instanceDescCreatedAt is the schema descriptor for created_at field.
	instanceDescCreatedAt := instanceFields[1].Descriptor()
	// instance.DefaultCreatedAt holds the default value on creation for the created_at field.
	instance.DefaultCreatedAt = instanceDescCreatedAt.Default.(func() time.Time)
	// instanceDescUpdatedAt is the schema descriptor for updated_at field.
	instanceDescUpdatedAt := instanceFields[2].Descriptor()
	// instance.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	instance.DefaultUpdatedAt = instanceDescUpdatedAt.Default.(func() time.Time)
	// instance.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	instance.UpdateDefaultUpdatedAt = instanceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// instanceDescID is the schema descriptor for id field.
	instanceDescID := instanceFields[0].Descriptor()
	// instance.DefaultID holds the default value on creation for the id field.
	instance.DefaultID = instanceDescID.Default.(func() uuid.UUID)
	instanceruntimeFields := schema.InstanceRuntime{}.Fields()
	_ = instanceruntimeFields
	// instanceruntimeDescID is the schema descriptor for id field.
	instanceruntimeDescID := instanceruntimeFields[0].Descriptor()
	// instanceruntime.DefaultID holds the default value on creation for the id field.
	instanceruntime.DefaultID = instanceruntimeDescID.Default.(func() uuid.UUID)
	logmsgFields := schema.LogMsg{}.Fields()
	_ = logmsgFields
	// logmsgDescID is the schema descriptor for id field.
	logmsgDescID := logmsgFields[0].Descriptor()
	// logmsg.DefaultID holds the default value on creation for the id field.
	logmsg.DefaultID = logmsgDescID.Default.(func() uuid.UUID)
	namespaceFields := schema.Namespace{}.Fields()
	_ = namespaceFields
	// namespaceDescCreatedAt is the schema descriptor for created_at field.
	namespaceDescCreatedAt := namespaceFields[1].Descriptor()
	// namespace.DefaultCreatedAt holds the default value on creation for the created_at field.
	namespace.DefaultCreatedAt = namespaceDescCreatedAt.Default.(func() time.Time)
	// namespaceDescUpdatedAt is the schema descriptor for updated_at field.
	namespaceDescUpdatedAt := namespaceFields[2].Descriptor()
	// namespace.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	namespace.DefaultUpdatedAt = namespaceDescUpdatedAt.Default.(func() time.Time)
	// namespace.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	namespace.UpdateDefaultUpdatedAt = namespaceDescUpdatedAt.UpdateDefault.(func() time.Time)
	// namespaceDescConfig is the schema descriptor for config field.
	namespaceDescConfig := namespaceFields[3].Descriptor()
	// namespace.DefaultConfig holds the default value on creation for the config field.
	namespace.DefaultConfig = namespaceDescConfig.Default.(string)
	// namespaceDescName is the schema descriptor for name field.
	namespaceDescName := namespaceFields[4].Descriptor()
	// namespace.NameValidator is a validator for the "name" field. It is called by the builders before save.
	namespace.NameValidator = namespaceDescName.Validators[0].(func(string) error)
	// namespaceDescID is the schema descriptor for id field.
	namespaceDescID := namespaceFields[0].Descriptor()
	// namespace.DefaultID holds the default value on creation for the id field.
	namespace.DefaultID = namespaceDescID.Default.(func() uuid.UUID)
	refFields := schema.Ref{}.Fields()
	_ = refFields
	// refDescImmutable is the schema descriptor for immutable field.
	refDescImmutable := refFields[1].Descriptor()
	// ref.DefaultImmutable holds the default value on creation for the immutable field.
	ref.DefaultImmutable = refDescImmutable.Default.(bool)
	// refDescName is the schema descriptor for name field.
	refDescName := refFields[2].Descriptor()
	// ref.NameValidator is a validator for the "name" field. It is called by the builders before save.
	ref.NameValidator = refDescName.Validators[0].(func(string) error)
	// refDescID is the schema descriptor for id field.
	refDescID := refFields[0].Descriptor()
	// ref.DefaultID holds the default value on creation for the id field.
	ref.DefaultID = refDescID.Default.(func() uuid.UUID)
	revisionFields := schema.Revision{}.Fields()
	_ = revisionFields
	// revisionDescCreatedAt is the schema descriptor for created_at field.
	revisionDescCreatedAt := revisionFields[1].Descriptor()
	// revision.DefaultCreatedAt holds the default value on creation for the created_at field.
	revision.DefaultCreatedAt = revisionDescCreatedAt.Default.(func() time.Time)
	// revisionDescID is the schema descriptor for id field.
	revisionDescID := revisionFields[0].Descriptor()
	// revision.DefaultID holds the default value on creation for the id field.
	revision.DefaultID = revisionDescID.Default.(func() uuid.UUID)
	routeFields := schema.Route{}.Fields()
	_ = routeFields
	// routeDescID is the schema descriptor for id field.
	routeDescID := routeFields[0].Descriptor()
	// route.DefaultID holds the default value on creation for the id field.
	route.DefaultID = routeDescID.Default.(func() uuid.UUID)
	vardataFields := schema.VarData{}.Fields()
	_ = vardataFields
	// vardataDescCreatedAt is the schema descriptor for created_at field.
	vardataDescCreatedAt := vardataFields[1].Descriptor()
	// vardata.DefaultCreatedAt holds the default value on creation for the created_at field.
	vardata.DefaultCreatedAt = vardataDescCreatedAt.Default.(func() time.Time)
	// vardataDescUpdatedAt is the schema descriptor for updated_at field.
	vardataDescUpdatedAt := vardataFields[2].Descriptor()
	// vardata.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	vardata.DefaultUpdatedAt = vardataDescUpdatedAt.Default.(func() time.Time)
	// vardata.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	vardata.UpdateDefaultUpdatedAt = vardataDescUpdatedAt.UpdateDefault.(func() time.Time)
	// vardataDescMimeType is the schema descriptor for mime_type field.
	vardataDescMimeType := vardataFields[6].Descriptor()
	// vardata.DefaultMimeType holds the default value on creation for the mime_type field.
	vardata.DefaultMimeType = vardataDescMimeType.Default.(string)
	// vardataDescID is the schema descriptor for id field.
	vardataDescID := vardataFields[0].Descriptor()
	// vardata.DefaultID holds the default value on creation for the id field.
	vardata.DefaultID = vardataDescID.Default.(func() uuid.UUID)
	varrefFields := schema.VarRef{}.Fields()
	_ = varrefFields
	// varrefDescName is the schema descriptor for name field.
	varrefDescName := varrefFields[1].Descriptor()
	// varref.NameValidator is a validator for the "name" field. It is called by the builders before save.
	varref.NameValidator = varrefDescName.Validators[0].(func(string) error)
	// varrefDescID is the schema descriptor for id field.
	varrefDescID := varrefFields[0].Descriptor()
	// varref.DefaultID holds the default value on creation for the id field.
	varref.DefaultID = varrefDescID.Default.(func() uuid.UUID)
	workflowFields := schema.Workflow{}.Fields()
	_ = workflowFields
	// workflowDescLive is the schema descriptor for live field.
	workflowDescLive := workflowFields[1].Descriptor()
	// workflow.DefaultLive holds the default value on creation for the live field.
	workflow.DefaultLive = workflowDescLive.Default.(bool)
	// workflowDescID is the schema descriptor for id field.
	workflowDescID := workflowFields[0].Descriptor()
	// workflow.DefaultID holds the default value on creation for the id field.
	workflow.DefaultID = workflowDescID.Default.(func() uuid.UUID)
}
