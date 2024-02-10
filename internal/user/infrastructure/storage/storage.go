package storage

import (
	"fmt"
	"github.com/lsendoya/cognitoDalas/internal/user/domain"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, user domain.User) error {
	if err := db.Create(&user).Error; err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}
	return nil
}
