package service

import (
	"go-login-crud/dto"
	"go-login-crud/internal/repository"
	"go-login-crud/internal/util"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Auth(body dto.AuthReq) (dto.AuthRes, error)
}

type authSvc struct {
	userRepo repository.UserRepository
}

func CreateAuthService(user repository.UserRepository) AuthService {
	return &authSvc{userRepo: user}
}

func (u *authSvc) Auth(body dto.AuthReq) (dto.AuthRes, error) {
	user, err := u.userRepo.FindByUsername(body.Username)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return dto.AuthRes{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		log.Printf("error: %s", err.Error())
		return dto.AuthRes{}, err
	}

	tkn := util.NewToken(body.Username)
	token, err := tkn.CreateToken()
	if err != nil {
		return dto.AuthRes{}, err
	}

	return dto.AuthRes{
		Token: token,
	}, nil
}

