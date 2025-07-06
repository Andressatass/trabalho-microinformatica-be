package repository

import (
	"github.com/Andressatass/trabalho-microinformatica-be/trabalho-microinformatica-be/db/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) Create(user entities.UserInfo) (uint, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (r UserRepository) FindById(id string) (entities.UserInfo, error) {
	var user entities.UserInfo

	result := r.db.Find(&user).Where("uuid = ?", user.ID)
	if result.Error != nil {
		return entities.UserInfo{}, result.Error
	}

	return user, nil
}
