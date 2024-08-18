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
