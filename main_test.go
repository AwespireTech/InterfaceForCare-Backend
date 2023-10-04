package main

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/AwespireTech/InterfaceForCare-Backend/database"
	"github.com/AwespireTech/InterfaceForCare-Backend/models"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	err := database.Init("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	gin.SetMode(gin.TestMode)
	m.Run()
}
func getRouter(t *testing.T) *gin.Engine {
	t.Helper()
	return createRouter()
}

func TestDefaultRoute(t *testing.T) {
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}
	router := getRouter(t)
	router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}

}

func TestGetRivers(t *testing.T) {
	recorder := httptest.NewRecorder()

	request, err := http.NewRequest("GET", "/api/rivers/18", nil)
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}

	router := getRouter(t)

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}

func TestCreateRiver(t *testing.T) {

	body := []byte(`{
		"name": "test",
		"agreement": "test",
	}`)

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("POST", "/api/rivers", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}

	router := getRouter(t)

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Errorf("expected status code %d, got %d", http.StatusCreated, recorder.Code)
	}
}

func TestUpdateRiver(t *testing.T) {
	body := []byte(`{
		"agreement": "test",
		"gen": 0,
	}`)

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest("PUT", "/api/rivers/18", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}

	router := getRouter(t)

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}
func TestJoinRiverById(t *testing.T) {
	recorder := httptest.NewRecorder()

	request, err := http.NewRequest("POST", "/api/river/join/18", nil)
	if err != nil {
		t.Fatalf("failed to create mock request: %v", err)
	}

	router := getRouter(t)

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, recorder.Code)
	}
}
func RandomRiver(t *testing.T) models.River {
	t.Helper()
	return models.River{
		ID:            rand.Int(),
		Name:          "test",
		Agreement:     "test",
		Generation:    0,
		CreatedTime:   time.Now(),
		ExpiredTime:   time.Now().Add(time.Hour * 24 * 7),
		Status:        rand.Intn(2),
		Owners:        []string{"test"},
		OwnersCount:   1,
		TokenId:       0,
		WalletAddress: "test",
		EventData:     []models.Event{},
	}
}
func RandomEvent(t *testing.T, rid int) models.Event {
	t.Helper()
	return models.Event{
		ID:                strconv.Itoa(rid) + "-" + strconv.Itoa(rand.Int()),
		Name:              "test",
		ImageUri:          "test",
		Description:       "test",
		Editions:          1,
		Status:            rand.Intn(2),
		Participants:      []string{"test"},
		ParticipantsCount: 1,
		Approvals:         []string{"test"},
		ApprovalsCount:    1,
	}
}

