// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of digest.hmac_sha256
// Arguments may be:
// - STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-hmac-sha256/
func Test_Digest_hmac_sha256(t *testing.T) {
	ret, err := Digest_hmac_sha256(
		&context.Context{},
		&value.String{Value: "key"},
		&value.String{Value: "input"},
	)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if ret.Type() != value.StringType {
		t.Errorf("Unexpected return type, expect=STRING, got=%s", ret.Type())
	}
	v := value.Unwrap[*value.String](ret)
	expect := "9e089ec13af881a8ac227a736c3e7c490ea3b4afca0c5f83dff6393b683a72e3"
	if v.Value != expect {
		t.Errorf("return value unmach, expect=%s, got=%s", expect, v.Value)
	}
}
