package logic

import (
	"errors"
	"os"
	"xyz-multifinance/internal/model"
	"xyz-multifinance/pkg"

	"github.com/golang-jwt/jwt/v5"
)

type (
	JWT struct {
		Token string
	}
)

func (logic Logic) Register(user model.User) error {
	if err := pkg.ValidateStruct(user); err != nil {
		return ErrInvalidArgument(err)
	}

	hashed, err := pkg.Hash(user.Password)
	if err != nil {
		return ErrInternal(err)
	}

	user.Password = hashed

	if err := logic.Repository.CreateUser(user); err != nil {
		return ErrInternal(err)
	}

	return nil
}

func (logic Logic) Login(user model.User) (*JWT, error) {
	var (
		plain  string
		hashed string
	)
	if err := pkg.ValidateStruct(user); err != nil {
		return &JWT{}, ErrInvalidArgument(err)
	}
	plain = user.Password

	user, err, ok := logic.Repository.FindUserByEmail(user.Email)

	if err != nil {
		return &JWT{}, ErrInternal(err)
	}

	if !ok {
		return &JWT{}, ErrInvalidArgument(errors.New("incorrect email or password"))
	}

	hashed = user.Password

	if !pkg.CheckHash(plain, hashed) {
		return &JWT{}, ErrInvalidArgument(errors.New("incorrect email or password"))
	}

	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := pkg.GenerateJWT(jwt.MapClaims{"user_id": user.ID}, secret)
	if err != nil {
		return &JWT{}, ErrInternal(err)
	}

	return &JWT{Token: token}, nil
}

func (logic Logic) AuthUser(token string) (*model.User, error) {
	if token == "" {
		return &model.User{}, ErrIllegal(errors.New("no token"))
	}

	secret := []byte(os.Getenv("JWT_SECRET"))
	claims, err := pkg.ParseJWT(token, secret)
	if err != nil {
		return &model.User{}, ErrIllegal(errors.New("unauthorized"))
	}

	userId, ok := claims["user_id"].(float64)
	if !ok {
		return &model.User{}, ErrInternal(errors.New("failed assert"))
	}

	user, err, ok := logic.Repository.FindUserById(int(userId))

	if !ok {
		return &model.User{}, ErrNotFound(errors.New("user not found"))
	}

	if err != nil {
		return &model.User{}, ErrInternal(err)
	}

	return &user, nil
}

func (logic Logic) RegisterInstallment(userId int, userDetail model.UserDetail) error {
	if err := pkg.ValidateStruct(userDetail); err != nil {
		return ErrInvalidArgument(err)
	}

	userDetail.UserID = uint(userId)
	userDetail.IsVerified = true
	err := logic.Repository.CreateUserDetail(userDetail)
	if err != nil {
		return ErrInternal(errors.New("failed update userDetail"))
	}

	return nil
}
