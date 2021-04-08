package user_repository

import "errors"

var ErrRecordNotFound = errors.New("Record Not Found")
var ErrRecordAlreadyExist = errors.New("Record already exist")
