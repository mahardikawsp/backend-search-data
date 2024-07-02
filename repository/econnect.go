package repository

import (
	"example/db/mysql"
	"example/helper"
	"example/model"
	"github.com/labstack/gommon/log"
)

type (
	EconnectRepository interface {
		Create(data *model.Econnect) error
		List(prm *model.FilterOptions) ([]model.Econnect, int64, error)
		Detail(prm *model.EconnectParam) (*model.Econnect, error)
		Update(prm *model.Econnect) error
	}

	repoEconnect struct {
		db *mysql.Client
	}
)

func NewEconnectRepository() EconnectRepository {
	return &repoEconnect{
		mysql.GetDB(),
	}
}

func (r repoEconnect) Create(data *model.Econnect) error {
	result := r.db.Create(data)
	if result.Error != nil {
		log.Error(result.Error.Error())
		return result.Error
	}

	return nil
}

func (r repoEconnect) List(prm *model.FilterOptions) ([]model.Econnect, int64, error) {
	page, limit := helper.GetLimitOffset(prm.Page, prm.Limit)
	var (
		total    int64
		econnect []model.Econnect
	)

	results := r.db.Limit(limit).Offset(page).Find(&econnect).Count(&total)
	if results.Error != nil {
		log.Error(results.Error.Error())
		return nil, 0, results.Error
	}

	return econnect, total, nil
}

func (r repoEconnect) Detail(prm *model.EconnectParam) (*model.Econnect, error) {
	var econnect *model.Econnect

	result := r.db.First(&econnect, "id = ?", prm.ID)
	if result.Error != nil {
		log.Error(result.Error.Error())
		return nil, result.Error
	}

	return econnect, nil
}

func (r repoEconnect) Update(prm *model.Econnect) error {
	result := r.db.Model(&model.Econnect{}).Where("id = ?", prm.ID).Updates(prm)
	if result.Error != nil {
		log.Error(result.Error.Error())
		return result.Error
	}

	return nil
}
