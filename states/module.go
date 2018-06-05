package states

import (
	"github.com/hashicorp/terraform/addrs"
)

// Module is a container for the states of objects within a particular module.
type Module struct {
	Addr addrs.ResourceInstance

	Resources    map[string]*Resource
	OutputValues map[string]*OutputValue
}
