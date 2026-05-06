package domain

type Peminjaman struct {
	ID            int64  `json:"id"`
	IDAnggota     int64  `json:"id_anggota"`
	TanggalPinjam string `json:"tanggal_pinjam"`
	JumlahPinjam  int64  `json:"jumlah_pinjam"`
	Status        string `json:"status"`
}
