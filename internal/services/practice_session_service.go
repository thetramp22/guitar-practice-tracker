package services

import (
	"github.com/thetramp22/rifflog/internal/models"
	"github.com/thetramp22/rifflog/internal/repository"
)

type PracticeSessionService struct {
	Repo *repository.PracticeSessionRepository
}

func NewPracticeSessionService(repo *repository.PracticeSessionRepository) *PracticeSessionService {
	return &PracticeSessionService{Repo: repo}
}

func (s *PracticeSessionService) CreatePracticeSession(req models.SessionsRequest) error {
	practiceSession := models.PracticeSession{
		SkillID:         req.SkillID,
		DurationMinutes: req.DurationMinutes,
		PracticedAt:     req.PracticedAt,
		Notes:           req.Notes,
		UserID:          req.UserID,
	}

	return s.Repo.CreatePracticeSession(practiceSession)
}
