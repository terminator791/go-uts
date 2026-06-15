package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"go-uts/internal/domain"
)

type PembayaranRepository interface {
	Create(ctx context.Context, data domain.Pembayaran) (domain.Pembayaran, error)
	GetAll(ctx context.Context) ([]domain.Pembayaran, error)
	GetByID(ctx context.Context, id int64) (domain.Pembayaran, error)
	GetByPeminjamanID(ctx context.Context, peminjamanID int64) ([]domain.Pembayaran, error)
	Update(ctx context.Context, id int64, data domain.Pembayaran) (domain.Pembayaran, error)
	Delete(ctx context.Context, id int64) (domain.Pembayaran, error)
	SumByPeminjamanID(ctx context.Context, peminjamanID int64) (int64, error)
}

const pembayaranDateLayout = "2006-01-02"

type PembayaranPGX struct {
	pool *pgxpool.Pool
}

func NewPembayaranPGX(pool *pgxpool.Pool) *PembayaranPGX {
	return &PembayaranPGX{pool: pool}
}

func (r *PembayaranPGX) Create(ctx context.Context, data domain.Pembayaran) (domain.Pembayaran, error) {
	var created domain.Pembayaran
	created = data

	parsedDate, err := time.Parse(pembayaranDateLayout, data.TanggalPembayaran)
	if err != nil {
		return created, err
	}

	query := `
		INSERT INTO pembayaran (id_peminjaman, tanggal_pembayaran, jumlah_pembayaran)
		VALUES ($1, $2, $3)
		RETURNING id, id_peminjaman, tanggal_pembayaran, jumlah_pembayaran
	`

	row := r.pool.QueryRow(ctx, query, data.IDPeminjaman, parsedDate, data.JumlahPembayaran)
	var tanggal time.Time
	if err := row.Scan(&created.ID, &created.IDPeminjaman, &tanggal, &created.JumlahPembayaran); err != nil {
		return created, err
	}
	created.TanggalPembayaran = tanggal.Format(pembayaranDateLayout)

	return created, nil
}

func (r *PembayaranPGX) GetAll(ctx context.Context) ([]domain.Pembayaran, error) {
	query := `
		SELECT id, id_peminjaman, tanggal_pembayaran, jumlah_pembayaran
		FROM pembayaran
		ORDER BY id
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Pembayaran
	for rows.Next() {
		var item domain.Pembayaran
		var tanggal time.Time
		if err := rows.Scan(&item.ID, &item.IDPeminjaman, &tanggal, &item.JumlahPembayaran); err != nil {
			return nil, err
		}
		item.TanggalPembayaran = tanggal.Format(pembayaranDateLayout)
		result = append(result, item)
	}

	return result, rows.Err()
}

func (r *PembayaranPGX) GetByID(ctx context.Context, id int64) (domain.Pembayaran, error) {
	query := `
		SELECT id, id_peminjaman, tanggal_pembayaran, jumlah_pembayaran
		FROM pembayaran
		WHERE id = $1
	`

	var result domain.Pembayaran
	var tanggal time.Time
	row := r.pool.QueryRow(ctx, query, id)
	if err := row.Scan(&result.ID, &result.IDPeminjaman, &tanggal, &result.JumlahPembayaran); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return result, ErrNotFound
		}
		return result, err
	}

	result.TanggalPembayaran = tanggal.Format(pembayaranDateLayout)
	return result, nil
}

func (r *PembayaranPGX) GetByPeminjamanID(ctx context.Context, peminjamanID int64) ([]domain.Pembayaran, error) {
	query := `
		SELECT id, id_peminjaman, tanggal_pembayaran, jumlah_pembayaran
		FROM pembayaran
		WHERE id_peminjaman = $1
		ORDER BY id
	`

	rows, err := r.pool.Query(ctx, query, peminjamanID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Pembayaran
	for rows.Next() {
		var item domain.Pembayaran
		var tanggal time.Time
		if err := rows.Scan(&item.ID, &item.IDPeminjaman, &tanggal, &item.JumlahPembayaran); err != nil {
			return nil, err
		}
		item.TanggalPembayaran = tanggal.Format(pembayaranDateLayout)
		result = append(result, item)
	}

	return result, rows.Err()
}

func (r *PembayaranPGX) Update(ctx context.Context, id int64, data domain.Pembayaran) (domain.Pembayaran, error) {
	parsedDate, err := time.Parse(pembayaranDateLayout, data.TanggalPembayaran)
	if err != nil {
		return domain.Pembayaran{}, err
	}

	query := `
		UPDATE pembayaran
		SET id_peminjaman = $1, tanggal_pembayaran = $2, jumlah_pembayaran = $3
		WHERE id = $4
		RETURNING id, id_peminjaman, tanggal_pembayaran, jumlah_pembayaran
	`

	var updated domain.Pembayaran
	var tanggal time.Time
	row := r.pool.QueryRow(ctx, query, data.IDPeminjaman, parsedDate, data.JumlahPembayaran, id)
	if err := row.Scan(&updated.ID, &updated.IDPeminjaman, &tanggal, &updated.JumlahPembayaran); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return updated, ErrNotFound
		}
		return updated, err
	}
	updated.TanggalPembayaran = tanggal.Format(pembayaranDateLayout)

	return updated, nil
}

func (r *PembayaranPGX) Delete(ctx context.Context, id int64) (domain.Pembayaran, error) {
	query := `
		DELETE FROM pembayaran
		WHERE id = $1
		RETURNING id, id_peminjaman, tanggal_pembayaran, jumlah_pembayaran
	`

	var deleted domain.Pembayaran
	var tanggal time.Time
	row := r.pool.QueryRow(ctx, query, id)
	if err := row.Scan(&deleted.ID, &deleted.IDPeminjaman, &tanggal, &deleted.JumlahPembayaran); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return deleted, ErrNotFound
		}
		return deleted, err
	}
	deleted.TanggalPembayaran = tanggal.Format(pembayaranDateLayout)

	return deleted, nil
}

func (r *PembayaranPGX) SumByPeminjamanID(ctx context.Context, peminjamanID int64) (int64, error) {
	query := `
		SELECT COALESCE(SUM(jumlah_pembayaran), 0)
		FROM pembayaran
		WHERE id_peminjaman = $1
	`

	var total int64
	if err := r.pool.QueryRow(ctx, query, peminjamanID).Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}
