// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

var Std_atof_ArgumentTypes = []value.Type{value.StringType}

func Std_atof_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough("std.atof", 1, args)
	}
	for i := range args {
		if args[i].Type() != Std_atof_ArgumentTypes[i] {
			return errors.TypeMismatch("std.atof", i+1, Std_atof_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of std.atof
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-atof/
func Std_atof(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Std_atof_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("std.atof")
}
