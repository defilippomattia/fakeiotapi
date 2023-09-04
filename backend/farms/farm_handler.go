package farms

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/defilippomattia/fakeiotapi/backend/common"
	"github.com/go-chi/chi"
)

type Farm struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Location        string `json:"location"`
	DateEstablished string `json:"date_established"`
	Latitude        string `json:"latitude"`
	Longitude       string `json:"longitude"`
}

func GetFarms(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)
	rows, err := handlerDeps.DB.Query("SELECT id, name, location, date_established, latitude, longitude FROM farm")
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	farms := []Farm{}
	for rows.Next() {
		var farm Farm
		err := rows.Scan(&farm.ID, &farm.Name, &farm.Location, &farm.DateEstablished, &farm.Latitude, &farm.Longitude)
		if err != nil {
			common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		farms = append(farms, farm)
	}
	common.WriteJSONResponse(w, http.StatusOK, farms)
}

func GetFarmByID(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)
	idStr := chi.URLParam(r, "farmID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	var farm Farm
	row := handlerDeps.DB.QueryRow("SELECT id, name, location, date_established, latitude, longitude FROM farm WHERE id = $1", id)
	err = row.Scan(&farm.ID, &farm.Name, &farm.Location, &farm.DateEstablished, &farm.Latitude, &farm.Longitude)
	if err == sql.ErrNoRows {
		errMsg := fmt.Sprintf("Farm with id %d not found", id)
		common.WriteErrorJSONResponse(w, http.StatusNotFound, errMsg)
		return
	} else if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSONResponse(w, http.StatusOK, farm)
}
