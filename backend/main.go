package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/defilippomattia/fakeiotapi/backend/common"
	"github.com/defilippomattia/fakeiotapi/backend/database"
	"github.com/defilippomattia/fakeiotapi/backend/employees"
	"github.com/defilippomattia/fakeiotapi/backend/farms"
	"github.com/defilippomattia/fakeiotapi/backend/fields"
	"github.com/defilippomattia/fakeiotapi/backend/sensors"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	var db *sql.DB

	var err error
	db, err = database.Connect()
	if err != nil {
		common.WriteErrorJSONResponse(nil, http.StatusInternalServerError, err.Error())
		return
	}

	defer db.Close()

	handlerDeps := common.HandlerDependencies{
		DB:       db,
		StrDummy: "str_dummy",
		IntDummy: 42,
	}
	r := chi.NewRouter()
	r.Use(CorsMiddleware)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "dependencies", handlerDeps)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	})

	// Register farms endpoints
	r.Get("/api/v1/farms", farms.GetFarms)
	r.Get("/api/v1/farms/{farmID}", farms.GetFarmByID)

	r.Get("/api/v1/employees", employees.GetAllEmployees)
	r.Get("/api/v1/employees/{emplID}", employees.GetEmployeeByID)
	r.Get("/api/v1/farms/{farmID}/employees", employees.GetEmployeesByFarmID)

	r.Get("/api/v1/fields", fields.GetAllFields)
	r.Get("/api/v1/fields/{fieldID}", fields.GetFieldByID)
	r.Get("/api/v1/farms/{farmID}/fields", fields.GetFieldsByFarmID)

	r.Get("/api/v1/sensors", sensors.GetAllSensors)
	r.Get("/api/v1/sensors/{sensorID}", sensors.GetSensorByID)
	r.Get("/api/v1/farms/{farmID}/sensors", sensors.GetSensorsByFarmID)
	r.Get("/api/v1/fields/{fieldID}/sensors", sensors.GetSensorsByFieldID)

	r.Get("/api/v1/sensors/{sensorID}/temperature/current", sensors.GetCurrentTemperature)
	r.Get("/api/v1/sensors/{sensorID}/temperature/history", sensors.GetTemperatureHistory)

	r.Get("/api/v1/sensors/{sensorID}/humidity/current", sensors.GetCurrentHumidity)
	r.Get("/api/v1/sensors/{sensorID}/humidity/history", sensors.GetHumidityHistory)

	r.Get("/api/v1/sensors/{sensorID}/pressure/current", sensors.GetCurrentPressure)
	r.Get("/api/v1/sensors/{sensorID}/pressure/history", sensors.GetPressureHistory)

	http.ListenAndServe(":8080", r)

}
