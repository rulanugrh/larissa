package repository

import (
	"github.com/rulanugrh/larissa/internal/config"
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/util"
)

type UserInterface interface {
	Create(req domain.UserRegister) (*domain.User, error)
	Find(req domain.UserLogin) (*domain.User, error)
	GetDoctor() (*[]domain.User, error)
	GetNurse() (*[]domain.User, error)
	Update(id uint, req domain.User) error
	ListAll() (*[]domain.User, error)
	Delete(id uint) error
}

type user struct {
	client *config.Postgres
}

func NewUser(client *config.Postgres) UserInterface {
	return &user{
		client: client,
	}
}
func (u *user) Create(req domain.UserRegister) (*domain.User, error) {
	create := domain.User{
		FName:    req.FName,
		LName:    req.LName,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		Age:      req.Age,
		TTL:      req.TTL,
		RoleID:   req.RoleID,
	}

	err := u.client.DB.Create(&create).Error
	if err != nil {
		return nil, err
	}

	return &create, nil
}

func (u *user) Find(req domain.UserLogin) (*domain.User, error) {
	var find domain.User
	rows := u.client.DB.Where("email = ?", req.Email).Find(&find)

	if rows.RowsAffected != 0 {
		return nil, util.EmailUsed()
	}

	return &find, nil
}

func (u *user) Update(id uint, req domain.User) error {
	var update domain.User
	err := u.client.DB.Model(&req).Where("id = ?", id).Updates(&update).Error
	if err != nil {
		return util.NotFound()
	}

	return nil
}

func (u *user) ListAll() (*[]domain.User, error) {
	var find []domain.User

	finds := u.client.DB.Find(&find)
	if finds.RowsAffected == 0 {
		return nil, util.NotFound()
	}

	if finds.Error != nil {
		return nil, util.Errors(finds.Error)
	}

	return &find, nil
}

func (u *user) Delete(id uint) error {
	var model domain.User
	err := u.client.DB.Where("id = ?", id).Delete(&model).Error
	if err != nil {
		return util.Errors(err)
	}

	return nil
}

func (u *user) GetDoctor() (*[]domain.User, error) {
	var model []domain.User
	err := u.client.DB.Preload("Role").Where("id = ?", 2).Find(&model)

	if err.RowsAffected == 0 {
		return nil, util.NotFound()
	}

	return &model, nil
}

func (u *user) GetNurse() (*[]domain.User, error) {
	var model []domain.User
	err := u.client.DB.Preload("Role").Where("id = ?", 3).Find(&model)

	if err.RowsAffected == 0 {
		return nil, util.NotFound()
	}

	return &model, nil
}
