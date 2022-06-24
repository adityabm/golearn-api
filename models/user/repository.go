package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(user User) (User, error)
	SearchByEmail(email string) (User, error)
	SearchByID(ID int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) SearchByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) SearchByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	
	return user, nil
}