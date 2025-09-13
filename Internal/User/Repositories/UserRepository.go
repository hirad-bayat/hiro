package Repositories

import (
	"Hiro/Models"
	"context"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *Models.User) error
	FindByID(ctx context.Context, id uint) (*Models.User, error)
	FindAll(ctx context.Context) ([]Models.User, error)
	FindWithPosts(ctx context.Context, id uint) (*Models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *Models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*Models.User, error) {
	var user Models.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindWithPosts(ctx context.Context, id uint) (*Models.User, error) {
	var user Models.User
	err := r.db.WithContext(ctx).Preload("Posts").First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindAll(ctx context.Context) ([]Models.User, error) {
	var users []Models.User
	err := r.db.WithContext(ctx).Find(&users).Error
	return users, err
}
