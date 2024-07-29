package postgres

import "errors"

var (
	ErrorExist              = errors.New("model is exist")
	ErrorCreateRole         = errors.New("can't create role")
	ErrorViolatesForeignKey = errors.New("violates key")
	ErrorRecordNotFound     = errors.New("model unexist")
)
