package states

import (
	"github.com/zclconf/go-cty/cty"
)

// ResourceInstanceObject is the local representation of a specific remote
// object associated with a resource instance. In practice not all remote
// objects are actually remote in the sense of being accessed over the network,
// but this is the most common case.
type ResourceInstanceObject struct {
	// SchemaVersion identifies which version of the resource type schema the
	// Attrs or AttrsFlat value conforms to. If this is less than the schema
	// version number given by the current provider version then the value
	// must be upgraded to the latest version before use. If it is greater
	// than the current version number then the provider must be upgraded
	// before any operations can be performed.
	SchemaVersion uint64

	// Attrs is the value (of the object type implied by the associated resource
	// type schema) that represents this remote object in Terraform Language
	// expressions and is compared with configuration when producing a diff.
	Attrs cty.Value

	// AttrsFlat is a legacy form of attributes used in older state file
	// formats, and in the new state format for objects that haven't yet been
	// upgraded. This attribute is mutually exclusive with Attrs: for any
	// ResourceInstanceObject, only one of these attributes may be populated
	// and the other must be nil.
	//
	// An instance object with this field populated should be upgraded to use
	// Attrs at the earliest opportunity, since this legacy flatmap-based
	// format will be phased out over time.
	AttrsFlat map[string]string

	// Internal is an opaque value set by the provider when this object was
	// last created or updated. Terraform Core does not use this value in
	// any way and it is not exposed anywhere in the user interface, so
	// a provider can use it for retaining any necessary private state.
	Private cty.Value

	// Status represents the "readiness" of the object as of the last time
	// it was updated.
	Status ObjectStatus
}

// ObjectStatus represents the status of a RemoteObject.
type ObjectStatus rune

//go:generate stringer -type ObjectStatus

const (
	// ObjectReady is an object status for an object that is ready to use.
	ObjectReady ObjectStatus = 'R'

	// ObjectTainted is an object status representing an object that is in
	// an unrecoverable bad state due to a partial failure during a create,
	// update, or delete operation. Since it cannot be moved into the
	// ObjectRead state, a tainted object must be replaced.
	ObjectTainted ObjectStatus = 'T'
)
