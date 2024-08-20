package credit

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/usuario/credit-api/pkg/response"
)

func RegisterRoutes(r *mux.Router) {
    r.HandleFunc("/applications", GetApplicationsHandler).Methods("GET")
    r.HandleFunc("/applications/{id}", GetApplicationHandler).Methods("GET")
    r.HandleFunc("/applications", CreateApplicationHandler).Methods("POST")
    r.HandleFunc("/applications/{id}/approve", ApproveApplicationHandler).Methods("POST")
    r.HandleFunc("/applications/{id}/reject", RejectApplicationHandler).Methods("POST")
}

func GetApplicationsHandler(w http.ResponseWriter, r *http.Request) {
    applications := GetAllApplicationsService()
    response.JSON(w, http.StatusOK, applications)
}

func GetApplicationHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    application := GetApplicationByIDService(id)
    if application == nil {
        response.Error(w, http.StatusNotFound, "Application not found")
        return
    }
    response.JSON(w, http.StatusOK, application)
}

func CreateApplicationHandler(w http.ResponseWriter, r *http.Request) {
    var application CreditApplication
    if err := json.NewDecoder(r.Body).Decode(&application); err != nil {
        response.Error(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    createdApplication := CreateApplicationService(application)
    response.JSON(w, http.StatusCreated, createdApplication)
}

func ApproveApplicationHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var payload struct {
        ApprovedAmount float64 `json:"approved_amount"`
    }
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        response.Error(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    application := ApproveApplicationService(id, payload.ApprovedAmount)
    if application == nil {
        response.Error(w, http.StatusNotFound, "Application not found")
        return
    }
    response.JSON(w, http.StatusOK, application)
}

func RejectApplicationHandler(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    application := RejectApplicationService(id)
    if application == nil {
        response.Error(w, http.StatusNotFound, "Application not found")
        return
    }
    response.JSON(w, http.StatusOK, application)
}
