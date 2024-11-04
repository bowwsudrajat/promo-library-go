package user

import (
	"github.com/bowwsudrajat/promo-library-go/library/types"
	"github.com/bowwsudrajat/promo-library-go/models"
	"github.com/gin-gonic/gin"
)

type Repository interface {
	FindAll(*gin.Context, models.FindAllUserParams) ([]*models.User, *types.Error)
	Find(*gin.Context, string) (*models.User, *types.Error)
	Create(*gin.Context, *models.User) (*models.User, *types.Error)
	Update(*gin.Context, *models.User) (*models.User, *types.Error)
	UpdateStatus(*gin.Context, string, string) (*models.User, *types.Error)
}

type PermissionRepository interface {
	FindAll(*gin.Context, models.FindAllUserPermissionParams) ([]*models.UserPermission, *types.Error)
	Find(*gin.Context, string) (*models.UserPermission, *types.Error)
	Create(*gin.Context, *models.CreateUpdateUserPermission) (*models.UserPermission, *types.Error)
	DeleteByUserID(*gin.Context, string) *types.Error

	CreateBunch(*gin.Context, string, models.FindAllUserPermissionParams) *types.Error
}
