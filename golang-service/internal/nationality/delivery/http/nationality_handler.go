package http

import (
	"net/http"

	"github.com/ChristianTertius/backend_developer_test/internal/domain"
	"github.com/gorilla/mux"
)

type NationalityHandler struct {
	usecase domain.NationalityUsecase
}

func NewNationalityHandler(r *mux.Router, uc domain.NationalityUsecase) {
	h := &NationalityHandler{usecase: uc}
	r.HandleFunc("/nationalities", h.Fetch).Methods(http.MethodGet)
}

func (h *NationalityHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	data, err := h.usecase.Fetch(r.Context())
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondData(w, http.StatusOK, data)
}
