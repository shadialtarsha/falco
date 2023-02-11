// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

var Uuid_version3_ArgumentTypes = []value.Type{value.StringType, value.StringType}

func Uuid_version3_Validate(args []value.Value) error {
	if len(args) != 2 {
		return errors.ArgumentNotEnough("uuid.version3", 2, args)
	}
	for i := range args {
		if args[i].Type() != Uuid_version3_ArgumentTypes[i] {
			return errors.TypeMismatch("uuid.version3", i+1, Uuid_version3_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of uuid.version3
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/uuid/uuid-version3/
func Uuid_version3(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Uuid_version3_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("uuid.version3")
}
