package gormdb

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/luongduc1246/ultility/reqparams"

	"github.com/google/uuid"
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
	Roles       []Role `gorm:"many2many:role_permissions;"`
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
	instance, _ = Connect(DbConfig{Url: txtUri})
	// users := []User{}
	instance.NamingStrategy = schema.NamingStrategy{
		TablePrefix: "auth.ntb_", // table name prefix, table for `User` would be `goqh_users`
	}
	tx := instance
	go func() {
		stm, _ := schema.ParseWithSpecialTableName(&Role{}, &sync.Map{}, tx.Statement.NamingStrategy, "")

		a := "uuid,name,permissions[name,uuid]"
		f := reqparams.NewField()
		f.Parse(a)
		fp := NewFieldPreload()
		fp.Parse(stm, f)
		bx := fp.BuildPreload(tx)
		users := []Role{}
		bx.Debug().Find(&users)
	}()
	go func() {

		tx.Statement.Parse(&Permission{})
		field := "uuid,name,roles[name,uuid]"
		rqf := reqparams.NewField()
		rqf.Parse(field)
		nfr := NewFieldPreload()
		nfr.Parse(tx.Statement.Schema, rqf)
		brd := nfr.BuildPreload(tx)
		pers := []Permission{}
		brd.Debug().Find(&pers)

	}()
	time.Sleep(2 * time.Second)
}
