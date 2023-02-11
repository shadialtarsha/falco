// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

var Querystring_get_ArgumentTypes = []value.Type{value.StringType, value.StringType}

func Querystring_get_Validate(args []value.Value) error {
	if len(args) != 2 {
		return errors.ArgumentNotEnough("querystring.get", 2, args)
	}
	for i := range args {
		if args[i].Type() != Querystring_get_ArgumentTypes[i] {
			return errors.TypeMismatch("querystring.get", i+1, Querystring_get_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of querystring.get
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-get/
func Querystring_get(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Querystring_get_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("querystring.get")
}
