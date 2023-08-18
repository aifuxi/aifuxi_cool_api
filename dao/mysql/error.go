package mysql

import "errors"

var (
	ErrorTagNotFound = errors.New("tag not found")
	ErrorTagExists   = errors.New("tag already exists")
)

var (
	ErrorUserNotFound = errors.New("user not found")
	ErrorUserExists   = errors.New("user already exists")
)
