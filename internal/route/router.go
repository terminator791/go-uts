package route

import (
	"net/http"
	"strings"

	"go-uts/internal/handler"
)

func NewRouter(
	anggotaHandler *handler.AnggotaHandler,
	peminjamanHandler *handler.PeminjamanHandler,
	pembayaranHandler *handler.PembayaranHandler,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/anggota", anggotaHandler.HandleCollection)
	mux.HandleFunc("/api/anggota/", anggotaHandler.HandleItem)

	mux.HandleFunc("/api/peminjaman", peminjamanHandler.HandleCollection)
	mux.HandleFunc("/api/peminjaman/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/pembayaran") {
			pembayaranHandler.HandleByPeminjaman(w, r)
			return
		}
		peminjamanHandler.HandleItem(w, r)
	})

	mux.HandleFunc("/api/pembayaran", pembayaranHandler.HandleCollection)
	mux.HandleFunc("/api/pembayaran/", pembayaranHandler.HandleItem)

	return mux
}
