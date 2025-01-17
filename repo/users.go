package repo

import (
	"github.com/bocianowski1/base/models"
	"gorm.io/gorm"
)

type IUserRepo interface {
	Create(user *models.User) error
	FindByID(id string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *models.User) error {
	// TODO: hash password
	return r.db.Create(user).Error
}

func (r *UserRepo) FindByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *UserRepo) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepo) Update(user *models.User) error {
	// TODO: hash password
	return r.db.Save(user).Error
}

func (r *UserRepo) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.User{}).Error
}
