package logic

import (
	"errors"
	"os"
	"xyz-multifinance/internal/model"
	"xyz-multifinance/pkg"
)

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
