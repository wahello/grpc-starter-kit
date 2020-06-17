// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: mkit/service/account/user/v1/user_service.proto

package userv1

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _user_service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ExistRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ExistRequest) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetId(); wrapper != nil {

		if err := m._validateUuid(wrapper.GetValue()); err != nil {
			return ExistRequestValidationError{
				field:  "Id",
				reason: "value must be a valid UUID",
				cause:  err,
			}
		}

	}

	if wrapper := m.GetUsername(); wrapper != nil {

		if l := utf8.RuneCountInString(wrapper.GetValue()); l < 4 || l > 16 {
			return ExistRequestValidationError{
				field:  "Username",
				reason: "value length must be between 4 and 16 runes, inclusive",
			}
		}

		if len(wrapper.GetValue()) > 256 {
			return ExistRequestValidationError{
				field:  "Username",
				reason: "value length must be at most 256 bytes",
			}
		}

		if !_ExistRequest_Username_Pattern.MatchString(wrapper.GetValue()) {
			return ExistRequestValidationError{
				field:  "Username",
				reason: "value does not match regex pattern \"^[a-z0-9_-]{3,15}$\"",
			}
		}

	}

	if wrapper := m.GetFirstName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return ExistRequestValidationError{
				field:  "FirstName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetLastName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return ExistRequestValidationError{
				field:  "LastName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetEmail(); wrapper != nil {

		if err := m._validateEmail(wrapper.GetValue()); err != nil {
			return ExistRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
		}

	}

	return nil
}

func (m *ExistRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *ExistRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

func (m *ExistRequest) _validateUuid(uuid string) error {
	if matched := _user_service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// ExistRequestValidationError is the validation error returned by
// ExistRequest.Validate if the designated constraints aren't met.
type ExistRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExistRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExistRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExistRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExistRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExistRequestValidationError) ErrorName() string { return "ExistRequestValidationError" }

// Error satisfies the builtin error interface
func (e ExistRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExistRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExistRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExistRequestValidationError{}

var _ExistRequest_Username_Pattern = regexp.MustCompile("^[a-z0-9_-]{3,15}$")

// Validate is disabled for ExistResponse. This method will always return nil.
func (m *ExistResponse) Validate() error {
	return nil
}

// ExistResponseValidationError is the validation error returned by
// ExistResponse.Validate if the designated constraints aren't met.
type ExistResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExistResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExistResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExistResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExistResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExistResponseValidationError) ErrorName() string { return "ExistResponseValidationError" }

// Error satisfies the builtin error interface
func (e ExistResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExistResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExistResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExistResponseValidationError{}

// Validate checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListRequest) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetLimit(); wrapper != nil {

		if val := wrapper.GetValue(); val < 1 || val > 100 {
			return ListRequestValidationError{
				field:  "Limit",
				reason: "value must be inside range [1, 100]",
			}
		}

	}

	if wrapper := m.GetPage(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			return ListRequestValidationError{
				field:  "Page",
				reason: "value must be greater than or equal to 1",
			}
		}

	}

	{
		tmp := m.GetSort()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListRequestValidationError{
					field:  "Sort",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if wrapper := m.GetUsername(); wrapper != nil {

		if l := utf8.RuneCountInString(wrapper.GetValue()); l < 4 || l > 16 {
			return ListRequestValidationError{
				field:  "Username",
				reason: "value length must be between 4 and 16 runes, inclusive",
			}
		}

		if len(wrapper.GetValue()) > 256 {
			return ListRequestValidationError{
				field:  "Username",
				reason: "value length must be at most 256 bytes",
			}
		}

		if !_ListRequest_Username_Pattern.MatchString(wrapper.GetValue()) {
			return ListRequestValidationError{
				field:  "Username",
				reason: "value does not match regex pattern \"^[a-z0-9_-]{3,15}$\"",
			}
		}

	}

	if wrapper := m.GetFirstName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return ListRequestValidationError{
				field:  "FirstName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetLastName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return ListRequestValidationError{
				field:  "LastName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetEmail(); wrapper != nil {

		if err := m._validateEmail(wrapper.GetValue()); err != nil {
			return ListRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
		}

	}

	return nil
}

func (m *ListRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *ListRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// ListRequestValidationError is the validation error returned by
// ListRequest.Validate if the designated constraints aren't met.
type ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestValidationError) ErrorName() string { return "ListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestValidationError{}

var _ListRequest_Username_Pattern = regexp.MustCompile("^[a-z0-9_-]{3,15}$")

// Validate is disabled for ListResponse. This method will always return nil.
func (m *ListResponse) Validate() error {
	return nil
}

// ListResponseValidationError is the validation error returned by
// ListResponse.Validate if the designated constraints aren't met.
type ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResponseValidationError) ErrorName() string { return "ListResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResponseValidationError{}

// Validate checks the field values on GetRequest with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *GetRequest) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetId(); wrapper != nil {

		if err := m._validateUuid(wrapper.GetValue()); err != nil {
			return GetRequestValidationError{
				field:  "Id",
				reason: "value must be a valid UUID",
				cause:  err,
			}
		}

	}

	if wrapper := m.GetUsername(); wrapper != nil {

		if l := utf8.RuneCountInString(wrapper.GetValue()); l < 4 || l > 16 {
			return GetRequestValidationError{
				field:  "Username",
				reason: "value length must be between 4 and 16 runes, inclusive",
			}
		}

		if len(wrapper.GetValue()) > 256 {
			return GetRequestValidationError{
				field:  "Username",
				reason: "value length must be at most 256 bytes",
			}
		}

		if !_GetRequest_Username_Pattern.MatchString(wrapper.GetValue()) {
			return GetRequestValidationError{
				field:  "Username",
				reason: "value does not match regex pattern \"^[a-z0-9_-]{3,15}$\"",
			}
		}

	}

	if wrapper := m.GetFirstName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return GetRequestValidationError{
				field:  "FirstName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetLastName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return GetRequestValidationError{
				field:  "LastName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetEmail(); wrapper != nil {

		if err := m._validateEmail(wrapper.GetValue()); err != nil {
			return GetRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
		}

	}

	return nil
}

