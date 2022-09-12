package values

import (
	"testing"

	"github.com/hashicorp/go-cty/cty"
)

func TestBool(t *testing.T) {

	v := Bool(cty.NullVal(cty.Bool))
	if v != nil {
		t.Errorf("Expected to be null, got %t", *v)
	}

	v = Bool(cty.BoolVal(true))
	if *v != true {
		t.Errorf("expected to be true, got %t", *v)
	}

	v = Bool(cty.BoolVal(false))
	if *v != false {
		t.Errorf("expected to be false, got %t", *v)
	}
}
