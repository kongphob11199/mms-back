package message

import "errors"

var (
	ErrorUserNotFound       = errors.New("NOTFOUND_USER")
	ErrorUserCreateCustomer = errors.New("FAILED_CREATE_USER")
	ErrorUserUpdateCustomer = errors.New("FAILED_UPDATE_USER")
	ErrorUserCreate         = errors.New("FAILED_CREATE_USER")
	ErrorUserUPDATE         = errors.New("FAILED_UPDATE_USER")
	ErrorUserDup            = errors.New("USER_DUP")
	ErrorPasswordHash       = errors.New("FAILED_TO_HASH")
)