func (m *GetRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *GetRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

func (m *GetRequest) _validateUuid(uuid string) error {
	if matched := _user_service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetRequestValidationError is the validation error returned by
// GetRequest.Validate if the designated constraints aren't met.
type GetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRequestValidationError) ErrorName() string { return "GetRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRequestValidationError{}

var _GetRequest_Username_Pattern = regexp.MustCompile("^[a-z0-9_-]{3,15}$")

// Validate is disabled for GetResponse. This method will always return nil.
func (m *GetResponse) Validate() error {
	return nil
}

// GetResponseValidationError is the validation error returned by
// GetResponse.Validate if the designated constraints aren't met.
type GetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResponseValidationError) ErrorName() string { return "GetResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResponseValidationError{}

// Validate checks the field values on CreateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CreateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetUsername(); wrapper != nil {

		if l := utf8.RuneCountInString(wrapper.GetValue()); l < 4 || l > 16 {
			return CreateRequestValidationError{
				field:  "Username",
				reason: "value length must be between 4 and 16 runes, inclusive",
			}
		}

		if len(wrapper.GetValue()) > 256 {
			return CreateRequestValidationError{
				field:  "Username",
				reason: "value length must be at most 256 bytes",
			}
		}

		if !_CreateRequest_Username_Pattern.MatchString(wrapper.GetValue()) {
			return CreateRequestValidationError{
				field:  "Username",
				reason: "value does not match regex pattern \"^[a-z0-9_-]{3,15}$\"",
			}
		}

	}

	if wrapper := m.GetFirstName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return CreateRequestValidationError{
				field:  "FirstName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetLastName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return CreateRequestValidationError{
				field:  "LastName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetEmail(); wrapper != nil {

		if err := m._validateEmail(wrapper.GetValue()); err != nil {
			return CreateRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
		}

	}

	return nil
}

func (m *CreateRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *CreateRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// CreateRequestValidationError is the validation error returned by
// CreateRequest.Validate if the designated constraints aren't met.
type CreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRequestValidationError) ErrorName() string { return "CreateRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRequestValidationError{}

var _CreateRequest_Username_Pattern = regexp.MustCompile("^[a-z0-9_-]{3,15}$")

// Validate is disabled for CreateResponse. This method will always return nil.
func (m *CreateResponse) Validate() error {
	return nil
}

// CreateResponseValidationError is the validation error returned by
// CreateResponse.Validate if the designated constraints aren't met.
type CreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResponseValidationError) ErrorName() string { return "CreateResponseValidationError" }

