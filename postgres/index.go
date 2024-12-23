package postgres

import (
	"context"
	"errors"
	"sync"

	"github.com/luongduc1246/ultility/gormdb"
	"github.com/luongduc1246/ultility/reqparams"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres(tx *gorm.DB) *Postgres {
	return &Postgres{
		db: tx,
	}
}

func (pgr Postgres) GetDB() *gorm.DB {
	return pgr.db
}

func (pgr Postgres) Create(ctx context.Context, model interface{}, extra ...interface{}) (err error) {
	tx := pgr.db
	scopes, clauses, _ := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	tx = tx.Create(model)
	err = tx.Error
	if err != nil {
		unWrap := errors.Unwrap(err)
		if unWrap != nil {
			err = unWrap
		}
		switch er := err.(type) {
		case *pgconn.PgError:
			if er.Code == "23505" {
				return ErrorExist
			}
			if er.Code == "23503" {
				return ErrorViolatesForeignKey
			}
			if er.Code == "55007" {
				return ErrorManualInsertID
			}
			return err
		default:
			return err
		}
	}
	return nil
}
func (pgr Postgres) CreateBatch(ctx context.Context, models interface{}, extra ...interface{}) (err error) {
	tx := pgr.db
	scopes, clauses, _ := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	tx = tx.Create(models)
	err = tx.Error
	if err != nil {
		unWrap := errors.Unwrap(err)
		if unWrap != nil {
			err = unWrap
		}
		switch er := err.(type) {
		case *pgconn.PgError:
			if er.Code == "23505" {
				return ErrorExist
			}
			if er.Code == "23503" {
				return ErrorViolatesForeignKey
			}
			if er.Code == "55007" {
				return ErrorManualInsertID
			}
			return err
		default:
			return err
		}
	}
	return nil
}

func (pgr Postgres) Update(ctx context.Context, info interface{}, extra ...interface{}) (err error) {
	tx := pgr.db
	scopes, clauses, _ := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	tx = tx.Updates(info)
	err = tx.Error
	if err != nil {
		unWrap := errors.Unwrap(err)
		if unWrap != nil {
			err = unWrap
		}
		switch er := err.(type) {
		case *pgconn.PgError:
			if er.Code == "23505" {
				return ErrorExist
			}
			if er.Code == "23503" {
				return ErrorViolatesForeignKey
			}
			if er.Code == "55008" {
				return ErrorManualUpdateID
			}
			return err
		default:
			return err
		}
	}
	return nil
}

/* Cẩn thận khi dùng extra ( dùng extra là sử dụng chung cho tất cả câu truy vấn từng model ) */
func (pgr Postgres) UpdateBatch(ctx context.Context, infos []interface{}, extra ...interface{}) (err error) {
	tx := pgr.db
	scopes, clauses, _ := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	tx = tx.Begin()
	for _, info := range infos {
		tx := tx.Session(&gorm.Session{Initialized: true})
		txUp := tx.Updates(info)
		err = txUp.Error
		if err != nil {
			tx.Rollback()
			unWrap := errors.Unwrap(err)
			if unWrap != nil {
				err = unWrap
			}
			switch er := err.(type) {
			case *pgconn.PgError:
				if er.Code == "23505" {
					return ErrorExist
				}
				if er.Code == "23503" {
					return ErrorViolatesForeignKey
				}
				if er.Code == "55008" {
					return ErrorManualUpdateID
				}
				return err
			default:
				return err
			}
		}
	}
	tx.Commit()
	return nil
}

func (pgr Postgres) UpdateBatchWithBatchInfo(ctx context.Context, infos []gormdb.BatchInfo) (err error) {
	tx := pgr.db.Begin()
	for _, info := range infos {
		/* làm việc với nhiều dữ liệu mà có sử dụng scope hoặc clause ta nên khởi tạo lại bằng cách dùng session */
		tx := tx.Session(&gorm.Session{Initialized: true})
		scopes, clauses, _ := parseExtra(info.Extra...)
		if len(scopes) > 0 {
			tx = tx.Scopes(scopes...)
		}
		if len(clauses) > 0 {
			tx = tx.Clauses(clauses...)
		}
		txUp := tx.Updates(info.Model)
		err = txUp.Error
		if err != nil {
			tx.Rollback()
			unWrap := errors.Unwrap(err)
			if unWrap != nil {
				err = unWrap
			}
			switch er := err.(type) {
			case *pgconn.PgError:
				if er.Code == "23505" {
					return ErrorExist
				}
				if er.Code == "23503" {
					return ErrorViolatesForeignKey
				}
				if er.Code == "55008" {
					return ErrorManualUpdateID
				}
				return err
			default:
				return err
			}
		}
	}
	tx.Commit()
	return nil
}

/* Cẩn thận khi dùng extra ( dùng extra là sử dụng chung cho tất cả câu truy vấn từng model ) */
func (pgr Postgres) AppendAssociation(ctx context.Context, asm *gormdb.AssociationModel, extra ...interface{}) (err error) {
	err = pgr.db.Transaction(func(txF *gorm.DB) error {
		tx := txF
		scopes, clauses, _ := parseExtra(extra...)
		if len(scopes) > 0 {
			tx = tx.Scopes(scopes...)
		}
		if len(clauses) > 0 {
			tx = tx.Clauses(clauses...)
		}
		for key, value := range asm.Associations {
			tx := tx.Session(&gorm.Session{Initialized: true})
			err = tx.Model(asm.Model).Clauses(clause.Returning{}).Association(key).Append(value.Model)
			if err != nil {
				unWrap := errors.Unwrap(err)
				if unWrap != nil {
					err = unWrap
				}
				switch er := err.(type) {
				case *pgconn.PgError:
					if er.Code == "23505" {
						return ErrorExist
					}
					if er.Code == "23503" {
						return ErrorViolatesForeignKey
					}
					if er.Code == "55007" {
						return ErrorManualInsertID
					}
					return err
				default:
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (pgr Postgres) DeleteAssociation(ctx context.Context, asm *gormdb.AssociationModel, extra ...interface{}) (err error) {
	err = pgr.db.Transaction(func(txF *gorm.DB) error {
		tx := txF
		scopes, clauses, _ := parseExtra(extra...)
		if len(scopes) > 0 {
			tx = tx.Scopes(scopes...)
		}
		if len(clauses) > 0 {
			tx = tx.Clauses(clauses...)
		}
		for key, value := range asm.Associations {
			tx := tx.Session(&gorm.Session{Initialized: true})
			err = tx.Model(asm.Model).Clauses(clause.Returning{}).Association(key).Delete(value.Model)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
func (pgr Postgres) Delete(ctx context.Context, infos interface{}, extra ...interface{}) (err error) {
	tx := pgr.db
	scopes, clauses, _ := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	tx = tx.Delete(infos)
	err = tx.Error
	if err != nil {
		return err
	}
	return nil
}

/* Cẩn thận khi dùng extra ( dùng extra là sử dụng chung cho tất cả câu truy vấn từng model ) */
func (pgr Postgres) RevertDelete(ctx context.Context, infos []interface{}, extra ...interface{}) (err error) {
	tx := pgr.db.Begin()
	scopes, clauses, _ := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	for _, info := range infos {
		tx := tx.Session(&gorm.Session{Initialized: true})
		txUp := tx.Unscoped().Clauses(clause.Returning{}).Model(info).Update("DeletedAt", nil)
		err = txUp.Error
		if err != nil {
			tx.Rollback()
			unWrap := errors.Unwrap(err)
			if unWrap != nil {
				err = unWrap
			}
			switch er := err.(type) {
			case *pgconn.PgError:
				if er.Code == "23505" {
					return ErrorExist
				}
				return err
			default:
				return err
			}
		}
	}
	tx.Commit()
	return nil
}
func (pgr Postgres) DeletePermanently(ctx context.Context, info interface{}, extra ...interface{}) (err error) {
	tx := pgr.db
	scopes, clauses, _ := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	tx = tx.Unscoped().Clauses(clause.Returning{}).Select(clause.Associations).Delete(info)
	err = tx.Error
	if err != nil {
		return err
	}
	return nil
}

func (pgr Postgres) Search(ctx context.Context, reqPamrams *reqparams.Search, models interface{}, extra ...interface{}) (err error) {
	tx := pgr.db
	stm, err := schema.ParseWithSpecialTableName(models, &sync.Map{}, tx.Statement.NamingStrategy, "")
	if err != nil {
		return err
	}
	cs := gormdb.NewClauseSearch()
	cs.Parse(stm, reqPamrams)
	exps := cs.Build()
	if reqPamrams.Fields != nil {
		fieldPreload := gormdb.NewFieldPreload()
		fields := reqparams.NewFields()
		fields.ParseFromQuerier(reqPamrams.Fields)
		fieldPreload.Parse(stm, fields)
		tx = fieldPreload.BuildPreload(tx)
	}
	/* làm việc với extra */
	scopes, clauses, _ := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	tx = tx.Clauses(exps...).Find(models)
	err = tx.Error
	if err != nil {
		return err
	}
	return nil
}
func (pgr Postgres) SearchSoftDelete(ctx context.Context, reqPamrams *reqparams.Search, models interface{}, extra ...interface{}) (err error) {
	tx := pgr.db
	stm, err := schema.ParseWithSpecialTableName(models, &sync.Map{}, tx.Statement.NamingStrategy, "")
	if err != nil {
		return err
	}
	cs := gormdb.NewClauseSearch()
	cs.Parse(stm, reqPamrams)
	exps := cs.Build()
	if reqPamrams.Fields != nil {
		fieldPreload := gormdb.NewFieldPreload()
		fields := reqparams.NewFields()
		fields.ParseFromQuerier(reqPamrams.Fields)
		fieldPreload.Parse(stm, fields)
		tx = fieldPreload.BuildPreload(tx)
	}
	/* làm việc với extra */
	scopes, clauses, _ := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	tx = tx.Unscoped().Clauses(exps...).Where("deleted_at IS NOT NULL").Find(models)
	err = tx.Error
	if err != nil {
		return err
	}
	return nil
}

/* phân tích extra ra scopes và clauses */
func parseExtra(extra ...interface{}) ([]func(*gorm.DB) *gorm.DB, []clause.Expression, []interface{}) {
	scopes := []func(*gorm.DB) *gorm.DB{}
	clauses := []clause.Expression{}
	ext := []interface{}{}
	for _, val := range extra {
		switch v := val.(type) {
		case func(*gorm.DB) *gorm.DB:
			scopes = append(scopes, v)
		case clause.Expression:
			clauses = append(clauses, v)
		default:
			ext = append(ext, v)
		}
	}
	return scopes, clauses, ext
}

/* scope lấy dữ liệu cho update  sử dụng để kiểm tra cache*/
func scopeLoadModel(model interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tx := db.Session(&gorm.Session{Initialized: true}).First(model)
		err := tx.Error
		if err != nil {
			switch err {
			case gorm.ErrRecordNotFound:
				db.AddError(ErrorRecordNotFound)
			}
		}
		return db
	}
}
