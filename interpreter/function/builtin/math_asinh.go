// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

var Math_asinh_ArgumentTypes = []value.Type{value.FloatType}

func Math_asinh_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough("math.asinh", 1, args)
	}
	for i := range args {
		if args[i].Type() != Math_asinh_ArgumentTypes[i] {
			return errors.TypeMismatch("math.asinh", i+1, Math_asinh_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of math.asinh
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-trig/math-asinh/
func Math_asinh(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Math_asinh_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("math.asinh")
}
