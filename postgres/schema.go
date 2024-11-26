package postgres

import "gorm.io/gorm"

type Schema struct {
	CatalogName string
	SchemaName  string
	SchemaOwner string
}

func CreateSchema(tx *gorm.DB, schemaName string) {
	var schema *Schema
	tx.Raw("SELECT * FROM information_schema.schemata WHERE schema_name =" + schemaName).Scan(&schema)
	if schema == nil {
		tx.Exec("create schema " + schemaName)
	}
}
