package values

import (
	"github.com/hashicorp/go-cty/cty"
)

// Bool accesses the value held by the bool flag key
// and type asserts it as a pointer to a bool.
func Bool(rawValue cty.Value) *bool {

	if rawValue.IsNull() {
		return nil
	}

	r := rawValue.True()
	return &r
}
