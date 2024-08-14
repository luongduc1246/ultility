package gormdb

import (
	"context"

	"github.com/luongduc1246/ultility/reqparams"
	"gorm.io/gorm"
)

type BatchInfo struct {
	Model interface{}
	Extra []interface{}
}

type DBRepository interface {
	GetDB() *gorm.DB
	Create(context.Context, interface{}, ...interface{}) error
	CreateBatch(context.Context, interface{}, ...interface{}) error
	Update(context.Context, interface{}, ...interface{}) error
	UpdateBatch(context.Context, []interface{}, ...interface{}) error
	UpdateBatchWithBatchInfo(context.Context, []BatchInfo) error
	Delete(context.Context, interface{}, ...interface{}) error
	RevertDelete(context.Context, []interface{}, ...interface{}) error
	SearchSoftDelete(context.Context, *reqparams.Search, interface{}, ...interface{}) (err error)
	AppendAssociation(context.Context, *AssociationModel, ...interface{}) error
	DeleteAssociation(context.Context, *AssociationModel, ...interface{}) error
	DeletePermanently(context.Context, interface{}, ...interface{}) error
	Search(context.Context, *reqparams.Search, interface{}, ...interface{}) (err error)
}
