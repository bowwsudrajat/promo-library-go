package promo

import (
	"github.com/bowwsudrajat/promo-library-go/library/types"
	"github.com/bowwsudrajat/promo-library-go/models"
	"github.com/gin-gonic/gin"
)

// Repository is the contract between Repository and usecase
type Repository interface {
	FindAll(*gin.Context, models.FindAllPromoParams) ([]*models.Promo, *types.Error)
	Find(*gin.Context, string) (*models.Promo, *types.Error)
	Create(*gin.Context, *models.Promo) (*models.Promo, *types.Error)
	Update(*gin.Context, *models.Promo) (*models.Promo, *types.Error)

	FindStatus(*gin.Context) ([]*models.Status, *types.Error)
	UpdateStatus(*gin.Context, string, string) (*models.Promo, *types.Error)

	// APPROVAL
	ApprovePromo(*gin.Context, string) (*models.Promo, *types.Error)
	RejectPromo(*gin.Context, string, string) (*models.Promo, *types.Error)
}

type DocumentRepository interface {
	FindAll(*gin.Context, models.FindAllPromoDocumentParams) ([]*models.PromoDocument, *types.Error)
	Find(*gin.Context, string) (*models.PromoDocument, *types.Error)
	Create(*gin.Context, *models.PromoDocument) (*models.PromoDocument, *types.Error)
	Update(*gin.Context, *models.PromoDocument) (*models.PromoDocument, *types.Error)

	UpdateStatus(*gin.Context, string, string) (*models.PromoDocument, *types.Error)
}
