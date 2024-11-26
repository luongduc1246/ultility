package gormdb

import (
	"fmt"
	"testing"
)

type Schema struct {
	CatalogName string
	SchemaName  string
	SchemaOwner string
}

func TestConnect(t *testing.T) {
	db, err := Connect(
		DbConfig{
			Url: "postgres://luongduc1246:Postgr3s@76uC1246@localhost:2235/naturalbuilder",
		},
	)
	var schema *Schema
	db.Raw("SELECT * FROM information_schema.schemata WHERE schema_name = 'v'").Scan(&schema)
	fmt.Println(schema)
	fmt.Println(db, err)
}
