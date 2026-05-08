package bootstrap

import (
	"github.com/thetramp22/rifflog/internal/models"
	"github.com/thetramp22/rifflog/internal/repository"
)

func PopulateSkillsList(r *repository.SkillRepository) error {
	skills := []models.Skill{
		{
			Name:        "Scales",
			Description: "Memorize note locations and scale patterns.",
		},
		{
			Name:        "Timing and Rhythm",
			Description: "Practice with a metronome to develop a solid sense of time and groove.",
		},
		{
			Name:        "Ear Training",
			Description: "Try playing to identify chords and melodies by ear.",
		},
	}

	for _, skill := range skills {
		err := r.SeedSkill(skill)
		if err != nil {
			return err
		}
	}

	return nil
}
