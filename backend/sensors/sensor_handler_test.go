package sensors

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

func TestGetTemparatureHistory(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/sensors/1/temperature/history", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/v1/sensors/{sensorID}/temperature/history", GetTemperatureHistory)

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

	var temperatures []Temperature
	if err := json.Unmarshal(rr.Body.Bytes(), &temperatures); err != nil {
		t.Errorf("failed to parse response body: %v", err)
	} else {
		t.Log("Response body parsed successfully. Assertion passed.")
	}
	expectedLength := 5
	if len(temperatures) != expectedLength {
		t.Errorf("expected length %d but got %d", expectedLength, len(temperatures))
	} else {
		t.Logf("Response body length is %d. Assertion passed.", expectedLength)
	}

	expectedTemperature := Temperature{
		Timestamp: "2023-07-01 08:00:00",
		Value:     25.5,
		SensorID:  1,
	}
	gotTemperature := temperatures[0]
	if expectedTemperature != gotTemperature {
		t.Errorf("expected temperature %+v but got %+v", expectedTemperature, gotTemperature)
	} else {
		t.Logf("Response body is %+v. Assertion passed.", gotTemperature)
	}

}

func TestGetHumidityHistory(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/sensors/2/humidity/history", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/v1/sensors/{sensorID}/humidity/history", GetHumidityHistory)

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

	var humidities []Humidity
	if err := json.Unmarshal(rr.Body.Bytes(), &humidities); err != nil {
		t.Errorf("failed to parse response body: %v", err)
	} else {
		t.Log("Response body parsed successfully. Assertion passed.")
	}
	expectedLength := 5
	if len(humidities) != expectedLength {
		t.Errorf("expected length %d but got %d", expectedLength, len(humidities))
	} else {
		t.Logf("Response body length is %d. Assertion passed.", expectedLength)
	}
	expectedHumidity := Humidity{
		Timestamp: "2023-07-01 08:00:00",
		Value:     60.2,
		SensorID:  2,
	}
	gotHumidity := humidities[0]
	if expectedHumidity != gotHumidity {
		t.Errorf("expected humidity %+v but got %+v", expectedHumidity, gotHumidity)
	} else {
		t.Logf("Response body is %+v. Assertion passed.", gotHumidity)
	}

}

func TestGetPressureHistory(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/sensors/3/pressure/history", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Get("/api/v1/sensors/{sensorID}/pressure/history", GetPressureHistory)

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

	var pressures []Pressure
	if err := json.Unmarshal(rr.Body.Bytes(), &pressures); err != nil {
		t.Errorf("failed to parse response body: %v", err)
	} else {
		t.Log("Response body parsed successfully. Assertion passed.")
	}
	expectedLength := 5
	if len(pressures) != expectedLength {
		t.Errorf("expected length %d but got %d", expectedLength, len(pressures))
	} else {
		t.Logf("Response body length is %d. Assertion passed.", expectedLength)
	}
	expectedPressure := Pressure{
		Timestamp: "2023-07-01 08:00:00",
		Value:     1013.2,
		SensorID:  3,
	}
	gotPressure := pressures[0]
	if expectedPressure != gotPressure {
		t.Errorf("expected pressure %+v but got %+v", expectedPressure, gotPressure)
	} else {
		t.Logf("Response body is %+v. Assertion passed.", gotPressure)
	}

}
