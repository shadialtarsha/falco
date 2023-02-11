// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

var Addr_is_ipv6_ArgumentTypes = []value.Type{value.IpType}

func Addr_is_ipv6_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough("addr.is_ipv6", 1, args)
	}
	for i := range args {
		if args[i].Type() != Addr_is_ipv6_ArgumentTypes[i] {
			return errors.TypeMismatch("addr.is_ipv6", i+1, Addr_is_ipv6_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of addr.is_ipv6
// Arguments may be:
// - IP
// Reference: https://developer.fastly.com/reference/vcl/functions/miscellaneous/addr-is-ipv6/
func Addr_is_ipv6(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Addr_is_ipv6_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("addr.is_ipv6")
}
