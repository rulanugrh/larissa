package service

import (
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/entity/web"
	"github.com/rulanugrh/larissa/internal/middleware"
	"github.com/rulanugrh/larissa/internal/repository"
	"github.com/rulanugrh/larissa/internal/util"
	"golang.org/x/crypto/bcrypt"
)

type UserInterface interface {
	Register(req domain.UserRegister) (*web.User, error)
	Login(req domain.UserLogin) (*string, error)
	Update(id uint, req domain.User) error
}

type user struct {
	urepo repository.UserInterface
	validate middleware.IValidation
}

func NewUser(urepo repository.UserInterface) UserInterface {
	return &user{
		urepo: urepo,
		validate: middleware.NewValidation(),
	}
}

func(u *user) Register(req domain.UserRegister) (*web.User, error) {
	err := u.validate.Validate(req)
	if err != nil {
		return nil, u.validate.Error(err)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nil, util.Errors(err)
	}

	created := domain.UserRegister{
		FName: req.FName,
		LName: req.LName,
		TTL: req.TTL,
		Age: req.Age,
		Address: req.Address,
		Password: string(hashPassword),
		RoleID: req.RoleID,
	}

	data, err := u.urepo.Create(created)
	if err != nil {
		return nil, util.Errors(err)
	}

	response := web.User{
		ID: data.ID,
		Email: data.Email,
		FName: data.FName,
		LName: data.LName,
		Age: data.Age,
		Address: data.Address,
		TTL: data.TTL,
	}


	return &response, nil
}

func(u *user) Login(req domain.UserLogin) (*string, error) {
	err := u.validate.Validate(req)
	if err != nil {
		return nil, u.validate.Error(err)
	}

	data, err := u.urepo.Find(req)
	if err != nil {
		return nil, util.Errors(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if err != nil {
		return nil, util.Errors(err)
	}

	token, err := middleware.GenerateToken(data.ID, data.RoleID, req.Email)
	if err != nil {
		return nil, err
	}

	return &token, nil
}


func(u *user) Update(id uint, req domain.User) error {
	err := u.urepo.Update(id, req)
	if err != nil {
		return util.Errors(err)
	}

	return nil
}
