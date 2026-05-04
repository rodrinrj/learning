package store

import "errors"

var ErrKeyNotFound = errors.New("key is not saved in store")
var ErrSetFailed = errors.New("key could not be saved")
var ErrDelFailed = errors.New("key could not be deleted")
var ErrKeysFailed = errors.New("keys could not be listed")
var ErrUnexpected = errors.New("unexpected error")
