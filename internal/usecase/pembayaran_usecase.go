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
	Update(ctx context.Context, id int64, data domain.Pembayaran) (domain.Pembayaran, error)
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
	if err := u.validateRemaining(ctx, data.IDPeminjaman, data.JumlahPembayaran, 0); err != nil {
		return data, err
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

func (u *pembayaranUsecase) Update(ctx context.Context, id int64, data domain.Pembayaran) (domain.Pembayaran, error) {
	if data.IDPeminjaman == 0 {
		return data, errors.New("id_peminjaman wajib")
	}
	if data.TanggalPembayaran == "" {
		return data, errors.New("tanggal_pembayaran wajib")
	}
	if data.JumlahPembayaran <= 0 {
		return data, errors.New("jumlah_pembayaran harus lebih dari 0")
	}

	current, err := u.pembayaranRepo.GetByID(ctx, id)
	if err != nil {
		return data, err
	}
	if err := u.validateRemaining(ctx, data.IDPeminjaman, data.JumlahPembayaran, current.JumlahPembayaran); err != nil {
		return data, err
	}

	updated, err := u.pembayaranRepo.Update(ctx, id, data)
	if err != nil {
		return updated, err
	}

	if err := u.updatePeminjamanStatus(ctx, data.IDPeminjaman); err != nil {
		return updated, err
	}
	if current.IDPeminjaman != data.IDPeminjaman {
		if err := u.updatePeminjamanStatus(ctx, current.IDPeminjaman); err != nil {
			return updated, err
		}
	}

	return updated, nil
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

func (u *pembayaranUsecase) validateRemaining(ctx context.Context, peminjamanID int64, amount int64, excludeAmount int64) error {
	peminjaman, err := u.peminjamanRepo.GetByID(ctx, peminjamanID)
	if err != nil {
		return err
	}
	if peminjaman.Status == "selesai" {
		return errors.New("peminjaman sudah selesai")
	}

	currentTotal, err := u.pembayaranRepo.SumByPeminjamanID(ctx, peminjamanID)
	if err != nil {
		return err
	}

	adjustedTotal := currentTotal - excludeAmount + amount
	if adjustedTotal > peminjaman.JumlahPinjam {
		return errors.New("jumlah_pembayaran melebihi sisa pinjaman")
	}

	return nil
}
