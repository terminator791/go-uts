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

type PembayaranHandler struct {
	usecase usecase.PembayaranUsecase
}

func NewPembayaranHandler(usecase usecase.PembayaranUsecase) *PembayaranHandler {
	return &PembayaranHandler{usecase: usecase}
}

func (h *PembayaranHandler) HandleCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	case http.MethodGet:
		h.getAll(w, r)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "method not allowed"})
	}
}

func (h *PembayaranHandler) HandleItem(w http.ResponseWriter, r *http.Request) {
	id, err := parsePembayaranID(r.URL.Path)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"message": "id tidak valid"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getByID(w, r, id)
	case http.MethodDelete:
		h.delete(w, r, id)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "method not allowed"})
	}
}

func (h *PembayaranHandler) HandleByPeminjaman(w http.ResponseWriter, r *http.Request) {
	peminjamanID, err := parsePeminjamanFromPaymentPath(r.URL.Path)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"message": "id peminjaman tidak valid"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getByPeminjamanID(w, r, peminjamanID)
	default:
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "method not allowed"})
	}
}

func (h *PembayaranHandler) create(w http.ResponseWriter, r *http.Request) {
	var payload domain.Pembayaran
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
		"message": "Data pembayaran berhasil ditambahkan",
		"data":    result,
	})
}

func (h *PembayaranHandler) getAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.usecase.GetAll(r.Context())
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"message": "gagal mengambil data"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{"data": result})
}

func (h *PembayaranHandler) getByID(w http.ResponseWriter, r *http.Request, id int64) {
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

func (h *PembayaranHandler) getByPeminjamanID(w http.ResponseWriter, r *http.Request, peminjamanID int64) {
	result, err := h.usecase.GetByPeminjamanID(r.Context(), peminjamanID)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"message": "gagal mengambil data"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{"data": result})
}

func (h *PembayaranHandler) delete(w http.ResponseWriter, r *http.Request, id int64) {
	result, err := h.usecase.Delete(r.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, repository.ErrNotFound) {
			status = http.StatusNotFound
		}
		writeJSON(w, status, map[string]string{"message": "data tidak ditemukan"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Data pembayaran berhasil dihapus",
		"data":    result,
	})
}

func parsePembayaranID(path string) (int64, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 3 {
		return 0, errors.New("id tidak ada")
	}
	return strconv.ParseInt(parts[2], 10, 64)
}

func parsePeminjamanFromPaymentPath(path string) (int64, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 4 {
		return 0, errors.New("id tidak ada")
	}
	return strconv.ParseInt(parts[2], 10, 64)
}
