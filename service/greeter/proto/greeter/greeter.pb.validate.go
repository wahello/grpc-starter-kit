// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service/greeter/proto/greeter/greeter.proto

package greeter

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _greeter_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on HelloRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HelloRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// HelloRequestValidationError is the validation error returned by
// HelloRequest.Validate if the designated constraints aren't met.
type HelloRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloRequestValidationError) ErrorName() string { return "HelloRequestValidationError" }

// Error satisfies the builtin error interface
func (e HelloRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloRequestValidationError{}

// Validate checks the field values on HelloResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HelloResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Msg

	return nil
}

// HelloResponseValidationError is the validation error returned by
// HelloResponse.Validate if the designated constraints aren't met.
type HelloResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloResponseValidationError) ErrorName() string { return "HelloResponseValidationError" }

// Error satisfies the builtin error interface
func (e HelloResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloResponseValidationError{}
