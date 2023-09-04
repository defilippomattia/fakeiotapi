package employees

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/defilippomattia/fakeiotapi/backend/common"

	"github.com/go-chi/chi"
)

type Employee struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Role    string `json:"role"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)

	rows, err := handlerDeps.DB.Query("SELECT id, name, surname, role, phone, email FROM employee")
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()
	employees := []Employee{}
	for rows.Next() {
		var employee Employee
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Surname, &employee.Role, &employee.Phone, &employee.Email)
		if err != nil {
			common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		employees = append(employees, employee)
	}
	common.WriteJSONResponse(w, http.StatusOK, employees)
}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	handlerDeps := r.Context().Value("dependencies").(common.HandlerDependencies)

	id := chi.URLParam(r, "emplID")

	var employee Employee
	row := handlerDeps.DB.QueryRow("SELECT id, name, surname, role, phone, email FROM employee WHERE id = $1", id)
	err := row.Scan(&employee.ID, &employee.Name, &employee.Surname, &employee.Role, &employee.Phone, &employee.Email)
	if err == sql.ErrNoRows {
		errMsg := fmt.Sprintf("Employee with id %s not found", id)
		common.WriteErrorJSONResponse(w, http.StatusNotFound, errMsg)
		return
	} else if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	common.WriteJSONResponse(w, http.StatusOK, employee)
}

func GetEmployeesByFarmID(w http.ResponseWriter, r *http.Request) {
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
	query := `
	SELECT e.id, e.name, e.surname, e.role, e.phone, e.email
	FROM EMPLOYEE e
	INNER JOIN farm2employee f2e ON e.id = f2e.employee_id
	WHERE f2e.farm_id = $1`

	rows, err := handlerDeps.DB.Query(query, farmID)
	if err != nil {
		common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	employees := []Employee{}
	for rows.Next() {
		var employee Employee
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Surname, &employee.Role, &employee.Phone, &employee.Email)
		if err != nil {
			common.WriteErrorJSONResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
		employees = append(employees, employee)
	}

	common.WriteJSONResponse(w, http.StatusOK, employees)
}
