package employees

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

func TestGetEmployees(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/employees", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/v1/employees", GetAllEmployees)

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

	var employees []Employee
	if err := json.Unmarshal(rr.Body.Bytes(), &employees); err != nil {
		t.Errorf("failed to parse response body: %v", err)
	} else {
		t.Log("Response body parsed successfully. Assertion passed.")
	}

	expectedLength := 8
	if len(employees) != expectedLength {
		t.Errorf("expected %d employees, but got %d", expectedLength, len(employees))
	} else {
		t.Logf("Employees list contains %d elements. Assertion passed.", expectedLength)
	}

	t.Log("GetEmployees test passed!")
}

func TestGetEmployeeByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/employees/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/v1/employees/{emplID}", GetEmployeeByID)

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

	expected := `{"id":1,"name":"Jane","surname":"Doe","role":"Manager","phone":"555-1234","email":"john.doe@example.com"}`
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	} else {
		t.Logf("Response body is %s. Assertion passed.", expected)
	}
}

func TestGetEmployeeByIDNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/employees/12345", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/v1/employees/{emplID}", GetEmployeeByID)

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
	expected := `{"error_message":"Employee with id 12345 not found"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	} else {
		t.Logf("Response body is %s. Assertion passed.", expected)
	}
}
