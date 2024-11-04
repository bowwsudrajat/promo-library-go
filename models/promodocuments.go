package models

import (
	"github.com/bowwsudrajat/promo-library-go/library/types"
)

type PromoDocumentBulk struct {
	ID          string `json:"ID" db:"id"`
	PromoID     string `json:"PromoID" db:"promo_id"`
	DocumentURL string `json:"DocumentURL" db:"document_url"`
	StatusID    string `json:"StatusID" db:"status_id"`
}

type PromoDocument struct {
	ID          string `json:"ID" db:"id"`
	PromoID     string `json:"PromoID" db:"promo_id" validate:"required"`
	DocumentURL string `json:"DocumentURL" db:"document_url" validate:"required"`
	StatusID    string `json:"StatusID" db:"status_id"`
}

type FindAllPromoDocumentParams struct {
	FindAllParams types.FindAllParams
	PromoID       string
}
