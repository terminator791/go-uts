package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"go-uts/internal/domain"
	"go-uts/internal/repository"
	"go-uts/internal/usecase"
)

type AnggotaHandler struct {
	usecase usecase.AnggotaUsecase
}

func NewAnggotaHandler(usecase usecase.AnggotaUsecase) *AnggotaHandler {
	return &AnggotaHandler{usecase: usecase}
}

func (h *AnggotaHandler) HandleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	case http.MethodGet:
		h.getAll(w, r)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "method not allowed"})
	}
}

func (h *AnggotaHandler) HandleItem(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getByID(w, r, id)
	case http.MethodPut:
		h.update(w, r, id)
	case http.MethodDelete:
		h.delete(w, r, id)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "method not allowed"})
	}
}

func (h *AnggotaHandler) create(w http.ResponseWriter, r *http.Request) {
	var payload domain.Anggota
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"message": "request body tidak valid"})
		return
	}

	result, err := h.usecase.Create(r.Context(), payload)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "Data anggota berhasil ditambahkan",
		"data":    result,
	})
}

func (h *AnggotaHandler) getAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.usecase.GetAll(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"message": "gagal mengambil data"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{"data": result})
}

func (h *AnggotaHandler) getByID(w http.ResponseWriter, r *http.Request, id int64) {
	result, err := h.usecase.GetByID(r.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, repository.ErrNotFound) {
			status = http.StatusNotFound
		}
		writeJSON(w, status, map[string]string{"message": "data tidak ditemukan"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{"data": result})
}

func (h *AnggotaHandler) update(w http.ResponseWriter, r *http.Request, id int64) {
	var payload domain.Anggota
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"message": "request body tidak valid"})
		return
	}

	result, err := h.usecase.Update(r.Context(), id, payload)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, repository.ErrNotFound) {
			status = http.StatusNotFound
		} else {
			status = http.StatusBadRequest
		}
		writeJSON(w, status, map[string]string{"message": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Data anggota berhasil diperbarui",
		"data":    result,
	})
}

func (h *AnggotaHandler) delete(w http.ResponseWriter, r *http.Request, id int64) {
	if err := h.usecase.Delete(r.Context(), id); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, repository.ErrNotFound) {
			status = http.StatusNotFound
		}
		writeJSON(w, status, map[string]string{"message": "data tidak ditemukan"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Data anggota berhasil dihapus"})
}
