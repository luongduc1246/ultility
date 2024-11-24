package gormdb

import (
	"github.com/luongduc1246/ultility/reqparams"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type fieldPreload struct {
	Columns []clause.Column
	Childs  map[string]*fieldPreload
}

func NewFieldPreload() *fieldPreload {
	return &fieldPreload{
		Columns: make([]clause.Column, 0),
		Childs:  make(map[string]*fieldPreload),
	}
}

func (f *fieldPreload) BuildPreload(tx *gorm.DB) *gorm.DB {
	if len(f.Columns) > 0 {
		tx = tx.Clauses(clause.Select{
			Columns: f.Columns,
		})
	}
	if len(f.Childs) > 0 {
		for key, fc := range f.Childs {
			if tx.Statement.Preloads == nil {
				tx.Statement.Preloads = map[string][]interface{}{}
			}
			tx.Statement.Preloads[key] = []interface{}{func(db *gorm.DB) *gorm.DB {
				return fc.BuildPreload(db)
			}}
		}
	}
	return tx
}

func (fp *fieldPreload) Parse(scm *schema.Schema, field *reqparams.Fields) {
	if field != nil {
		fieldsByDB := scm.FieldsByDBName
		for _, name := range field.Columns {
			if _, ok := fieldsByDB[name]; ok {
				fp.Columns = append(fp.Columns, clause.Column{
					Table: scm.Table,
					Name:  name,
				})
			}
		}
		if field.Relatives != nil {
			for key, f := range field.Relatives {
				relations := scm.Relationships.Relations
				if rs, ok := relations[key]; ok {
					if len(fp.Columns) > 0 {
						fp.Columns = append(fp.Columns, clause.PrimaryColumn)
					}
					fpc := NewFieldPreload()
					fpc.Parse(rs.FieldSchema, f)
					fpc.Columns = append(fpc.Columns, clause.PrimaryColumn)
					if len(f.Columns) > 0 {
						fp.Childs[key] = fpc
					}
				}
			}
		}
	}
}
