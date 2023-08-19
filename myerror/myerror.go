package myerror

import "errors"

var (
	ErrorTagNotFound = errors.New("tag not found")
	ErrorTagExists   = errors.New("tag already exists")
)

var (
	ErrorUserNotFound      = errors.New("user not found")
	ErrorUserExists        = errors.New("user already exists")
	ErrorIncorrectPassword = errors.New("incorrect password ")
)

var (
	ErrorNoAuthorized     = errors.New("no authorized")
	ErrorInvalidToken     = errors.New("invalid token")
	ErrorParseTokenFailed = errors.New("failed to parse token")
)
