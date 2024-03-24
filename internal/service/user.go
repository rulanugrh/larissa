package service

import (
	"github.com/rulanugrh/larissa/internal/entity/domain"
	"github.com/rulanugrh/larissa/internal/entity/web"
	"github.com/rulanugrh/larissa/internal/middleware"
	"github.com/rulanugrh/larissa/internal/repository"
	"github.com/rulanugrh/larissa/internal/util"
	"github.com/rulanugrh/larissa/pkg"
	"golang.org/x/crypto/bcrypt"
)

type UserInterface interface {
	Register(req domain.UserRegister) (*web.User, error)
	Login(req domain.UserLogin) (*string, error)
	Update(id uint, req domain.User) error
	GotDoctor() (*[]web.User, error)
	GotNurse() (*[]web.User, error)
}

type user struct {
	urepo    repository.UserInterface
	validate middleware.IValidation
	log pkg.ILogrust
}

func NewUser(urepo repository.UserInterface, log pkg.ILogrust) UserInterface {
	return &user{
		urepo:    urepo,
		validate: middleware.NewValidation(),
		log: log,
	}
}

func (u *user) Register(req domain.UserRegister) (*web.User, error) {
	err := u.validate.Validate(req)
	if err != nil {
		u.log.StartLogger("user_service", "register").Error("cannot validate field")
		return nil, u.validate.Error(err)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {

		u.log.StartLogger("user_service", "register").Error("cannot hash password")
		return nil, util.Errors(err)
	}

	created := domain.UserRegister{
		FName:    req.FName,
		LName:    req.LName,
		TTL:      req.TTL,
		Age:      req.Age,
		Address:  req.Address,
		Password: string(hashPassword),
		RoleID:   req.RoleID,
	}

	data, err := u.urepo.Create(created)
	if err != nil {

		u.log.StartLogger("user_service", "register").Error("cannot create into db")
		return nil, util.Errors(err)
	}

	response := web.User{
		ID:      data.ID,
		Email:   data.Email,
		FName:   data.FName,
		LName:   data.LName,
		Age:     data.Age,
		Address: data.Address,
		TTL:     data.TTL,
	}

	u.log.StartLogger("user_service", "register").Info("sucess create user")
	return &response, nil
}

func (u *user) Login(req domain.UserLogin) (*string, error) {
	err := u.validate.Validate(req)
	if err != nil {

		u.log.StartLogger("user_service", "login").Error("cannot validate field")
		return nil, u.validate.Error(err)
	}

	data, err := u.urepo.Find(req)
	if err != nil {

		u.log.StartLogger("user_service", "login").Error("cannot find email in DB")
		return nil, util.Errors(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if err != nil {

		u.log.StartLogger("user_service", "login").Error("cannot compare password")
		return nil, util.Errors(err)
	}

	token, err := middleware.GenerateToken(data.ID, data.RoleID, req.Email)
	if err != nil {

		u.log.StartLogger("user_service", "login").Error("cannot generate token")
		return nil, err
	}


	u.log.StartLogger("user_service", "login").Info("success login")
	return &token, nil
}

func (u *user) Update(id uint, req domain.User) error {
	err := u.urepo.Update(id, req)
	if err != nil {

		u.log.StartLogger("user_service", "update").Error(err.Error())
		return util.Errors(err)
	}


	u.log.StartLogger("user_service", "update").Info("success update user")
	return nil
}

func (u *user) GotDoctor() (*[]web.User, error) {
	data, err := u.urepo.GetDoctor()
	if err != nil {

		u.log.StartLogger("user_service", "gotDoctor").Error(err.Error())
		return nil, util.Errors(err)
	}

	var response []web.User
	for _, v := range *data {
		result := web.User{
			ID:      v.ID,
			Email:   v.Email,
			FName:   v.FName,
			LName:   v.LName,
			Age:     v.Age,
			Address: v.Address,
		}

		response = append(response, result)
	}

	u.log.StartLogger("user_service", "gotDoctor").Info("success got all doctor")
	return &response, nil
}

func (u *user) GotNurse() (*[]web.User, error) {
	data, err := u.urepo.GetNurse()
	if err != nil {
		u.log.StartLogger("user_service", "gotNurse").Info(err.Error())
		return nil, util.Errors(err)
	}

	var response []web.User
	for _, v := range *data {
		result := web.User{
			ID:      v.ID,
			Email:   v.Email,
			FName:   v.FName,
			LName:   v.LName,
			Age:     v.Age,
			Address: v.Address,
		}

		response = append(response, result)
	}

	u.log.StartLogger("user_service", "gotNurse").Info("success got all nurse")
	return &response, nil
}
