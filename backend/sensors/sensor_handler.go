package sensors

import (
	"database/sql"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/defilippomattia/fakeiotapi/backend/common"
	"github.com/go-chi/chi"
)

type Sensor struct {
	ID           int    `json:"id"`
	Manufacturer string `json:"manufacturer"`
	SensorType   string `json:"type"`
	Name         string `json:"name"`
	FieldID      int    `json:"field_id"`
}

type Temperature struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
	SensorID  int     `json:"sensor_id"`
}

type Humidity struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
	SensorID  int     `json:"sensor_id"`
}

type Pressure struct {
	Timestamp string  `json:"timestamp"`
	Value     float64 `json:"value"`
	SensorID  int     `json:"sensor_id"`
}

type ErrorResponse struct {
	ErrorMsg string `json:"error_msg"`
}

func GetAllSensors(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)
	rows, err := handlerDeps.DB.Query("SELECT id, manufacturer, type, name, field_id FROM sensor")
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	sensors := []Sensor{}
	for rows.Next() {
		var sensor Sensor
		err := rows.Scan(&sensor.ID, &sensor.Manufacturer, &sensor.SensorType, &sensor.Name, &sensor.FieldID)
		if err != nil {
			common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		sensors = append(sensors, sensor)
	}
	common.WriteJSONResponse(w, http.StatusOK, sensors)
}

func GetSensorByID(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)

	id := chi.URLParam(r, "sensorID")

	var sensor Sensor
	row := handlerDeps.DB.QueryRow("SELECT id, manufacturer, type, name, field_id FROM sensor WHERE id = $1", id)
	err := row.Scan(&sensor.ID, &sensor.Manufacturer, &sensor.SensorType, &sensor.Name, &sensor.FieldID)
	if err == sql.ErrNoRows {
		errMsg := fmt.Sprintf("Sensor with ID %s not found", id)
		common.WriteErrorJSONResponse(w, http.StatusNotFound, errMsg)
		return
	} else if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSONResponse(w, http.StatusOK, sensor)
}

func GetSensorsByFieldID(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)

	id := chi.URLParam(r, "fieldID")

	rows, err := handlerDeps.DB.Query("SELECT id, manufacturer, type, name, field_id FROM sensor WHERE field_id = $1", id)
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	sensors := []Sensor{}
	for rows.Next() {
		var sensor Sensor
		err := rows.Scan(&sensor.ID, &sensor.Manufacturer, &sensor.SensorType, &sensor.Name, &sensor.FieldID)
		if err != nil {
			common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		sensors = append(sensors, sensor)
	}
	common.WriteJSONResponse(w, http.StatusOK, sensors)
}

