package models

import "errors"

var ErrNotFound = errors.New("object not found")
var ErrTooManyRows = errors.New("too many rows returned")
var ErrInvalidID = errors.New("invalid id")
var ErrEmailNotSent = errors.New("email not sent")
