package businessconfig

import (
	"github.com/bowwsudrajat/promo-library-go/library/types"
	"github.com/bowwsudrajat/promo-library-go/models"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	FindAll(*gin.Context, models.FindAllBusinessConfigParams) ([]*models.BusinessConfig, *types.Error)
	Find(*gin.Context, int) (*models.BusinessConfig, *types.Error)
	Create(*gin.Context, *models.BusinessConfig) (*models.BusinessConfig, *types.Error)
	Update(*gin.Context, *models.BusinessConfig) (*models.BusinessConfig, *types.Error)
}
