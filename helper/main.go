package helper

import "gorm.io/gorm"

type InTransaction func(tx *gorm.DB) error

func DoInManualQuery(tx *gorm.DB, fn InTransaction) error {
	if tx.Error != nil {
		return tx.Error
	}

	err := fn(tx)
	if err != nil {
		xerr := tx.Rollback().Error
		if xerr != nil {
			return xerr
		}
		return err
	}
	if err = tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
