package user

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/quzhen12/plugins/db"
)

type repo struct {
	DB *gorm.DB
}

func NewMysqlRepo() Repository {
	return &repo{
		DB: db.Client,
	}
}

func (r *repo) CreateUser(ctx context.Context, email, password string) error {
	u := &User{
		Email:    email,
		Password: password,
	}
	return r.DB.Create(u).Error
}

func (r *repo) FindUserByUsername(ctx context.Context, username string) (*User, error) {
	u := &User{}
	result := r.DB.Where("email=?", username).First(&u)
	return u, result.Error
}

func (r *repo) DeleteUserById(ctx context.Context, id int) (int, error) {
	result := r.DB.Where("id=?", id).Delete(&User{})
	return 0, result.Error
}

func (r *repo) UpdateUser(ctx context.Context, user *User) (int64, error) {
	result := r.DB.Model(&User{}).Where("email=?", user.Email).Update(user)
	return result.RowsAffected, result.Error
}
