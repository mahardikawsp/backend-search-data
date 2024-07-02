package service

import (
	"context"
	"example/helper"
	"example/model"
	"example/repository"
	"net/http"
)

type (
	EconnectService interface {
		Create(param *model.EconnectParam, ctx context.Context) (error, int)
		List(param *model.FilterOptions) (*helper.PaginateResponse, error)
		Detail(param *model.EconnectParam) (*model.EconnectResponse, error)
		Update(param *model.EconnectParam, ctx context.Context) (error, int)
	}

	EconnectServiceImpl struct {
		repo repository.EconnectRepository
	}
)

func NewEconnectService() EconnectService {
	return &EconnectServiceImpl{
		repository.NewEconnectRepository(),
	}
}

func (e EconnectServiceImpl) Create(param *model.EconnectParam, ctx context.Context) (error, int) {
	prm, err := param.Validate(ctx)
	if err != nil {
		return err, http.StatusBadRequest
	}

	err = e.repo.Create(prm)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (e EconnectServiceImpl) List(param *model.FilterOptions) (*helper.PaginateResponse, error) {
	list, total, err := e.repo.List(param)
	if err != nil {
		return nil, err
	}

	var response []*model.EconnectResponse
	for _, item := range list {
		response = append(response, item.Response())
	}

	return helper.Paginate(&helper.PaginateParam{
		Page:  param.Page,
		Limit: param.Limit,
		Count: int(total),
	}, response), nil
}

func (e EconnectServiceImpl) Detail(param *model.EconnectParam) (*model.EconnectResponse, error) {
	data, err := e.repo.Detail(param)
	if err != nil {
		return nil, err
	}

	return data.Response(), nil
}

func (e EconnectServiceImpl) Update(param *model.EconnectParam, ctx context.Context) (error, int) {
	data, err := e.repo.Detail(param)
	if err != nil {
		return err, http.StatusBadRequest
	}

	econnect, err := param.ValidateUpdate(ctx, data)
	if err != nil {
		return err, http.StatusBadRequest
	}

	err = e.repo.Update(econnect)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
