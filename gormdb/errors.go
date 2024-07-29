package gormdb

import (
	"strings"
)

type ErrGormDb struct {
	Err error
}

func (e ErrGormDb) Error() string {
	var mess strings.Builder
	mess.WriteString("GormDB error ")
	mess.WriteString(e.Err.Error())
	return mess.String()
}
