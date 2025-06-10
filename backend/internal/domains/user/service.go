package user

import (
	"pokemon/pkg/utils"

	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
    CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error)
    GetUser(ctx context.Context, id uuid.UUID) (*User, error)
    UpdateUser(ctx context.Context, id uuid.UUID, req *UpdateUserRequest) (*User, error)
    DeleteUser(ctx context.Context, id uuid.UUID) error
    ListUsers(ctx context.Context, limit, offset int) ([]*User, error)
    Login(ctx context.Context, req *LoginRequest) (string, error)
}

type service struct {
    repo Repository
}

func NewService(repo Repository) Service {
    return &service{
        repo: repo,
    }
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    
    user := &User{
        Email:     req.Email,
        Username:  req.Username,
        Password:  string(hashedPassword),
        FirstName: req.FirstName,
        LastName:  req.LastName,
        Active:    true,
    }
    
    err = s.repo.Create(ctx, user)
    if err != nil {
        return nil, err
    }
    
    return user, nil
}

func (s *service) GetUser(ctx context.Context, id uuid.UUID) (*User, error) {
    return s.repo.GetByID(ctx, id)
}

func (s *service) UpdateUser(ctx context.Context, id uuid.UUID, req *UpdateUserRequest) (*User, error) {
    updates := make(map[string]interface{})
    
    if req.FirstName != nil {
        updates["first_name"] = *req.FirstName
    }
    if req.LastName != nil {
        updates["last_name"] = *req.LastName
    }
    if req.Active != nil {
        updates["active"] = *req.Active
    }
    
    err := s.repo.Update(ctx, id, updates)
    if err != nil {
        return nil, err
    }
    
    return s.repo.GetByID(ctx, id)
}

func (s *service) DeleteUser(ctx context.Context, id uuid.UUID) error {
    return s.repo.Delete(ctx, id)
}

func (s *service) ListUsers(ctx context.Context, limit, offset int) ([]*User, error) {
    return s.repo.List(ctx, limit, offset)
}

func (s *service) Login(ctx context.Context, req *LoginRequest) (string, error) {
    user, err := s.repo.GetByEmail(ctx, req.Email)
    if err != nil {
        return "", errors.New("invalid credentials")
    }
    
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
    if err != nil {
        return "", errors.New("invalid credentials")
    }
    
    if !user.Active {
        return "", errors.New("account is deactivated")
    }
    
    token, err := utils.GenerateJWT(user.ID, user.Email)
    if err != nil {
        return "", err
    }
    
    return token, nil
}
