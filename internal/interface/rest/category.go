package rest

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/categories"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryHand struct {
	uc categories.ICategoryCrud
}

func NewCategoryHand(uc categories.ICategoryCrud) CategoryHand {
	return CategoryHand{
		uc,
	}
}

func (rh *CategoryHand) GetAllCategories(c *gin.Context) {

	rs, err := rh.uc.GetAllCategories(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
		Data:    rs,
	})
}

func (rh *CategoryHand) CreateCategory(c *gin.Context) {
	var cat domain.Category

	if err := c.BindJSON(&cat); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	if err := rh.uc.CreateCategory(c, cat); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
	})
}
