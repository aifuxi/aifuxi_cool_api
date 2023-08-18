package mysql

import "errors"

var (
	ErrorTagNotFound = errors.New("tag not found")
	ErrorTagExists   = errors.New("tag already exists")
)
