package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"go-uts/internal/domain"
)

var ErrNotFound = errors.New("peminjaman not found")

const dateLayout = "2006-01-02"

type PeminjamanRepository interface {
	Create(ctx context.Context, data domain.Peminjaman) (domain.Peminjaman, error)
	GetAll(ctx context.Context) ([]domain.Peminjaman, error)
	GetByID(ctx context.Context, id int64) (domain.Peminjaman, error)
	UpdateStatus(ctx context.Context, id int64, status string) (domain.Peminjaman, error)
	Delete(ctx context.Context, id int64) error
}

type PeminjamanPGX struct {
	pool *pgxpool.Pool
}

func NewPeminjamanPGX(pool *pgxpool.Pool) *PeminjamanPGX {
	return &PeminjamanPGX{pool: pool}
}

func (r *PeminjamanPGX) Create(ctx context.Context, data domain.Peminjaman) (domain.Peminjaman, error) {
	var created domain.Peminjaman
	created = data

	parsedDate, err := time.Parse(dateLayout, data.TanggalPinjam)
	if err != nil {
		return created, err
	}

	query := `
		INSERT INTO peminjaman (id_anggota, tanggal_pinjam, jumlah_pinjam, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, id_anggota, tanggal_pinjam, jumlah_pinjam, status
	`

	row := r.pool.QueryRow(ctx, query, data.IDAnggota, parsedDate, data.JumlahPinjam, data.Status)
	var tanggal time.Time
	if err := row.Scan(&created.ID, &created.IDAnggota, &tanggal, &created.JumlahPinjam, &created.Status); err != nil {
		return created, err
	}
	created.TanggalPinjam = tanggal.Format(dateLayout)

	return created, nil
}

func (r *PeminjamanPGX) GetAll(ctx context.Context) ([]domain.Peminjaman, error) {
	query := `
		SELECT id, id_anggota, tanggal_pinjam, jumlah_pinjam, status
		FROM peminjaman
		ORDER BY id
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Peminjaman
	for rows.Next() {
		var item domain.Peminjaman
		var tanggal time.Time
		if err := rows.Scan(&item.ID, &item.IDAnggota, &tanggal, &item.JumlahPinjam, &item.Status); err != nil {
			return nil, err
		}
		item.TanggalPinjam = tanggal.Format(dateLayout)
		result = append(result, item)
	}

	return result, rows.Err()
}

func (r *PeminjamanPGX) GetByID(ctx context.Context, id int64) (domain.Peminjaman, error) {
	query := `
		SELECT id, id_anggota, tanggal_pinjam, jumlah_pinjam, status
		FROM peminjaman
		WHERE id = $1
	`

	var result domain.Peminjaman
	var tanggal time.Time
	row := r.pool.QueryRow(ctx, query, id)
	if err := row.Scan(&result.ID, &result.IDAnggota, &tanggal, &result.JumlahPinjam, &result.Status); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return result, ErrNotFound
		}
		return result, err
	}

	result.TanggalPinjam = tanggal.Format(dateLayout)
	return result, nil
}

func (r *PeminjamanPGX) UpdateStatus(ctx context.Context, id int64, status string) (domain.Peminjaman, error) {
	query := `
		UPDATE peminjaman
		SET status = $1
		WHERE id = $2
		RETURNING id, id_anggota, tanggal_pinjam, jumlah_pinjam, status
	`

	var result domain.Peminjaman
	var tanggal time.Time
	row := r.pool.QueryRow(ctx, query, status, id)
	if err := row.Scan(&result.ID, &result.IDAnggota, &tanggal, &result.JumlahPinjam, &result.Status); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return result, ErrNotFound
		}
		return result, err
	}

	result.TanggalPinjam = tanggal.Format(dateLayout)
	return result, nil
}

func (r *PeminjamanPGX) Delete(ctx context.Context, id int64) error {
	commandTag, err := r.pool.Exec(ctx, `DELETE FROM peminjaman WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
