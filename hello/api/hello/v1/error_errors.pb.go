// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/limes-cloud/kratos/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 为某个枚举单独设置错误码
func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 200
}

// 为某个枚举单独设置错误码
func ErrorUserNotFoundFormat(format string, args ...interface{}) *errors.Error {
	return errors.New(200, ErrorReason_USER_NOT_FOUND.String(), ":"+fmt.Sprintf(format, args...))
}

// 为某个枚举单独设置错误码
func ErrorUserNotFound() *errors.Error {
	return errors.New(200, ErrorReason_USER_NOT_FOUND.String(), "")
}

func IsContentMissing(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CONTENT_MISSING.String() && e.Code == 200
}

func ErrorContentMissingFormat(format string, args ...interface{}) *errors.Error {
	return errors.New(200, ErrorReason_CONTENT_MISSING.String(), ":"+fmt.Sprintf(format, args...))
}

func ErrorContentMissing() *errors.Error {
	return errors.New(200, ErrorReason_CONTENT_MISSING.String(), "")
}