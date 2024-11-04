package brand

import (
	"github.com/bowwsudrajat/promo-library-go/library/types"
	"github.com/bowwsudrajat/promo-library-go/models"
	"github.com/gin-gonic/gin"
)

// Usecase is the contract between Repository and usecase
type Usecase interface {
	FindAll(context *gin.Context, params models.FindAllBrandParams) ([]*models.Brand, *types.Error)
	Find(context *gin.Context, id string) (*models.Brand, *types.Error)
	Count(context *gin.Context, params models.FindAllBrandParams) (int, *types.Error)
	Create(context *gin.Context, newData models.Brand) (*models.Brand, *types.Error)
	Update(context *gin.Context, id string, updatedData models.Brand) (*models.Brand, *types.Error)

	UpdateStatus(*gin.Context, string, string) (*models.Brand, *types.Error)
}
