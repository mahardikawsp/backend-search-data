package service

import (
	"example/model"
	"example/repository"
)

type (
	MasterDataService interface {
		ListSkpd() ([]model.Skpd, error)
		ListUrusan() ([]model.Urusan, error)
	}

	MasterDataServiceImpl struct {
		repo repository.MasterDataRepository
	}
)

func NewMasterDataService() MasterDataService {
	return &MasterDataServiceImpl{
		repository.NewMasterDataRepository(),
	}
}

func (m MasterDataServiceImpl) ListSkpd() ([]model.Skpd, error) {
	list, err := m.repo.ListSkpd()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (m MasterDataServiceImpl) ListUrusan() ([]model.Urusan, error) {
	list, err := m.repo.ListUrusan()
	if err != nil {
		return nil, err
	}

	return list, nil
}
