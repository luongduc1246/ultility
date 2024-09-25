package postgres

import "gorm.io/gorm"

func CreateFuncTriggerInsert(tx *gorm.DB) error {
	txTgInsert := tx.Exec(`CREATE OR REPLACE FUNCTION prevent_manual_id_insert()
			RETURNS TRIGGER AS $$
			DECLARE 
				id integer;
			BEGIN
				id := nextval(pg_get_serial_sequence(quote_ident(TG_TABLE_SCHEMA)||'.'|| quote_ident(TG_TABLE_NAME) , 'id')) - 1;
			IF NEW.id <> id  THEN
				RAISE EXCEPTION 'cannot manually insert a value into the id column' USING ERRCODE = '55007';
			END IF;
				id =setval(pg_get_serial_sequence(quote_ident(TG_TABLE_SCHEMA)||'.'|| quote_ident(TG_TABLE_NAME) , 'id'),id);
			RETURN NEW;
			END;
			$$ LANGUAGE plpgsql;`)
	if txTgInsert.Error != nil {
		return txTgInsert.Error
	}
	return nil
}

func CreateFuncTriggerUpdate(tx *gorm.DB) error {
	txTgUpdate := tx.Exec(`CREATE OR REPLACE FUNCTION prevent_id_update()
		RETURNS TRIGGER AS $$
		BEGIN
			IF NEW.id <> OLD.id THEN
				RAISE EXCEPTION 'cannot update id field' USING ERRCODE = '55008';
			END IF;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;`)
	if txTgUpdate.Error != nil {
		return txTgUpdate.Error
	}
	return nil
}
