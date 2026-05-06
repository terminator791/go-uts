package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"go-uts/internal/domain"
	"go-uts/internal/repository"
	"go-uts/internal/usecase"
)

type PeminjamanHandler struct {
	usecase usecase.PeminjamanUsecase
}

func NewPeminjamanHandler(usecase usecase.PeminjamanUsecase) *PeminjamanHandler {
	return &PeminjamanHandler{usecase: usecase}
}

func (h *PeminjamanHandler) HandleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	case http.MethodGet:
		h.getAll(w, r)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "method not allowed"})
	}
}

func (h *PeminjamanHandler) HandleItem(w http.ResponseWriter, r *http.Request) {
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

func (h *PeminjamanHandler) create(w http.ResponseWriter, r *http.Request) {
	var payload domain.Peminjaman
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
		"message": "Data peminjaman berhasil ditambahkan",
		"data":    result,
	})
}

func (h *PeminjamanHandler) getAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.usecase.GetAll(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"message": "gagal mengambil data"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{"data": result})
}

func (h *PeminjamanHandler) getByID(w http.ResponseWriter, r *http.Request, id int64) {
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

func (h *PeminjamanHandler) update(w http.ResponseWriter, r *http.Request, id int64) {
	var payload struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"message": "request body tidak valid"})
		return
	}

	result, err := h.usecase.UpdateStatus(r.Context(), id, payload.Status)
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
		"message": "Data peminjaman berhasil diperbarui",
		"data":    result,
	})
}

func (h *PeminjamanHandler) delete(w http.ResponseWriter, r *http.Request, id int64) {
	if err := h.usecase.Delete(r.Context(), id); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, repository.ErrNotFound) {
			status = http.StatusNotFound
		}
		writeJSON(w, status, map[string]string{"message": "data tidak ditemukan"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": "Data peminjaman berhasil dihapus"})
}

func parseID(path string) (int64, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 3 {
		return 0, errors.New("id tidak ada")
	}
	return strconv.ParseInt(parts[2], 10, 64)
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