// Error satisfies the builtin error interface
func (e CreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResponseValidationError{}

// Validate checks the field values on UpdateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UpdateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetId(); wrapper != nil {

		if err := m._validateUuid(wrapper.GetValue()); err != nil {
			return UpdateRequestValidationError{
				field:  "Id",
				reason: "value must be a valid UUID",
				cause:  err,
			}
		}

	}

	if wrapper := m.GetUsername(); wrapper != nil {

		if l := utf8.RuneCountInString(wrapper.GetValue()); l < 4 || l > 16 {
			return UpdateRequestValidationError{
				field:  "Username",
				reason: "value length must be between 4 and 16 runes, inclusive",
			}
		}

		if len(wrapper.GetValue()) > 256 {
			return UpdateRequestValidationError{
				field:  "Username",
				reason: "value length must be at most 256 bytes",
			}
		}

		if !_UpdateRequest_Username_Pattern.MatchString(wrapper.GetValue()) {
			return UpdateRequestValidationError{
				field:  "Username",
				reason: "value does not match regex pattern \"^[a-z0-9_-]{3,15}$\"",
			}
		}

	}

	if wrapper := m.GetFirstName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return UpdateRequestValidationError{
				field:  "FirstName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetLastName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return UpdateRequestValidationError{
				field:  "LastName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetEmail(); wrapper != nil {

		if err := m._validateEmail(wrapper.GetValue()); err != nil {
			return UpdateRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
		}

	}

	return nil
}

func (m *UpdateRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *UpdateRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

func (m *UpdateRequest) _validateUuid(uuid string) error {
	if matched := _user_service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// UpdateRequestValidationError is the validation error returned by
// UpdateRequest.Validate if the designated constraints aren't met.
type UpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRequestValidationError) ErrorName() string { return "UpdateRequestValidationError" }

// Error satisfies the builtin error interface
func (e UpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRequestValidationError{}

var _UpdateRequest_Username_Pattern = regexp.MustCompile("^[a-z0-9_-]{3,15}$")

// Validate is disabled for UpdateResponse. This method will always return nil.
func (m *UpdateResponse) Validate() error {
	return nil
}

// UpdateResponseValidationError is the validation error returned by
// UpdateResponse.Validate if the designated constraints aren't met.
type UpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateResponseValidationError) ErrorName() string { return "UpdateResponseValidationError" }

// Error satisfies the builtin error interface
func (e UpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateResponseValidationError{}

// Validate checks the field values on DeleteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeleteRequest) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetId(); wrapper != nil {

		if err := m._validateUuid(wrapper.GetValue()); err != nil {
			return DeleteRequestValidationError{
				field:  "Id",
				reason: "value must be a valid UUID",
				cause:  err,
			}
		}

	}

	if wrapper := m.GetUsername(); wrapper != nil {

		if l := utf8.RuneCountInString(wrapper.GetValue()); l < 4 || l > 16 {
			return DeleteRequestValidationError{
				field:  "Username",
				reason: "value length must be between 4 and 16 runes, inclusive",
			}
		}

		if len(wrapper.GetValue()) > 256 {
			return DeleteRequestValidationError{
				field:  "Username",
				reason: "value length must be at most 256 bytes",
			}
		}

		if !_DeleteRequest_Username_Pattern.MatchString(wrapper.GetValue()) {
			return DeleteRequestValidationError{
				field:  "Username",
				reason: "value does not match regex pattern \"^[a-z0-9_-]{3,15}$\"",
			}
		}

	}

	if wrapper := m.GetFirstName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return DeleteRequestValidationError{
				field:  "FirstName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetLastName(); wrapper != nil {

		if utf8.RuneCountInString(wrapper.GetValue()) < 3 {
			return DeleteRequestValidationError{
				field:  "LastName",
				reason: "value length must be at least 3 runes",
			}
		}

	}

	if wrapper := m.GetEmail(); wrapper != nil {

		if err := m._validateEmail(wrapper.GetValue()); err != nil {
			return DeleteRequestValidationError{
				field:  "Email",
				reason: "value must be a valid email address",
				cause:  err,
			}
		}

	}

	return nil
}

func (m *DeleteRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *DeleteRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

func (m *DeleteRequest) _validateUuid(uuid string) error {
	if matched := _user_service_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// DeleteRequestValidationError is the validation error returned by
// DeleteRequest.Validate if the designated constraints aren't met.
type DeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRequestValidationError) ErrorName() string { return "DeleteRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRequestValidationError{}

var _DeleteRequest_Username_Pattern = regexp.MustCompile("^[a-z0-9_-]{3,15}$")

// Validate is disabled for DeleteResponse. This method will always return nil.
func (m *DeleteResponse) Validate() error {
	return nil
}

// DeleteResponseValidationError is the validation error returned by
// DeleteResponse.Validate if the designated constraints aren't met.
type DeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResponseValidationError) ErrorName() string { return "DeleteResponseValidationError" }

// Error satisfies the builtin error interface
func (e DeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResponseValidationError{}
