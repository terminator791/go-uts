package route

import (
	"net/http"

	"go-uts/internal/handler"
)

func NewRouter(peminjamanHandler *handler.PeminjamanHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/peminjaman", peminjamanHandler.HandleCollection)
	mux.HandleFunc("/api/peminjaman/", peminjamanHandler.HandleItem)

	return mux
}
