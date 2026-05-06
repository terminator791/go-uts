package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"go-uts/internal/config"
	"go-uts/internal/db"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()
	pool, err := db.NewPool(ctx, cfg)
	if err != nil {
		log.Fatalf("gagal konek database: %v", err)
	}
	defer pool.Close()

	if err := seedData(ctx, pool); err != nil {
		log.Fatalf("seed gagal: %v", err)
	}

	log.Println("seed selesai")
}

func seedData(ctx context.Context, pool *pgxpool.Pool) error {
	queries := []string{
		"TRUNCATE TABLE pembayaran, peminjaman, anggota RESTART IDENTITY CASCADE",
		"INSERT INTO anggota (nama, alamat, telepon) VALUES ('John Doe', 'Jalan Merdeka No. 123', '081234567890')",
		"INSERT INTO anggota (nama, alamat, telepon) VALUES ('Jane Smith', 'Jalan Gatot Subroto', '082345678901')",
		"INSERT INTO peminjaman (id_anggota, tanggal_pinjam, jumlah_pinjam, status) VALUES (1, '2024-04-10', 1000000, 'pending')",
		"INSERT INTO peminjaman (id_anggota, tanggal_pinjam, jumlah_pinjam, status) VALUES (2, '2024-04-12', 750000, 'selesai')",
		"INSERT INTO pembayaran (id_peminjaman, tanggal_pembayaran, jumlah_pembayaran) VALUES (1, '2024-04-15', 500000)",
		"INSERT INTO pembayaran (id_peminjaman, tanggal_pembayaran, jumlah_pembayaran) VALUES (1, '2024-04-20', 500000)",
	}

	for _, query := range queries {
		if _, err := pool.Exec(ctx, query); err != nil {
			return err
		}
	}
	return nil
}
