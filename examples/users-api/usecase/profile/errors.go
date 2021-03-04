package profile_usecase

import "errors"

var ErrDatabaseError = errors.New("Database Error")
var ErrRecordNotFound = errors.New("Record Not Found")
var ErrClient = errors.New("Err Client")
