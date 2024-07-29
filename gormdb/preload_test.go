package gormdb

import (
	"fmt"
	"os"
	"testing"

	"github.com/luongduc1246/ultility/reqparams"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// func PreloadFromFieldPreload(tx *gorm.DB, f *fieldPreload) *gorm.DB {
// 	if len(f.Columns) > 0 {
// 		tx = tx.Clauses(clause.Select{
// 			Columns: f.Columns,
// 		})
// 	}
// 	if len(f.Childs) > 0 {
// 		for key, fc := range f.Childs {
// 			if tx.Statement.Preloads == nil {
// 				tx.Statement.Preloads = map[string][]interface{}{}
// 			}
// 			tx.Statement.Preloads[key] = []interface{}{func(db *gorm.DB) *gorm.DB {
// 				return PreloadFromFieldPreload(db, fc)
// 			}}
// 		}
// 	}
// 	return tx
// }
// func PreloadFromFieldPreload2(tx *gorm.DB, f *fieldPreload) *gorm.DB {
// 	if len(f.Columns) > 0 {
// 		tx = tx.Clauses(clause.Select{
// 			Columns: f.Columns,
// 		})
// 	} else {
// 		tx = tx.Clauses(clause.Select{})
// 	}
// 	if len(f.Childs) > 0 {
// 		for key, fc := range f.Childs {
// 			tx.Preload(key, func(db *gorm.DB) *gorm.DB {
// 				return PreloadFromFieldPreload2(db, fc)
// 			})
// 		}
// 	}
// 	return tx
// }

type Permission struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"type:uuid;unique;default:uuid_generate_v4();->"`
	Name        string    `gorm:"unique;not null;size:500"`
	Status      int       `gorm:"not null"`
	Description *string
}

type Role struct {
	gorm.Model
	UUID        uuid.UUID `gorm:"type:uuid;unique;default:uuid_generate_v4();->"`
	Name        string    `gorm:"unique;not null;size:200"`
	Status      int       `gorm:"not null"`
	Description *string
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

func TestFieldPreload(t *testing.T) {

	// var config = &gorm.Config{
	// 	// Logger: logger.Default.LogMode(logger.Silent),
	// 	// SkipDefaultTransaction: true,
	// 	NamingStrategy: schema.NamingStrategy{
	// 		TablePrefix: "auth.ntb_", // table name prefix, table for `User` would be `goqh_users`
	// 	},
	// }
	txtUri, bol := os.LookupEnv("DATABASE_URL")
	if !bol {
		txtUri = "postgres://luongduc1246:Postgr3s@76uC1246@localhost:2235/naturalbuilder"
	}
	instance, _ = gorm.Open(postgres.Open(txtUri))
	// users := []User{}
	instance.NamingStrategy = schema.NamingStrategy{
		TablePrefix: "auth.ntb_", // table name prefix, table for `User` would be `goqh_users`
	}
	tx := instance.Session(&gorm.Session{})
	tx.Statement.Parse(&Role{})
	fmt.Printf("%+v \n", tx.Statement.Schema)
	a := "uuid,name,permissions[name,uuid]"
	f := reqparams.NewField()
	f.Parse(a)
	fp := NewFieldPreload()
	fp.Parse(tx.Statement.Schema, f)
	fmt.Println(fp)
	bx := fp.BuildPreload(tx)
	users := []Role{}
	bx.Debug().Find(&users)
	fmt.Println(users[2])
}