func GetSensorsByFarmID(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)

	id := chi.URLParam(r, "farmID")

	rows, err := handlerDeps.DB.Query("SELECT id, manufacturer, type, name, field_id FROM sensor WHERE field_id IN (SELECT id FROM field WHERE farm_id = $1)", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	sensors := []Sensor{}
	for rows.Next() {
		var sensor Sensor
		err := rows.Scan(&sensor.ID, &sensor.Manufacturer, &sensor.SensorType, &sensor.Name, &sensor.FieldID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		sensors = append(sensors, sensor)
	}
	common.WriteJSONResponse(w, http.StatusOK, sensors)
}

func GetTemperatureHistory(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)
	sensorID := chi.URLParam(r, "sensorID")
	rows, err := handlerDeps.DB.Query("SELECT timestamp, value, sensor_id FROM temperature WHERE sensor_id=$1", sensorID)
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	temperatures := []Temperature{}
	for rows.Next() {
		var temperature Temperature
		err := rows.Scan(&temperature.Timestamp, &temperature.Value, &temperature.SensorID)
		if err != nil {
			common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		parsedTime, _ := time.Parse("2006-01-02T15:04:05Z", temperature.Timestamp)
		temperature.Timestamp = parsedTime.Format("2006-01-02 15:04:05")
		temperatures = append(temperatures, temperature)
	}
	common.WriteJSONResponse(w, http.StatusOK, temperatures)

}

func GetHumidityHistory(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)
	sensorID := chi.URLParam(r, "sensorID")
	rows, err := handlerDeps.DB.Query("SELECT timestamp, value, sensor_id FROM humidity WHERE sensor_id=$1", sensorID)
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	humidities := []Humidity{}
	for rows.Next() {
		var humidity Humidity
		err := rows.Scan(&humidity.Timestamp, &humidity.Value, &humidity.SensorID)
		if err != nil {
			common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		parsedTime, _ := time.Parse("2006-01-02T15:04:05Z", humidity.Timestamp)
		humidity.Timestamp = parsedTime.Format("2006-01-02 15:04:05")
		humidities = append(humidities, humidity)
	}
	common.WriteJSONResponse(w, http.StatusOK, humidities)

}

func GetPressureHistory(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)
	sensorID := chi.URLParam(r, "sensorID")
	rows, err := handlerDeps.DB.Query("SELECT timestamp, value, sensor_id FROM pressure WHERE sensor_id=$1", sensorID)
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	pressures := []Pressure{}
	for rows.Next() {
		var pressure Pressure
		err := rows.Scan(&pressure.Timestamp, &pressure.Value, &pressure.SensorID)
		if err != nil {
			common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		parsedTime, _ := time.Parse("2006-01-02T15:04:05Z", pressure.Timestamp)
		pressure.Timestamp = parsedTime.Format("2006-01-02 15:04:05")
		pressures = append(pressures, pressure)
	}
	common.WriteJSONResponse(w, http.StatusOK, pressures)

}

func GetCurrentTemperature(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)

	id := chi.URLParam(r, "sensorID")

	var sensorType string
	err := handlerDeps.DB.QueryRow("SELECT type FROM sensor WHERE id = $1", id).Scan(&sensorType)
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if strings.ToLower(sensorType) != "temperature" {
		errorMsg := fmt.Sprintf("Sensor with ID %s is not a Temperature sensor. It is a %s sensor.", id, sensorType)
		common.WriteErrorJSONResponse(w, http.StatusBadRequest, errorMsg)
		return
	}

	currentTimestamp := getCurrentTimestamp()
	currentTemperature := generateRandomNumber(20, 23)
	idAsInt, _ := strconv.Atoi(id) //todo check error
	temperature := Temperature{
		Timestamp: currentTimestamp,
		Value:     currentTemperature,
		SensorID:  idAsInt,
	}

	common.WriteJSONResponse(w, http.StatusOK, temperature)
}

func GetCurrentHumidity(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)

	id := chi.URLParam(r, "sensorID")

	var sensorType string
	err := handlerDeps.DB.QueryRow("SELECT type FROM sensor WHERE id = $1", id).Scan(&sensorType)
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if strings.ToLower(sensorType) != "humidity" {
		errorMsg := fmt.Sprintf("Sensor with ID %s is not a Humidity sensor. It is a %s sensor.", id, sensorType)
		common.WriteErrorJSONResponse(w, http.StatusBadRequest, errorMsg)
		return
	}

	currentTimestamp := getCurrentTimestamp()
	currentHumidity := generateRandomNumber(58, 61)
	idAsInt, _ := strconv.Atoi(id) //todo check error
	humidity := Humidity{
		Timestamp: currentTimestamp,
		Value:     currentHumidity,
		SensorID:  idAsInt,
	}

	common.WriteJSONResponse(w, http.StatusOK, humidity)
}

func GetCurrentPressure(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)

	id := chi.URLParam(r, "sensorID")

	var sensorType string
	err := handlerDeps.DB.QueryRow("SELECT type FROM sensor WHERE id = $1", id).Scan(&sensorType)
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if strings.ToLower(sensorType) != "pressure" {
		errorMsg := fmt.Sprintf("Sensor with ID %s is not a Pressure sensor. It is a %s sensor.", id, sensorType)
		common.WriteErrorJSONResponse(w, http.StatusBadRequest, errorMsg)
		return
	}

	currentTimestamp := getCurrentTimestamp()
	currentPressure := generateRandomNumber(1010, 1013)
	idAsInt, _ := strconv.Atoi(id) //todo check error

	pressure := Pressure{
		Timestamp: currentTimestamp,
		Value:     currentPressure,
		SensorID:  idAsInt,
	}

	common.WriteJSONResponse(w, http.StatusOK, pressure)
}

func generateRandomNumber(min, max int) float64 {
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator with the current time

	randomFloat := rand.Float64()
	randomNumber := float64(max-min)*randomFloat + float64(min)
	randomNumber = math.Round(randomNumber*100) / 100
	return randomNumber
}

func getCurrentTimestamp() string {
	now := time.Now()
	layout := "2006-01-02 15:04:05"
	return now.Format(layout)
}
