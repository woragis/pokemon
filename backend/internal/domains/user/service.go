package user

import (
	"pokemon/pkg/utils"

	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService interface {
    createUser(ctx context.Context, req *createUserRequest) (*User, error)
    getUser(ctx context.Context, id uuid.UUID) (*User, error)
    updateUser(ctx context.Context, id uuid.UUID, req *updateUserRequest) (*User, error)
    deleteUser(ctx context.Context, id uuid.UUID) error
    listUsers(ctx context.Context, limit, offset int) ([]*User, error)
    login(ctx context.Context, req *loginRequest) (string, error)
}

type service struct {
    repo userRepository
}

func NewService(repo userRepository) userService {
    return &service{
        repo: repo,
    }
}

func (s *service) createUser(ctx context.Context, req *createUserRequest) (*User, error) {
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
    
    err = s.repo.create(ctx, user)
    if err != nil {
        return nil, err
    }
    
    return user, nil
}

func (s *service) getUser(ctx context.Context, id uuid.UUID) (*User, error) {
    return s.repo.getByID(ctx, id)
}

func (s *service) updateUser(ctx context.Context, id uuid.UUID, req *updateUserRequest) (*User, error) {
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
    
    err := s.repo.update(ctx, id, updates)
    if err != nil {
        return nil, err
    }
    
    return s.repo.getByID(ctx, id)
}

func (s *service) deleteUser(ctx context.Context, id uuid.UUID) error {
    return s.repo.delete(ctx, id)
}

func (s *service) listUsers(ctx context.Context, limit, offset int) ([]*User, error) {
    return s.repo.list(ctx, limit, offset)
}

func (s *service) login(ctx context.Context, req *loginRequest) (string, error) {
    user, err := s.repo.getByEmail(ctx, req.Email)
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
