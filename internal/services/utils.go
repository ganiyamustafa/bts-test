package services

import (
	apperror "github.com/ganiyamustafa/bts/utils/app_error"
	"gorm.io/gorm"
)

func deferTransaction(tx *gorm.DB, err *apperror.AppError) {
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
}
