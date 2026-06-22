package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/ChristianTertius/backend_developer_test/internal/domain"
	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	usecase domain.CustomerUsecase
}

func NewCustomerHandler(r *mux.Router, uc domain.CustomerUsecase) {
	h := &CustomerHandler{usecase: uc}
	r.HandleFunc("/customers", h.Fetch).Methods(http.MethodGet)
	r.HandleFunc("/customers", h.Store).Methods(http.MethodPost)
	r.HandleFunc("/customers/{id:[0-9]+}", h.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/customers/{id:[0-9]+}", h.Update).Methods(http.MethodPut)
	r.HandleFunc("/customers/{id:[0-9]+}", h.Delete).Methods(http.MethodDelete)
}

func (h *CustomerHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	customers, err := h.usecase.Fetch(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondData(w, http.StatusOK, customers)
}

func (h *CustomerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "id not found!")
		return
	}
	customer, err := h.usecase.GetByID(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}
	respondData(w, http.StatusOK, customer)
}

func (h *CustomerHandler) Store(w http.ResponseWriter, r *http.Request) {
	var c domain.Customer
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		respondError(w, http.StatusBadRequest, "body json tidak valid")
		return
	}
	if err := h.usecase.Store(r.Context(), &c); err != nil {
		handleError(w, err)
		return
	}
	respondData(w, http.StatusCreated, c)
}

func (h *CustomerHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "id not found!")
		return
	}

	var c domain.Customer
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	c.ID = id
	if err := h.usecase.Update(r.Context(), &c); err != nil {
		handleError(w, err)
		return
	}

	respondData(w, http.StatusOK, c)

}

func (h *CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.usecase.Delete(r.Context(), id); err != nil {
		handleError(w, err)
		return
	}

	respondMessage(w, http.StatusOK, "successfully delete customer!")
}

func parseID(r *http.Request) (int64, error) {
	return strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		respondError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, domain.ErrBadRequest):
		respondError(w, http.StatusBadRequest, err.Error())
	default:
		respondError(w, http.StatusInternalServerError, err.Error())
	}
}
