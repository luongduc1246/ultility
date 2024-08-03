package gormdb

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	db, err := Connect(
		DbConfig{
			Url: "postgres://luongduc1246:Postgr3s@76uC1246@localhost:2235/naturalbuilder",
		},
	)
	fmt.Println(db, err)
}
