// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

var Ratelimit_ratecounter_increment_ArgumentTypes = []value.Type{value.IdentType, value.StringType, value.IntegerType}

func Ratelimit_ratecounter_increment_Validate(args []value.Value) error {
	if len(args) != 3 {
		return errors.ArgumentNotEnough("ratelimit.ratecounter_increment", 3, args)
	}
	for i := range args {
		if args[i].Type() != Ratelimit_ratecounter_increment_ArgumentTypes[i] {
			return errors.TypeMismatch("ratelimit.ratecounter_increment", i+1, Ratelimit_ratecounter_increment_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of ratelimit.ratecounter_increment
// Arguments may be:
// - ID, STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/rate-limiting/ratelimit-ratecounter-increment/
func Ratelimit_ratecounter_increment(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Ratelimit_ratecounter_increment_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("ratelimit.ratecounter_increment")
}
