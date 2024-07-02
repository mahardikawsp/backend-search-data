package service

import (
	"context"
	"example/helper"
	"example/model"
	"example/repository"
)

type (
	SearchService interface {
		List(ctx context.Context, param *model.SearchParam) (interface{}, error)
	}

	SearchServiceImpl struct {
		repo repository.SearchRepository
	}
)

func NewSearchService() SearchService {
	return &SearchServiceImpl{
		repository.NewSearchRepository(),
	}
}

func (s SearchServiceImpl) List(ctx context.Context, param *model.SearchParam) (interface{}, error) {
	searchParam := param.ValidateParam(ctx)

	list, total, err := s.repo.List(ctx, searchParam)
	if err != nil {
		return nil, err
	}

	var response []*model.SatudataResponse
	for _, item := range list {
		response = append(response, item.Response())
	}

	return helper.Paginate(&helper.PaginateParam{
		Page:  param.Page,
		Limit: param.Limit,
		Count: total,
	}, response), nil
}
