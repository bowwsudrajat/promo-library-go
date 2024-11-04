package business

import (
	"github.com/bowwsudrajat/promo-library-go/library/types"
	"github.com/bowwsudrajat/promo-library-go/models"
	"github.com/gin-gonic/gin"
)

// Repository is the contract between Repository and usecase
type Repository interface {
	FindAll(*gin.Context, models.FindAllBusinessParams) ([]*models.Business, *types.Error)
	Find(*gin.Context, string) (*models.Business, *types.Error)
	Create(*gin.Context, *models.Business) (*models.Business, *types.Error)
	Update(*gin.Context, *models.Business) (*models.Business, *types.Error)

	UpdateStatus(*gin.Context, string, string) (*models.Business, *types.Error)
}
