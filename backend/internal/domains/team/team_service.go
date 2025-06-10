package team

import (
	"fmt"

	"github.com/google/uuid"
)

/*********************
 * SERVICE INTERFACE *
 *********************/

type TeamService interface {
	CreateTeam(team *Team) error
	GetTeam(id uuid.UUID) (*Team, error)
	ListTeams(userID uuid.UUID) ([]Team, error)
	UpdateTeam(team *Team) error
	DeleteTeam(id uuid.UUID) error
}

/**************************
 * SERVICE IMPLEMENTATION *
 **************************/

type teamService struct {
	repo TeamRepository
}

func NewTeamService(repo TeamRepository) TeamService {
	return &teamService{repo}
}

func (s *teamService) CreateTeam(team *Team) error {
	if err := team.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return s.repo.Create(team)
}

func (s *teamService) GetTeam(id uuid.UUID) (*Team, error) {
	return s.repo.GetByID(id)
}

func (s *teamService) ListTeams(userID uuid.UUID) ([]Team, error) {
	return s.repo.ListByUser(userID)
}

func (s *teamService) UpdateTeam(team *Team) error {
	if err := team.Validate(); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return s.repo.Update(team)
}

func (s *teamService) DeleteTeam(id uuid.UUID) error {
	return s.repo.Delete(id)
}
