package postgres

import "gorm.io/gorm"

func MirgationModels(tx *gorm.DB, models ...interface{}) error {
	for _, v := range models {
		if !tx.Migrator().HasTable(v) {
			err := tx.AutoMigrate(v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func MirgationModelWithTrigger(tx *gorm.DB, model interface{}) error {
	err := tx.AutoMigrate(model)
	if err != nil {
		return err
	}
	/* add trigger update and insert */
	stmt := &gorm.Statement{DB: tx}
	stmt.Parse(model)
	scm := stmt.Schema
	txTg := tx.Exec(`CREATE TRIGGER prevent_manual_id_insert_trigger
				BEFORE INSERT on ` + scm.Table + `
				 FOR EACH ROW
				EXECUTE FUNCTION prevent_manual_id_insert()`)
	if txTg.Error != nil {
		return txTg.Error
	}
	txTg = tx.Exec(`CREATE TRIGGER prevent_id_update_trigger
				BEFORE UPDATE on ` + scm.Table + `
				 FOR EACH ROW
				EXECUTE FUNCTION prevent_id_update()`)
	if txTg.Error != nil {
		return txTg.Error
	}
	return nil
}
