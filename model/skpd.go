package model

type (
	Skpd struct {
		ID       string `json:"id"`
		KodeSkpd string `gorm:"column:KODE_SKPD" json:"kode_skpd" mapstructure:"KODE_SKPD"`
		NamaSkpd string `gorm:"column:NAMA_SKPD" json:"nama_skpd" mapstructure:"NAMA_SKPD"`
	}

	Urusan struct {
		ID         string `json:"id"`
		KodeUrusan string `gorm:"column:KODE_URUSAN" json:"kode_urusan" mapstructure:"KODE_URUSAN"`
		NamaUrusan string `gorm:"column:URUSAN" json:"nama_urusan" mapstructure:"URUSAN"`
	}
)

func (Skpd) TableName() string {
	return "m_skpd"
}

func (Urusan) TableName() string {
	return "m_urusan"
}
