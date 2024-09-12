package gormdb

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func LoadSchema(db *gorm.DB, model any) *schema.Schema {
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(model)
	return stmt.Schema
}
