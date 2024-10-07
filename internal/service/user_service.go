package service

import (
	"go-login-crud/dto"
	"go-login-crud/internal/repository"
	"go-login-crud/model"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req dto.UserRegist) (string, error)
	Update(req dto.UserData) (string, error)
	FetchData(req dto.UserReq) ([]model.User, error)
	FindId(id string) (model.User, error)
	FindUsername(username string) (model.User, error)
	Delete(id string) (string, error)
}

type userSvc struct {
	userRepo repository.UserRepository
}

func CreateUserService(user repository.UserRepository) UserService {
	return &userSvc{userRepo: user}
}

func (u *userSvc) Register(req dto.UserRegist) (string, error) {
	// exist, err := u.userRepo.FindByUsername(req.Username)
	// if err != nil {
	// 	log.Printf("error: %s", err.Error())
	// 	return "", err
	// }

	// if exist != "" {
	// 	return exist, errors.New("username taken")
	// }

	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 12)

	user := model.User{
		Username: req.Username,
		Phone:    req.Phone,
		Email:    req.Email,
		Password: string(hashedPass),
	}

	result, err := u.userRepo.InsertUser(user)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return "", err
	}

	return result, nil
}

func (u *userSvc) Update(req dto.UserData) (string, error) {
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 12)

	user := model.User{
		UserId: req.UserId,
		Username: req.Username,
		Phone:    req.Phone,
		Email:    req.Email,
		Password: string(hashedPass),
	}

	result, err := u.userRepo.UpdateUser(user)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return "", err
	}

	return result, nil
}

func (u *userSvc) FetchData(req dto.UserReq) ([]model.User, error) {
	result, err := u.userRepo.GetAllUsers(req.Limit, req.Page)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return []model.User{}, err
	}

	return result, nil
}

func (u *userSvc) FindId(id string) (model.User, error) {
	result, err := u.userRepo.FindById(id)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return model.User{}, err
	}

	return result, nil
}

func (u *userSvc) FindUsername(username string) (model.User, error) {
	result, err := u.userRepo.FindByUsername(username)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return model.User{}, err
	}

	return result, nil
}

func (u *userSvc) Delete(id string) (string, error) {
	result, err := u.userRepo.DeleteUser(id)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return "", err
	}

	return result, nil
}