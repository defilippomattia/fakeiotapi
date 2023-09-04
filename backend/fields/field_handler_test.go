package fields

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/defilippomattia/fakeiotapi/backend/common"
	"github.com/defilippomattia/fakeiotapi/backend/database"
	"github.com/go-chi/chi"
)

func TestGetAllFields(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/fields", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/v1/fields", GetAllFields)

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
	var fields []Field
	if err := json.Unmarshal(rr.Body.Bytes(), &fields); err != nil {
		t.Errorf("failed to parse response body: %v", err)
	} else {
		t.Log("Response body parsed successfully. Assertion passed.")
	}
	expectedLength := 14
	if len(fields) != expectedLength {
		t.Errorf("expected length %d but got %d", expectedLength, len(fields))
	} else {
		t.Logf("Response body length is %d. Assertion passed.", expectedLength)
	}

}

func TestGetFieldByID(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/fields/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/v1/fields/{fieldID}", GetFieldByID)

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

	var expectedField Field
	var gotField Field

	if err := json.Unmarshal(rr.Body.Bytes(), &gotField); err != nil {
		t.Errorf("failed to parse response body: %v", err)
	} else {
		t.Log("Response body parsed successfully. Assertion passed.")
	}
	expectedField = Field{
		ID:        1,
		FieldType: "Corn",
		M2:        1000,
		FarmID:    "1",
	}
	if expectedField != gotField {
		t.Errorf("expected field %+v but got %+v", expectedField, gotField)
	} else {
		t.Logf("Response body is %+v. Assertion passed.", gotField)
	}
}

func TestGetFieldByIDNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/fields/12345", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/v1/fields/{fieldID}", GetFieldByID)

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
		t.Logf("Response status code is %d. Assertion passed.", http.StatusNotFound)
	}

	var expectedResponse common.ErrorResponse
	var gotResponse common.ErrorResponse

	if err := json.Unmarshal(rr.Body.Bytes(), &gotResponse); err != nil {
		t.Errorf("failed to parse response body: %v", err)
	} else {
		t.Log("Response body parsed successfully. Assertion passed.")
	}
	expectedResponse = common.ErrorResponse{
		ErrorMessage: "Field with id 12345 not found",
	}

	if expectedResponse != gotResponse {
		t.Errorf("expected response %+v but got %+v", expectedResponse, gotResponse)
	} else {
		t.Logf("Response body is %+v. Assertion passed.", gotResponse)
	}

}
