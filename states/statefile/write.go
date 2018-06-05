package statefile

import (
	"fmt"
	"io"
)

// WriteState writes the given state to the given writer in the current state
// serialization format.
func WriteState(s *File, w io.Writer) error {
	return fmt.Errorf("WriteState is not yet implemented")
}
