package fields

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/defilippomattia/fakeiotapi/backend/common"
	"github.com/go-chi/chi"
)

type Field struct {
	ID        int     `json:"id"`
	FieldType string  `json:"type"`
	M2        float64 `json:"m2"`
	FarmID    string  `json:"farm_id"`
}

func GetAllFields(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)

	rows, err := handlerDeps.DB.Query("SELECT id, type, m2, farm_id FROM field")
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	fields := []Field{}
	for rows.Next() {
		var field Field
		err := rows.Scan(&field.ID, &field.FieldType, &field.M2, &field.FarmID)
		if err != nil {
			common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		fields = append(fields, field)
	}
	common.WriteJSONResponse(w, http.StatusOK, fields)
}

func GetFieldByID(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)

	id := chi.URLParam(r, "fieldID")

	var field Field
	row := handlerDeps.DB.QueryRow("SELECT id, type, m2, farm_id FROM field WHERE id = $1", id)
	err := row.Scan(&field.ID, &field.FieldType, &field.M2, &field.FarmID)
	if err == sql.ErrNoRows {
		errMsg := fmt.Sprintf("Field with id %s not found", id)
		common.WriteErrorJSONResponse(w, http.StatusNotFound, errMsg)
		return
	} else if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSONResponse(w, http.StatusOK, field)
}

func GetFieldsByFarmID(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)

	farmID := chi.URLParam(r, "farmID")
	row := handlerDeps.DB.QueryRow("SELECT id FROM farm WHERE id = $1", farmID)
	err := row.Scan(&farmID)
	if err == sql.ErrNoRows {
		errMsg := fmt.Sprintf("Farm with id %s not found", farmID)
		common.WriteErrorJSONResponse(w, http.StatusNotFound, errMsg)
		return
	} else if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	rows, err := handlerDeps.DB.Query("SELECT id, type, m2, farm_id FROM field WHERE farm_id = $1", farmID)
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	fields := []Field{}
	for rows.Next() {
		var field Field
		err := rows.Scan(&field.ID, &field.FieldType, &field.M2, &field.FarmID)
		if err != nil {
			common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		fields = append(fields, field)
	}
	common.WriteJSONResponse(w, http.StatusOK, fields)
}
