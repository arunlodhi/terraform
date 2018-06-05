package states

import "fmt"

// Generations is a container type representing the remote objects for the
// current and any deposed generations of a particular resource instance.
type Generations struct {
	// Current, if non-nil, is the remote object that is currently represented
	// by the corresponding resource instance.
	Current *ResourceInstanceObject

	// Deposed, if len > 0, contains any remote objects that were previously
	// represented by the corresponding resource instance but have been
	// replaced and are pending destruction due to the create_before_destroy
	// lifecycle mode.
	Deposed []*ResourceInstanceObject
}

func (g Generations) Get(gen Generation) *ResourceInstanceObject {
	if gen == CurrentGen {
		return g.Current
	}
	if dg, ok := gen.(DeposedGen); ok {
		return g.Deposed[int(dg)]
	}
	if gen == nil {
		panic(fmt.Sprintf("get with nil Generation"))
	}
	// Should never fall out here, since the above covers all possible
	// Generation values.
	panic(fmt.Sprintf("get invalid Generation %#v", gen))
}

// Generation is used to represent multiple objects in a succession of objects
// represented by a single resource instance address. A resource instance can
// have multiple generations over its lifetime due to object replacement
// (when a change can't be applied without destroying and re-creating), and
// multiple generations can exist at the same time when create_before_destroy
// is used.
//
// A Generation value can either be the value of the variable "CurrentGen" or
// a value of type DeposedGen. Generation values can be compared for equality
// using "==" and used as map keys. The zero value of Generation (nil) is not
// a valid generation and must not be used.
type Generation interface {
	generation()
}

// CurrentGen is the Generation representing the currently-active object for
// a resource instance.
var CurrentGen Generation

type currentGen struct{}

func (g currentGen) generation() {}

// DeposedGen is a Generation type representing deposed objects. An object is
// deposed during its replacement when create_before_destroy is in use, so
// that Terraform can retain the id of the prior object in order to destroy
// it after the new primary object is ready.
//
// Create a Deposed generation by converting from int:
//
//     DeposedGen(0)
//
// The value is an index into the Generations.Tainted slice for a particular
// resource instance.
type DeposedGen int

var _ Generation = DeposedGen(0)

func (g DeposedGen) generation() {}

func init() {
	CurrentGen = currentGen{}
}
