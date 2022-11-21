// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/group/v1/group.proto

package v1

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

// Validate checks the field values on CreateGroupRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateGroupRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// CreateGroupRequestValidationError is the validation error returned by
// CreateGroupRequest.Validate if the designated constraints aren't met.
type CreateGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateGroupRequestValidationError) ErrorName() string {
	return "CreateGroupRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateGroupRequestValidationError{}

// Validate checks the field values on CreateGroupReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateGroupReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CreateGroupReplyValidationError is the validation error returned by
// CreateGroupReply.Validate if the designated constraints aren't met.
type CreateGroupReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateGroupReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateGroupReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateGroupReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateGroupReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateGroupReplyValidationError) ErrorName() string { return "CreateGroupReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateGroupReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateGroupReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateGroupReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateGroupReplyValidationError{}

// Validate checks the field values on UpdateGroupRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateGroupRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetGroup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateGroupRequestValidationError{
				field:  "Group",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateGroupRequestValidationError is the validation error returned by
// UpdateGroupRequest.Validate if the designated constraints aren't met.
type UpdateGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateGroupRequestValidationError) ErrorName() string {
	return "UpdateGroupRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateGroupRequestValidationError{}

// Validate checks the field values on UpdateGroupReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdateGroupReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateGroupReplyValidationError is the validation error returned by
// UpdateGroupReply.Validate if the designated constraints aren't met.
type UpdateGroupReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateGroupReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateGroupReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateGroupReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateGroupReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateGroupReplyValidationError) ErrorName() string { return "UpdateGroupReplyValidationError" }

// Error satisfies the builtin error interface
func (e UpdateGroupReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateGroupReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateGroupReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateGroupReplyValidationError{}

// Validate checks the field values on DeleteGroupRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteGroupRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// DeleteGroupRequestValidationError is the validation error returned by
// DeleteGroupRequest.Validate if the designated constraints aren't met.
type DeleteGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteGroupRequestValidationError) ErrorName() string {
	return "DeleteGroupRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteGroupRequestValidationError{}

// Validate checks the field values on DeleteGroupReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeleteGroupReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteGroupReplyValidationError is the validation error returned by
// DeleteGroupReply.Validate if the designated constraints aren't met.
type DeleteGroupReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteGroupReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteGroupReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteGroupReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteGroupReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteGroupReplyValidationError) ErrorName() string { return "DeleteGroupReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteGroupReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteGroupReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteGroupReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteGroupReplyValidationError{}

// Validate checks the field values on GetGroupRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetGroupRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// GetGroupRequestValidationError is the validation error returned by
// GetGroupRequest.Validate if the designated constraints aren't met.
type GetGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGroupRequestValidationError) ErrorName() string { return "GetGroupRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGroupRequestValidationError{}

// Validate checks the field values on GetGroupReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetGroupReply) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetGroup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetGroupReplyValidationError{
				field:  "Group",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetGroupReplyValidationError is the validation error returned by
// GetGroupReply.Validate if the designated constraints aren't met.
type GetGroupReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGroupReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGroupReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGroupReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGroupReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGroupReplyValidationError) ErrorName() string { return "GetGroupReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetGroupReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGroupReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGroupReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGroupReplyValidationError{}

// Validate checks the field values on ListGroupRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListGroupRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListGroupRequestValidationError is the validation error returned by
// ListGroupRequest.Validate if the designated constraints aren't met.
type ListGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListGroupRequestValidationError) ErrorName() string { return "ListGroupRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListGroupRequestValidationError{}

// Validate checks the field values on ListGroupReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListGroupReply) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetGroups() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListGroupReplyValidationError{
					field:  fmt.Sprintf("Groups[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListGroupReplyValidationError is the validation error returned by
// ListGroupReply.Validate if the designated constraints aren't met.
type ListGroupReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListGroupReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListGroupReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListGroupReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListGroupReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListGroupReplyValidationError) ErrorName() string { return "ListGroupReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListGroupReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListGroupReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListGroupReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListGroupReplyValidationError{}

// Validate checks the field values on GetUsersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetUsersRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Group

	return nil
}

// GetUsersRequestValidationError is the validation error returned by
// GetUsersRequest.Validate if the designated constraints aren't met.
type GetUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersRequestValidationError) ErrorName() string { return "GetUsersRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersRequestValidationError{}

// Validate checks the field values on GetUsersReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetUsersReply) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetUsersReplyValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetUsersReplyValidationError is the validation error returned by
// GetUsersReply.Validate if the designated constraints aren't met.
type GetUsersReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersReplyValidationError) ErrorName() string { return "GetUsersReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetUsersReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersReplyValidationError{}

// Validate checks the field values on UpdateGroupRequest_Group with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateGroupRequest_Group) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	return nil
}

// UpdateGroupRequest_GroupValidationError is the validation error returned by
// UpdateGroupRequest_Group.Validate if the designated constraints aren't met.
type UpdateGroupRequest_GroupValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateGroupRequest_GroupValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateGroupRequest_GroupValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateGroupRequest_GroupValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateGroupRequest_GroupValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateGroupRequest_GroupValidationError) ErrorName() string {
	return "UpdateGroupRequest_GroupValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateGroupRequest_GroupValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateGroupRequest_Group.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateGroupRequest_GroupValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateGroupRequest_GroupValidationError{}

// Validate checks the field values on GetGroupReply_Group with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetGroupReply_Group) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	return nil
}

// GetGroupReply_GroupValidationError is the validation error returned by
// GetGroupReply_Group.Validate if the designated constraints aren't met.
type GetGroupReply_GroupValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetGroupReply_GroupValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetGroupReply_GroupValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetGroupReply_GroupValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetGroupReply_GroupValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetGroupReply_GroupValidationError) ErrorName() string {
	return "GetGroupReply_GroupValidationError"
}

// Error satisfies the builtin error interface
func (e GetGroupReply_GroupValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetGroupReply_Group.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetGroupReply_GroupValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetGroupReply_GroupValidationError{}

// Validate checks the field values on ListGroupReply_Group with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListGroupReply_Group) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	return nil
}

// ListGroupReply_GroupValidationError is the validation error returned by
// ListGroupReply_Group.Validate if the designated constraints aren't met.
type ListGroupReply_GroupValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListGroupReply_GroupValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListGroupReply_GroupValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListGroupReply_GroupValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListGroupReply_GroupValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListGroupReply_GroupValidationError) ErrorName() string {
	return "ListGroupReply_GroupValidationError"
}

// Error satisfies the builtin error interface
func (e ListGroupReply_GroupValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListGroupReply_Group.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListGroupReply_GroupValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListGroupReply_GroupValidationError{}

// Validate checks the field values on GetUsersReply_User with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetUsersReply_User) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	return nil
}

// GetUsersReply_UserValidationError is the validation error returned by
// GetUsersReply_User.Validate if the designated constraints aren't met.
type GetUsersReply_UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersReply_UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersReply_UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersReply_UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersReply_UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersReply_UserValidationError) ErrorName() string {
	return "GetUsersReply_UserValidationError"
}

// Error satisfies the builtin error interface
func (e GetUsersReply_UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersReply_User.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersReply_UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersReply_UserValidationError{}
