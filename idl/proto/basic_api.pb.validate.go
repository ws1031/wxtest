// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: basic_api.proto

package proto

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on BasicReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BasicReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BasicReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BasicReqMultiError, or nil
// if none found.
func (m *BasicReq) ValidateAll() error {
	return m.validate(true)
}

func (m *BasicReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return BasicReqMultiError(errors)
	}

	return nil
}

// BasicReqMultiError is an error wrapping multiple validation errors returned
// by BasicReq.ValidateAll() if the designated constraints aren't met.
type BasicReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BasicReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BasicReqMultiError) AllErrors() []error { return m }

// BasicReqValidationError is the validation error returned by
// BasicReq.Validate if the designated constraints aren't met.
type BasicReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BasicReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BasicReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BasicReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BasicReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BasicReqValidationError) ErrorName() string { return "BasicReqValidationError" }

// Error satisfies the builtin error interface
func (e BasicReqValidationError) Error() string {
	errmsg := e.Reason()

	if e.cause != nil {
		errmsg = fmt.Sprintf("%s: %s", errmsg, e.cause.Error())
	}

	return errmsg
}

var _ error = BasicReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BasicReqValidationError{}

// Validate checks the field values on BasicResp with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BasicResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BasicResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BasicRespMultiError, or nil
// if none found.
func (m *BasicResp) ValidateAll() error {
	return m.validate(true)
}

func (m *BasicResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Errno

	// no validation rules for Errmsg

	if len(errors) > 0 {
		return BasicRespMultiError(errors)
	}

	return nil
}

// BasicRespMultiError is an error wrapping multiple validation errors returned
// by BasicResp.ValidateAll() if the designated constraints aren't met.
type BasicRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BasicRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BasicRespMultiError) AllErrors() []error { return m }

// BasicRespValidationError is the validation error returned by
// BasicResp.Validate if the designated constraints aren't met.
type BasicRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BasicRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BasicRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BasicRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BasicRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BasicRespValidationError) ErrorName() string { return "BasicRespValidationError" }

// Error satisfies the builtin error interface
func (e BasicRespValidationError) Error() string {
	errmsg := e.Reason()

	if e.cause != nil {
		errmsg = fmt.Sprintf("%s: %s", errmsg, e.cause.Error())
	}

	return errmsg
}

var _ error = BasicRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BasicRespValidationError{}