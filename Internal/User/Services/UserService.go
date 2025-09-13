package Services

import (
	"Hiro/Internal/User/Repositories"
	"Hiro/Models"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, user *Models.User) error
	GetUser(ctx context.Context, id uint) (*Models.User, error)
	GetUserWithPosts(ctx context.Context, id uint) (*Models.User, error)
	GetAllUsers(ctx context.Context) ([]Models.User, error)
}

type userService struct {
	userRepo Repositories.UserRepository
}

func NewUserService(userRepo Repositories.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(ctx context.Context, user *Models.User) error {
	return s.userRepo.Create(ctx, user)
}

func (s *userService) GetUser(ctx context.Context, id uint) (*Models.User, error) {
	return s.userRepo.FindByID(ctx, id)
}

func (s *userService) GetUserWithPosts(ctx context.Context, id uint) (*Models.User, error) {
	return s.userRepo.FindWithPosts(ctx, id)
}

func (s *userService) GetAllUsers(ctx context.Context) ([]Models.User, error) {
	return s.userRepo.FindAll(ctx)
}
