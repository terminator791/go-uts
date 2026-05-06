package usecase

import (
	"context"
	"errors"

	"go-uts/internal/domain"
	"go-uts/internal/repository"
)

type PembayaranUsecase interface {
	Create(ctx context.Context, data domain.Pembayaran) (domain.Pembayaran, error)
	GetAll(ctx context.Context) ([]domain.Pembayaran, error)
	GetByID(ctx context.Context, id int64) (domain.Pembayaran, error)
	GetByPeminjamanID(ctx context.Context, peminjamanID int64) ([]domain.Pembayaran, error)
	Delete(ctx context.Context, id int64) (domain.Pembayaran, error)
}

type pembayaranUsecase struct {
	pembayaranRepo repository.PembayaranRepository
	peminjamanRepo repository.PeminjamanRepository
}

func NewPembayaranUsecase(
	pembayaranRepo repository.PembayaranRepository,
	peminjamanRepo repository.PeminjamanRepository,
) PembayaranUsecase {
	return &pembayaranUsecase{
		pembayaranRepo: pembayaranRepo,
		peminjamanRepo: peminjamanRepo,
	}
}

func (u *pembayaranUsecase) Create(ctx context.Context, data domain.Pembayaran) (domain.Pembayaran, error) {
	if data.IDPeminjaman == 0 {
		return data, errors.New("id_peminjaman wajib")
	}
	if data.TanggalPembayaran == "" {
		return data, errors.New("tanggal_pembayaran wajib")
	}
	if data.JumlahPembayaran <= 0 {
		return data, errors.New("jumlah_pembayaran harus lebih dari 0")
	}

	created, err := u.pembayaranRepo.Create(ctx, data)
	if err != nil {
		return created, err
	}

	if err := u.updatePeminjamanStatus(ctx, data.IDPeminjaman); err != nil {
		return created, err
	}

	return created, nil
}

func (u *pembayaranUsecase) GetAll(ctx context.Context) ([]domain.Pembayaran, error) {
	return u.pembayaranRepo.GetAll(ctx)
}

func (u *pembayaranUsecase) GetByID(ctx context.Context, id int64) (domain.Pembayaran, error) {
	return u.pembayaranRepo.GetByID(ctx, id)
}

func (u *pembayaranUsecase) GetByPeminjamanID(ctx context.Context, peminjamanID int64) ([]domain.Pembayaran, error) {
	return u.pembayaranRepo.GetByPeminjamanID(ctx, peminjamanID)
}

func (u *pembayaranUsecase) Delete(ctx context.Context, id int64) (domain.Pembayaran, error) {
	deleted, err := u.pembayaranRepo.Delete(ctx, id)
	if err != nil {
		return deleted, err
	}

	if err := u.updatePeminjamanStatus(ctx, deleted.IDPeminjaman); err != nil {
		return deleted, err
	}

	return deleted, nil
}

func (u *pembayaranUsecase) updatePeminjamanStatus(ctx context.Context, peminjamanID int64) error {
	peminjaman, err := u.peminjamanRepo.GetByID(ctx, peminjamanID)
	if err != nil {
		return err
	}

	total, err := u.pembayaranRepo.SumByPeminjamanID(ctx, peminjamanID)
	if err != nil {
		return err
	}

	status := "pending"
	if total >= peminjaman.JumlahPinjam {
		status = "selesai"
	}

	_, err = u.peminjamanRepo.UpdateStatus(ctx, peminjamanID, status)
	return err
}
