// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

var Math_acos_ArgumentTypes = []value.Type{value.FloatType}

func Math_acos_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough("math.acos", 1, args)
	}
	for i := range args {
		if args[i].Type() != Math_acos_ArgumentTypes[i] {
			return errors.TypeMismatch("math.acos", i+1, Math_acos_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of math.acos
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-acos/
func Math_acos(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Math_acos_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("math.acos")
}
