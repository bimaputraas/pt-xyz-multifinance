package logic

import (
	"errors"
	"gorm.io/gorm"
	"xyz-multifinance/internal/model"
	"xyz-multifinance/pkg"
)

func (logic *Logic) AuthUser(token string) (*model.User, error) {
	if token == "" {
		return &model.User{}, ErrIllegal(errors.New("no token"))
	}

	secret := []byte(logic.config.JWTSecret)
	claims, err := pkg.ParseJWT(token, secret)
	if err != nil {
		return &model.User{}, ErrIllegal(errors.New("invalid token"))
	}

	userId, ok := claims["user_id"].(float64)
	if !ok {
		return &model.User{}, ErrInternal(errors.New("failed assert"))
	}

	user, err := logic.repo.UserRepository.FindById(int(userId))

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &model.User{}, ErrNotFound(errors.New("user not found"))
	}

	if err != nil {
		return &model.User{}, ErrInternal(err)
	}

	return &user, nil
}
