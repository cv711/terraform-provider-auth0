// Package value contains helper functions to convert from cty.Value to Go types.
//
// The input value must still conform to the implied type of the given schema,
// or else these functions may produce garbage results or panic. This is usually
// okay because type consistency is enforced when deserializing the value
// returned from the provider over the RPC wire protocol anyway.
package value

import (
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
)

// Bool evaluates the typed value of the value
// and coerces to a pointer of a boolean.
func Bool(rawValue cty.Value) *bool {
	if rawValue.IsNull() {
		return nil
	}

	value := rawValue.True()
	return &value
}

// String evaluates the typed value of the value
// and coerces to a pointer of a string.
func String(rawValue cty.Value) *string {
	if rawValue.IsNull() {
		return nil
	}

	value := rawValue.AsString()
	return &value
}

// Int evaluates the typed value of the value
// and coerces to a pointer of an int.
func Int(rawValue cty.Value) *int {
	if rawValue.IsNull() {
		return nil
	}

	int64Value, _ := rawValue.AsBigFloat().Int64()
	value := int(int64Value)
	return &value
}

// Float64 evaluates the typed value of the value
// and coerces to a pointer of a float64.
func Float64(rawValue cty.Value) *float64 {
	if rawValue.IsNull() {
		return nil
	}

	value, _ := rawValue.AsBigFloat().Float64()
	return &value
}

// Strings evaluates the typed value of the value
// and coerces to a pointer of a slice of strings.
func Strings(rawValues cty.Value) *[]string {
	if rawValues.IsNull() {
		return nil
	}

	value := make([]string, 0)
	for _, rawValue := range rawValues.AsValueSlice() {
		value = append(value, rawValue.AsString())
	}

	return &value
}

// Map evaluates the typed value of the value
// and coerces to a map[string]interface{}.
func Map(rawValue cty.Value) map[string]interface{} {
	if rawValue.IsNull() {
		return nil
	}

	m := make(map[string]interface{})
	for key, value := range rawValue.AsValueMap() {
		m[key] = value.AsString()
	}

	return m
}

// MapOfStrings evaluates the typed value of the value
// and coerces to a pointer of a map of strings.
func MapOfStrings(rawValue cty.Value) *map[string]string {
	if rawValue.IsNull() {
		return nil
	}

	m := make(map[string]string)
	for key, value := range rawValue.AsValueMap() {
		m[key] = value.AsString()
	}

	return &m
}

// MapFromJSON evaluates the typed value of the value
// and coerces to a map[string]interface{}.
func MapFromJSON(rawValue cty.Value) (map[string]interface{}, error) {
	if rawValue.IsNull() {
		return nil, nil
	}

	return structure.ExpandJsonFromString(rawValue.AsString())
}
