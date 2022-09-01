package value

import (
	"log"

	"github.com/auth0/go-auth0"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/go-cty/cty/gocty"
)

// Bool TODO description
func Bool(field cty.Value) *bool {
	// add check that field is of type bool before maybe?

	if field.IsNull() {
		return nil
	}

	return auth0.Bool(field.True())
}

// String TODO description
func String(field cty.Value) *string {
	// add check that field is of type string before maybe?
	if field.IsNull() {
		return nil
	}

	// you can probably write the bool and string funcs to be consistent with float meaning using the gocty.FromCtyValue

	return auth0.String(field.AsString())
}

// Float64 TODO description
// Useful reading material https://github.com/hashicorp/go-cty/blob/master/docs/gocty.md
func Float64(field cty.Value) *float64 {
	if field.IsNull() {
		return nil
	}

	var value float64
	err := gocty.FromCtyValue(field, &value)
	if err != nil {
		log.Printf("error bla bla")
		// return nil or default value for float?
		// or even ignore the error completely as
		// we know for sure that this field is of that exact type?
		// cuz of the schema validation?
		// what's certain is that we don't want to return the go type
		// and also an error to make the code look like
		// it was written from the pits of hell
		return nil
	}

	return &value
}
