// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/internal_redirect/safe_cross_scheme/v3/safe_cross_scheme_config.proto

package envoy_extensions_internal_redirect_safe_cross_scheme_v3

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

// Validate checks the field values on SafeCrossSchemeConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SafeCrossSchemeConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SafeCrossSchemeConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SafeCrossSchemeConfigMultiError, or nil if none found.
func (m *SafeCrossSchemeConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *SafeCrossSchemeConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SafeCrossSchemeConfigMultiError(errors)
	}
	return nil
}

// SafeCrossSchemeConfigMultiError is an error wrapping multiple validation
// errors returned by SafeCrossSchemeConfig.ValidateAll() if the designated
// constraints aren't met.
type SafeCrossSchemeConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SafeCrossSchemeConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SafeCrossSchemeConfigMultiError) AllErrors() []error { return m }

// SafeCrossSchemeConfigValidationError is the validation error returned by
// SafeCrossSchemeConfig.Validate if the designated constraints aren't met.
type SafeCrossSchemeConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SafeCrossSchemeConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SafeCrossSchemeConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SafeCrossSchemeConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SafeCrossSchemeConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SafeCrossSchemeConfigValidationError) ErrorName() string {
	return "SafeCrossSchemeConfigValidationError"
}

// Error satisfies the builtin error interface
func (e SafeCrossSchemeConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSafeCrossSchemeConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SafeCrossSchemeConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SafeCrossSchemeConfigValidationError{}