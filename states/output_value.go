package states

import (
	"github.com/zclconf/go-cty/cty"
)

// OutputValue represents the state of a particular output value.
type OutputValue struct {
	Value     cty.Value
	Sensitive bool
}
