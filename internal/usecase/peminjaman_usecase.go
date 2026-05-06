package usecase

import (
	"context"
	"errors"

	"go-uts/internal/domain"
	"go-uts/internal/repository"
)

type PeminjamanUsecase interface {
	Create(ctx context.Context, data domain.Peminjaman) (domain.Peminjaman, error)
	GetAll(ctx context.Context) ([]domain.Peminjaman, error)
	GetByID(ctx context.Context, id int64) (domain.Peminjaman, error)
	UpdateStatus(ctx context.Context, id int64, status string) (domain.Peminjaman, error)
	Delete(ctx context.Context, id int64) error
}

type peminjamanUsecase struct {
	repo repository.PeminjamanRepository
}

func NewPeminjamanUsecase(repo repository.PeminjamanRepository) PeminjamanUsecase {
	return &peminjamanUsecase{repo: repo}
}

func (u *peminjamanUsecase) Create(ctx context.Context, data domain.Peminjaman) (domain.Peminjaman, error) {
	if data.IDAnggota == 0 {
		return data, errors.New("id_anggota wajib")
	}
	if data.TanggalPinjam == "" {
		return data, errors.New("tanggal_pinjam wajib")
	}
	if data.JumlahPinjam <= 0 {
		return data, errors.New("jumlah_pinjam harus lebih dari 0")
	}
	if data.Status == "" {
		data.Status = "pending"
	}

	return u.repo.Create(ctx, data)
}

func (u *peminjamanUsecase) GetAll(ctx context.Context) ([]domain.Peminjaman, error) {
	return u.repo.GetAll(ctx)
}

func (u *peminjamanUsecase) GetByID(ctx context.Context, id int64) (domain.Peminjaman, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *peminjamanUsecase) UpdateStatus(ctx context.Context, id int64, status string) (domain.Peminjaman, error) {
	if status == "" {
		return domain.Peminjaman{}, errors.New("status wajib")
	}
	return u.repo.UpdateStatus(ctx, id, status)
}

func (u *peminjamanUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}
