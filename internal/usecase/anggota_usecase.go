package usecase

import (
	"context"
	"errors"

	"go-uts/internal/domain"
	"go-uts/internal/repository"
)

type AnggotaUsecase interface {
	Create(ctx context.Context, data domain.Anggota) (domain.Anggota, error)
	GetAll(ctx context.Context) ([]domain.Anggota, error)
	GetByID(ctx context.Context, id int64) (domain.Anggota, error)
	Update(ctx context.Context, id int64, data domain.Anggota) (domain.Anggota, error)
	Delete(ctx context.Context, id int64) error
}

type anggotaUsecase struct {
	repo repository.AnggotaRepository
}

func NewAnggotaUsecase(repo repository.AnggotaRepository) AnggotaUsecase {
	return &anggotaUsecase{repo: repo}
}

func (u *anggotaUsecase) Create(ctx context.Context, data domain.Anggota) (domain.Anggota, error) {
	if data.Nama == "" {
		return data, errors.New("nama wajib diisi")
	}
	if data.Alamat == "" {
		return data, errors.New("alamat wajib diisi")
	}
	if data.Telepon == "" {
		return data, errors.New("telepon wajib diisi")
	}
	return u.repo.Create(ctx, data)
}

func (u *anggotaUsecase) GetAll(ctx context.Context) ([]domain.Anggota, error) {
	return u.repo.GetAll(ctx)
}

func (u *anggotaUsecase) GetByID(ctx context.Context, id int64) (domain.Anggota, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *anggotaUsecase) Update(ctx context.Context, id int64, data domain.Anggota) (domain.Anggota, error) {
	if data.Nama == "" {
		return data, errors.New("nama wajib diisi")
	}
	if data.Alamat == "" {
		return data, errors.New("alamat wajib diisi")
	}
	if data.Telepon == "" {
		return data, errors.New("telepon wajib diisi")
	}
	return u.repo.Update(ctx, id, data)
}

func (u *anggotaUsecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}
