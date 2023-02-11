// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

var Strftime_ArgumentTypes = []value.Type{value.StringType, value.TimeType}

func Strftime_Validate(args []value.Value) error {
	if len(args) != 2 {
		return errors.ArgumentNotEnough("strftime", 2, args)
	}
	for i := range args {
		if args[i].Type() != Strftime_ArgumentTypes[i] {
			return errors.TypeMismatch("strftime", i+1, Strftime_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of strftime
// Arguments may be:
// - STRING, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/strftime/
func Strftime(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Strftime_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("strftime")
}
