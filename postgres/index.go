package postgres

import (
	"context"

	"github.com/luongduc1246/ultility/gormdb"
	"github.com/luongduc1246/ultility/reqparams"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres(tx *gorm.DB) *Postgres {
	return &Postgres{
		db: tx,
	}
}

func (pgr Postgres) Create(ctx context.Context, model interface{}, extra ...interface{}) (err error) {
	tx := pgr.db
	scopes, clauses := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	tx = tx.Debug().Create(model)
	err = tx.Error
	if err != nil {
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
	return nil
}
func (pgr Postgres) CreateBatch(ctx context.Context, models interface{}, extra ...interface{}) (err error) {
	tx := pgr.db
	scopes, clauses := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	tx = tx.Debug().Create(models)
	err = tx.Error
	if err != nil {
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
	return nil
}

func (pgr Postgres) Update(ctx context.Context, info interface{}, extra ...interface{}) (err error) {
	tx := pgr.db
	scopes, clauses := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	tx = tx.Debug().Updates(info)
	err = tx.Error
	if err != nil {
		switch er := err.(type) {
		case *pgconn.PgError:
			if er.Code == "23505" {
				return ErrorExist
			}
			if er.Code == "23503" {
				return ErrorViolatesForeignKey
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
	scopes, clauses := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	tx = tx.Begin()
	for _, info := range infos {
		tx := tx.Session(&gorm.Session{Initialized: true})
		txUp := tx.Debug().Updates(info)
		err = txUp.Error
		if err != nil {
			tx.Rollback()
			switch er := err.(type) {
			case *pgconn.PgError:
				if er.Code == "23505" {
					return ErrorExist
				}
				if er.Code == "23503" {
					return ErrorViolatesForeignKey
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
		scopes, clauses := parseExtra(info.Extra...)
		if len(scopes) > 0 {
			tx = tx.Scopes(scopes...)
		}
		if len(clauses) > 0 {
			tx = tx.Clauses(clauses...)
		}
		txUp := tx.Debug().Updates(info.Model)
		err = txUp.Error
		if err != nil {
			tx.Rollback()
			switch er := err.(type) {
			case *pgconn.PgError:
				if er.Code == "23505" {
					return ErrorExist
				}
				if er.Code == "23503" {
					return ErrorViolatesForeignKey
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
		scopes, clauses := parseExtra(extra...)
		if len(scopes) > 0 {
			tx = tx.Scopes(scopes...)
		}
		if len(clauses) > 0 {
			tx = tx.Clauses(clauses...)
		}
		for key, value := range asm.Associations {
			tx := tx.Session(&gorm.Session{Initialized: true})
			err = tx.Model(asm.Model).Debug().Clauses(clause.Returning{}).Association(key).Append(value.Model)
			if err != nil {
				switch er := err.(type) {
				case *pgconn.PgError:
					if er.Code == "23503" {
						return ErrorViolatesForeignKey
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
		scopes, clauses := parseExtra(extra...)
		if len(scopes) > 0 {
			tx = tx.Scopes(scopes...)
		}
		if len(clauses) > 0 {
			tx = tx.Clauses(clauses...)
		}
		for key, value := range asm.Associations {
			tx := tx.Session(&gorm.Session{Initialized: true})
			err = tx.Model(asm.Model).Debug().Clauses(clause.Returning{}).Association(key).Delete(value.Model)
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
	scopes, clauses := parseExtra(extra...)
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
	scopes, clauses := parseExtra(extra...)
	if len(scopes) > 0 {
		tx = tx.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		tx = tx.Clauses(clauses...)
	}
	for _, info := range infos {
		tx := tx.Session(&gorm.Session{Initialized: true})
		txUp := tx.Debug().Unscoped().Clauses(clause.Returning{}).Model(info).Update("DeletedAt", nil)
		err = txUp.Error
		if err != nil {
			tx.Rollback()
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
	scopes, clauses := parseExtra(extra...)
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
	tx := pgr.db.Session(&gorm.Session{NewDB: true})
	tx.Statement.Parse(models)
	cs := gormdb.NewClauseSearch()
	cs.Parse(tx.Statement.Schema, reqPamrams)
	exps := cs.Build()
	fieldPreload := gormdb.NewFieldPreload()
	fieldPreload.Parse(tx.Statement.Schema, reqPamrams.Field)
	fb := fieldPreload.BuildPreload(tx)
	/* làm việc với extra */
	scopes, clauses := parseExtra(extra...)
	if len(scopes) > 0 {
		fb = fb.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		fb = fb.Clauses(clauses...)
	}
	fb.Debug().Clauses(exps...).Find(models)
	err = fb.Error
	if err != nil {
		return err
	}
	return nil
}
func (pgr Postgres) SearchSoftDelete(ctx context.Context, reqPamrams *reqparams.Search, models interface{}, extra ...interface{}) (err error) {
	tx := pgr.db.Session(&gorm.Session{NewDB: true})
	tx.Statement.Parse(models)
	cs := gormdb.NewClauseSearch()
	cs.Parse(tx.Statement.Schema, reqPamrams)
	exps := cs.Build()
	fieldPreload := gormdb.NewFieldPreload()
	fieldPreload.Parse(tx.Statement.Schema, reqPamrams.Field)
	fb := fieldPreload.BuildPreload(tx)
	/* làm việc với extra */
	scopes, clauses := parseExtra(extra...)
	if len(scopes) > 0 {
		fb = fb.Scopes(scopes...)
	}
	if len(clauses) > 0 {
		fb = fb.Clauses(clauses...)
	}
	fb.Debug().Unscoped().Clauses(exps...).Where("deleted_at IS NOT NULL").Find(models)
	err = fb.Error
	if err != nil {
		return err
	}
	return nil
}

/* phân tích extra ra scopes và clauses */
func parseExtra(extra ...interface{}) ([]func(*gorm.DB) *gorm.DB, []clause.Expression) {
	scopes := []func(*gorm.DB) *gorm.DB{}
	clauses := []clause.Expression{}
	for _, val := range extra {
		switch v := val.(type) {
		case func(*gorm.DB) *gorm.DB:
			scopes = append(scopes, v)
		default:
			if c, ok := v.(clause.Expression); ok {
				clauses = append(clauses, c)
			}
		}
	}
	return scopes, clauses
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
