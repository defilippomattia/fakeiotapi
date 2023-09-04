package farms

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/defilippomattia/fakeiotapi/backend/common"
	"github.com/defilippomattia/fakeiotapi/backend/database"
	"github.com/go-chi/chi"
)

func TestGetFarms(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/farms", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/v1/farms", GetFarms)

	var db *sql.DB
	db, err = database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handlerDeps := common.HandlerDependencies{
		DB:       db,
		StrDummy: "str_dummy",
		IntDummy: 42,
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx, "dependencies", handlerDeps)
	req = req.WithContext(ctx)

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, rr.Code)
	} else {
		t.Logf("Response status code is %d. Assertion passed.", http.StatusOK)
	}

	var farms []Farm
	if err := json.Unmarshal(rr.Body.Bytes(), &farms); err != nil {
		t.Errorf("failed to parse response body: %v", err)
	} else {
		t.Log("Response body parsed successfully. Assertion passed.")
	}

	expectedLength := 3
	if len(farms) != expectedLength {
		t.Errorf("expected %d farms, but got %d", expectedLength, len(farms))
	} else {
		t.Logf("Farms list contains %d elements. Assertion passed.", expectedLength)
	}

	t.Log("GetFarms test passed!")
}

func TestGetFarmByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/farms/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/v1/farms/{farmID}", GetFarmByID)

	var db *sql.DB
	db, err = database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handlerDeps := common.HandlerDependencies{
		DB:       db,
		StrDummy: "str_dummy",
		IntDummy: 42,
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx, "dependencies", handlerDeps)
	req = req.WithContext(ctx)

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, rr.Code)
	} else {
		t.Logf("Response status code is %d. Assertion passed.", http.StatusOK)
	}

	expected := `{"id":1,"name":"Jonhson Familiy Farm","location":"123 Main St","date_established":"2022-01-15T00:00:00Z","latitude":"40.123456","longitude":"-75.654321"}`
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	} else {
		t.Logf("Response body is %s. Assertion passed.", expected)
	}

}

func TestGetFarmByIDNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/farms/12345", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/v1/farms/{farmID}", GetFarmByID)

	var db *sql.DB
	db, err = database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handlerDeps := common.HandlerDependencies{
		DB:       db,
		StrDummy: "str_dummy",
		IntDummy: 42,
	}

	ctx := req.Context()
	ctx = context.WithValue(ctx, "dependencies", handlerDeps)
	req = req.WithContext(ctx)

	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status %d but got %d", http.StatusNotFound, rr.Code)
	} else {
		t.Logf("Response status code is %d. Assertion passed.", http.StatusOK)
	}

	//check the response body is what we expect.
	expected := `{"error_message":"Farm with id 12345 not found"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	} else {
		t.Logf("Response body is %s. Assertion passed.", expected)
	}

}
