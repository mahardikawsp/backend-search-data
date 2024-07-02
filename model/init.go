package model

type (
	FilterOptions struct {
		Page  int `query:"page"`
		Limit int `query:"limit"`
	}
)
