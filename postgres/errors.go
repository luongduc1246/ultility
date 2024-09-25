package postgres

import "errors"

var (
	ErrorExist              = errors.New("model is exist")
	ErrorCreateRole         = errors.New("can't create role")
	ErrorViolatesForeignKey = errors.New("violates key")
	ErrorRecordNotFound     = errors.New("model unexist")
	ErrorManualInsertID     = errors.New("cannot manually insert a value into the id")
	ErrorManualUpdateID     = errors.New("cannot manually update a value into the id")
)
