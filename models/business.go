package models

import (
	"github.com/bowwsudrajat/promo-library-go/library/types"
)

type BusinessBulk struct {
	ID         string `json:"ID" db:"id"`
	Name       string `json:"Name" db:"name"`
	Code       string `json:"Code" db:"code"`
	LogoImgURL string `json:"LogoImgURL" db:"logo_img_url"`
	CompanyID  string `json:"CompanyID" db:"company_id"`
	StatusID   string `json:"StatusID" db:"status_id"`

	StatusName string `json:"StatusName" db:"status_name"`
}

type Business struct {
	ID         string `json:"ID" db:"id"`
	Name       string `json:"Name" db:"name"`
	Code       string `json:"Code" db:"code"`
	LogoImgURL string `json:"LogoImgURL" db:"logo_img_url"`
	CompanyID  string `json:"CompanyID" db:"company_id"`
	StatusID   string `json:"StatusID" db:"status_id"`

	Status Status `json:"Status"`
}

type FindAllBusinessParams struct {
	FindAllParams types.FindAllParams
	CompanyID     string
}
