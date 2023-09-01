package user

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"demo-gorm-gen/internal/infrastructure/database/gormgen"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type Repo interface {
	GetByID(ctx context.Context, ID uint64) (User, error)
}

func New(db *gorm.DB) Repo {
	return &repoImpl{
		db: db,
	}
}

type repoImpl struct {
	db *gorm.DB
}

func (r *repoImpl) GetByID(ctx context.Context, ID uint64) (User, error) {
	u := gormgen.User

	user, err := u.WithContext(ctx).Where(u.ID.Eq(ID)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, ErrUserNotFound
		}
		return User{}, err
	}

	return User{
		ID:   user.ID,
		Name: user.Name,
		Age:  user.Age,
	}, nil
}
