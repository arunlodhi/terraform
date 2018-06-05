package states

import (
	"github.com/hashicorp/terraform/addrs"
)

// Resource represents the state of a resource.
type Resource struct {
	// Addr is the module-relative address for the resource this state object
	// belongs to.
	Addr addrs.Resource

	// InstanceKeyType is the key type currently in use for this resource.
	// Within Instances, keys not of this type are considered to be "orphans"
	// that will be destroyed by a subsequent plan.
	InstanceKeyType addrs.InstanceKeyType

	// Instances contains the potentially-multiple instances associated with
	// this resource. This map can contain a mixture of different key types,
	// but only the ones of InstanceKeyType are considered current.
	Instances map[addrs.InstanceKey]*ResourceInstance

	// ProviderConfig is the absolute address for the provider configuration that
	// most recently managed this resource. This is used to connect a resource
	// with a provider configuration when the resource configuration block is
	// not available, such as if it has been removed from configuration
	// altogether.
	ProviderConfig addrs.AbsProviderConfig
}

// ResourceInstance represents the state of a particular instance of a resource.
type ResourceInstance struct {
	// Generations is a container for the potentially-many objects that are
	// either currently represent or formerly represented by this resource
	// instance.
	Generations Generations

	// Dependencies is a set of other addresses in the same module which
	// this instance depended on the last time it was updated. This is used
	// to construct the dependency relationships for an instance whose
	// configuration is no longer available, such as if it has been removed
	// from configuration altogether.
	Dependencies []addrs.Referenceable
}
