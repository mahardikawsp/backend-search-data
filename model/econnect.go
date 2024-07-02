package model

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
	"time"
)

type (
	Econnect struct {
		ID                  int `gorm:"primaryKey"`
		Indikator           string
		NamaData            string
		DataSetDibuat       time.Time
		DataSetDiperbaharui time.Time
		Satuan              string
		Urusan              string
		Kategori            string
		Produsen            string
		KontakProdusen      string
		EmailProdusen       string
		Frekuensi           string
		Url                 string
		IsPublik            int32
	}

	EconnectParam struct {
		ID             int    `form:"id" param:"id" query:"id"`
		NamaData       string `form:"nama_data"`
		Satuan         string `form:"satuan"`
		Urusan         string `form:"urusan"`
		Kategori       string `form:"kategori"`
		Produsen       string `form:"produsen"`
		KontakProdusen string `form:"kontak_produsen"`
		EmailProdusen  string `form:"email_produsen"`
		Frekuensi      string `form:"frekuensi"`
		Url            string `form:"url"`
		IsPublik       int32  `form:"is_publik"`
	}

	EconnectResponse struct {
		ID                  int       `json:"id"`
		Indikator           string    `json:"indikator"`
		NamaData            string    `json:"nama_data"`
		DataSetDibuat       time.Time `json:"data_set_dibuat"`
		DataSetDiperbaharui time.Time `json:"data_set_diperbaharui"`
		Satuan              string    `json:"satuan"`
		Urusan              string    `json:"urusan"`
		Kategori            string    `json:"kategori"`
		Produsen            string    `json:"produsen"`
		KontakProdusen      string    `json:"kontak_produsen"`
		EmailProdusen       string    `json:"email_produsen"`
		Frekuensi           string    `json:"frekuensi"`
		Url                 string    `json:"url"`
		IsPublik            int32     `json:"is_publik"`
	}
)

func CreateEconnectTable(db *gorm.DB) error {
	check := db.Migrator().HasTable(&Econnect{})
	if check {
		log.Error("Table econnect already exist")
		return nil
	}

	err := db.Migrator().CreateTable(&Econnect{})
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (e *Econnect) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = int(time.Now().Unix())

	e.DataSetDibuat = time.Now()
	e.DataSetDiperbaharui = time.Now()

	return nil
}

func (prm *EconnectParam) Validate(ctx context.Context) (*Econnect, error) {
	if prm.NamaData == "" {
		return nil, errors.New("nama_data tidak boleh kosong")
	} else {
		prm.NamaData = cases.Title(language.Indonesian).String(prm.NamaData)
	}

	if prm.Satuan == "" {
		return nil, errors.New("satuan tidak boleh kosong")
	} else {
		prm.Satuan = cases.Title(language.Indonesian).String(prm.Satuan)
	}

	if prm.Urusan == "" {
		return nil, errors.New("urusan tidak boleh kosong")
	} else {
		prm.Urusan = cases.Title(language.Indonesian).String(prm.Urusan)
	}

	if prm.Kategori == "" {
		return nil, errors.New("kategori tidak boleh kosong")
	} else {
		prm.Kategori = cases.Title(language.Indonesian).String(prm.Kategori)
	}

	if prm.Produsen == "" {
		return nil, errors.New("produsen tidak boleh kosong")
	} else {
		prm.Produsen = cases.Title(language.Indonesian).String(prm.Produsen)
	}

	if prm.Frekuensi == "" {
		return nil, errors.New("frekuensi tidak boleh kosong")
	}

	if prm.Url == "" {
		return nil, errors.New("url tidak boleh kosong")
	}

	return &Econnect{
		Indikator:      prm.NamaData,
		NamaData:       prm.NamaData,
		Satuan:         prm.Satuan,
		Urusan:         prm.Urusan,
		Kategori:       prm.Kategori,
		Produsen:       prm.Produsen,
		KontakProdusen: prm.KontakProdusen,
		EmailProdusen:  prm.EmailProdusen,
		Frekuensi:      prm.Frekuensi,
		Url:            prm.Url,
		IsPublik:       prm.IsPublik,
	}, nil
}

func (data *Econnect) Response() *EconnectResponse {
	if data == nil {
		return nil
	}

	return &EconnectResponse{
		ID:                  data.ID,
		Indikator:           data.Indikator,
		NamaData:            data.NamaData,
		DataSetDibuat:       data.DataSetDibuat,
		DataSetDiperbaharui: data.DataSetDiperbaharui,
		Satuan:              data.Satuan,
		Urusan:              data.Urusan,
		Kategori:            data.Kategori,
		Produsen:            data.Produsen,
		KontakProdusen:      data.KontakProdusen,
		EmailProdusen:       data.EmailProdusen,
		Frekuensi:           data.Frekuensi,
		Url:                 data.Url,
		IsPublik:            data.IsPublik,
	}
}

func (prm *EconnectParam) ValidateUpdate(ctx context.Context, data *Econnect) (*Econnect, error) {
	if prm.NamaData != "" {
		data.NamaData = cases.Title(language.Indonesian).String(prm.NamaData)
		data.Indikator = cases.Title(language.Indonesian).String(prm.NamaData)
	}

	if prm.Satuan != "" {
		data.Satuan = cases.Title(language.Indonesian).String(prm.Satuan)
	}

	if prm.Urusan != "" {
		data.Urusan = cases.Title(language.Indonesian).String(prm.Urusan)
	}

	if prm.Kategori != "" {
		data.Kategori = cases.Title(language.Indonesian).String(prm.Kategori)
	}

	if prm.Produsen != "" {
		data.Produsen = cases.Title(language.Indonesian).String(prm.Produsen)
	}

	if prm.KontakProdusen != "" {
		data.KontakProdusen = cases.Title(language.Indonesian).String(prm.KontakProdusen)
	}

	if prm.EmailProdusen != "" {
		data.EmailProdusen = cases.Title(language.Indonesian).String(prm.EmailProdusen)
	}

	if prm.Frekuensi != "" {
		data.Frekuensi = cases.Title(language.Indonesian).String(prm.Frekuensi)
	}

	if prm.Url != "" {
		data.Url = cases.Title(language.Indonesian).String(prm.Url)
	}

	data.DataSetDiperbaharui = time.Now()

	return data, nil
}
