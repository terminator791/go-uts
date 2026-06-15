package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"go-uts/internal/domain"
)

type AnggotaRepository interface {
	Create(ctx context.Context, data domain.Anggota) (domain.Anggota, error)
	GetAll(ctx context.Context) ([]domain.Anggota, error)
	GetByID(ctx context.Context, id int64) (domain.Anggota, error)
	Update(ctx context.Context, id int64, data domain.Anggota) (domain.Anggota, error)
	Delete(ctx context.Context, id int64) error
}

type AnggotaPGX struct {
	pool *pgxpool.Pool
}

func NewAnggotaPGX(pool *pgxpool.Pool) *AnggotaPGX {
	return &AnggotaPGX{pool: pool}
}

func (r *AnggotaPGX) Create(ctx context.Context, data domain.Anggota) (domain.Anggota, error) {
	query := `
		INSERT INTO anggota (nama, alamat, telepon)
		VALUES ($1, $2, $3)
		RETURNING id, nama, alamat, telepon
	`
	var created domain.Anggota
	err := r.pool.QueryRow(ctx, query, data.Nama, data.Alamat, data.Telepon).
		Scan(&created.ID, &created.Nama, &created.Alamat, &created.Telepon)
	return created, err
}

func (r *AnggotaPGX) GetAll(ctx context.Context) ([]domain.Anggota, error) {
	query := `SELECT id, nama, alamat, telepon FROM anggota ORDER BY id`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Anggota
	for rows.Next() {
		var item domain.Anggota
		if err := rows.Scan(&item.ID, &item.Nama, &item.Alamat, &item.Telepon); err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, rows.Err()
}

func (r *AnggotaPGX) GetByID(ctx context.Context, id int64) (domain.Anggota, error) {
	query := `SELECT id, nama, alamat, telepon FROM anggota WHERE id = $1`
	var result domain.Anggota
	err := r.pool.QueryRow(ctx, query, id).Scan(&result.ID, &result.Nama, &result.Alamat, &result.Telepon)
	if errors.Is(err, pgx.ErrNoRows) {
		return result, ErrNotFound
	}
	return result, err
}

func (r *AnggotaPGX) Update(ctx context.Context, id int64, data domain.Anggota) (domain.Anggota, error) {
	query := `
		UPDATE anggota
		SET nama = $1, alamat = $2, telepon = $3
		WHERE id = $4
		RETURNING id, nama, alamat, telepon
	`
	var updated domain.Anggota
	err := r.pool.QueryRow(ctx, query, data.Nama, data.Alamat, data.Telepon, id).
		Scan(&updated.ID, &updated.Nama, &updated.Alamat, &updated.Telepon)
	if errors.Is(err, pgx.ErrNoRows) {
		return updated, ErrNotFound
	}
	return updated, err
}

func (r *AnggotaPGX) Delete(ctx context.Context, id int64) error {
	commandTag, err := r.pool.Exec(ctx, `DELETE FROM anggota WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
