package repositories

import (
	"AuthService/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthrepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) FindByUsername(username string) (*models.MUser) {
	var user models.MUser
	r.db.Where("username = ?", username).First(&user)

	if user == (models.MUser{}) {
		return nil
	}

	return &user
}