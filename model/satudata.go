package model

import (
	"context"
	"strings"
	"time"
)

type (
	SearchParam struct {
		Keyword  string `query:"keyword"`
		Urusan   string `query:"urusan"`
		Skpd     string `query:"skpd"`
		Kategori string `query:"kategori"`
		Page     int    `query:"page"`
		Limit    int    `query:"limit"`
		Offset   int    `query:"offset"`
		Sort     string `query:"sort"`
		Order    string `query:"order"`
	}

	NewSearchParam struct {
		Keyword  string
		Urusan   []string
		Skpd     []string
		Kategori []string
		Page     int
		Limit    int
		Offset   int
		Sort     string
		Order    string
	}

	Satudata struct {
		ID                  string  `json:"id"`
		Indikator           string  `json:"indikator" mapstructure:"INDIKATOR"`
		NamaData            string  `json:"nama_data" mapstructure:"NAMA_DATA"`
		Tags                string  `json:"tags"`
		Source              string  `json:"source"`
		DataSetDibuat       int64   `json:"data_set_dibuat" mapstructure:"DATA_SET_DIBUAT"`
		DataSetDiperbaharui int64   `json:"data_set_diperbaharui" mapstructure:"DATA_SET_DIPERBAHARUI"`
		Satuan              string  `json:"satuan"`
		Urusan              string  `json:"urusan"`
		Kategori            string  `json:"kategori"`
		Produsen            string  `json:"produsen"`
		KontakProdusen      string  `json:"kontak_produsen" mapstructure:"KONTAK_PRODUSEN"`
		EmailProdusen       string  `json:"email_produsen" mapstructure:"EMAIL_PRODUSEN"`
		Frekuensi           string  `json:"frekuensi"`
		Url                 string  `json:"url"`
		IsPublik            float64 `json:"is_publik" mapstructure:"IS_PUBLIK"`
	}

	SatudataResponse struct {
		ID                  string    `json:"id"`
		Indikator           string    `json:"indikator"`
		NamaData            string    `json:"nama_data"`
		Tags                string    `json:"tags"`
		Source              string    `json:"source"`
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
		IsPublik            float64   `json:"is_publik"`
	}
)

func (prm *SearchParam) ValidateParam(ctx context.Context) *NewSearchParam {
	var (
		urusan   []string
		skpd     []string
		kategori []string
	)
	if prm.Urusan != "" {
		urusan = strings.Split(prm.Urusan, ",")
	} else {
		urusan = nil
	}

	if prm.Skpd != "" {
		skpd = strings.Split(prm.Skpd, ",")
	} else {
		skpd = nil
	}

	if prm.Kategori != "" {
		kategori = strings.Split(prm.Kategori, ",")
	} else {
		kategori = nil
	}

	return &NewSearchParam{
		Keyword:  prm.Keyword,
		Urusan:   urusan,
		Skpd:     skpd,
		Kategori: kategori,
		Page:     prm.Page,
		Limit:    prm.Limit,
		Offset:   prm.Offset,
		Sort:     prm.Sort,
		Order:    prm.Order,
	}
}

func (data *Satudata) Response() *SatudataResponse {
	if data == nil {
		return nil
	}

	return &SatudataResponse{
		ID:                  data.ID,
		Indikator:           data.Indikator,
		NamaData:            data.NamaData,
		Tags:                data.Tags,
		Source:              data.Source,
		DataSetDibuat:       time.UnixMilli(data.DataSetDibuat),
		DataSetDiperbaharui: time.UnixMilli(data.DataSetDiperbaharui),
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
