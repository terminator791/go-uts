package domain

type Pembayaran struct {
	ID               int64  `json:"id"`
	IDPeminjaman     int64  `json:"id_peminjaman"`
	TanggalPembayaran string `json:"tanggal_pembayaran"`
	JumlahPembayaran int64  `json:"jumlah_pembayaran"`
}
