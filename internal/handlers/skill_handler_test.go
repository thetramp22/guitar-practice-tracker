package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/thetramp22/rifflog/internal/bootstrap"
	"github.com/thetramp22/rifflog/internal/database"
	"github.com/thetramp22/rifflog/internal/models"
	"github.com/thetramp22/rifflog/internal/repository"
	"github.com/thetramp22/rifflog/internal/services"
)

func testSetup(t *testing.T) (*gin.Engine, *pgx.Conn) {
	t.Log("starting setup")
	err := godotenv.Load("../../.env.test")
	if err != nil {
		t.Log("No .env file found")
	}

	t.Log(os.Getwd())

	t.Log("connecting to database")
	conn := database.NewConnection()

	router := gin.Default()

	skillRepo := repository.NewSkillRepository(conn)
	skillService := services.NewSkillService(skillRepo)
	skillHandler := NewSkillHandler(skillService)

	t.Log("seeding skills")
	bootstrap.PopulateSkillsList(skillRepo)

	router.GET("/skills", skillHandler.ListSkills)

	return router, conn
}

func TestSkillsEndpoint(t *testing.T) {
	t.Log("creating router")
	router, conn := testSetup(t)
	defer conn.Close(context.Background())

	t.Log("creating request")
	req := httptest.NewRequest("GET", "http://localhost:8080/skills", nil)
	rec := httptest.NewRecorder()

	t.Log("ServeHTTP call")
	router.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("expected 200, got %v", status)
	}

	want := []models.Skill{
		{
			Name:        "Ear Training",
			Description: "Try playing to identify chords and melodies by ear.",
		},
		{
			Name:        "Scales",
			Description: "Memorize note locations and scale patterns.",
		},
		{
			Name:        "Timing and Rhythm",
			Description: "Practice with a metronome to develop a solid sense of time and groove.",
		},
	}

	var got []models.Skill
	err := json.Unmarshal(rec.Body.Bytes(), &got)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	opts := cmpopts.IgnoreFields(models.Skill{}, "ID", "CreatedAt")
	if diff := cmp.Diff(want, got, opts); diff != "" {
		t.Errorf("Values mismatch (-want +got):\n%s", diff)
	}
}
