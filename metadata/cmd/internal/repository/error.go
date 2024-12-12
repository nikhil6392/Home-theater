package repository

import "errors"

//ErrNotFound is returned when a requested record is not Found

var ErrNotFound = errors.New("Not Found")
