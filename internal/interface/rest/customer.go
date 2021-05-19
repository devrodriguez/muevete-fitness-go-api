package rest

import (
	"github.com/devrodriguez/muevete-fitness-go-api/internal/customers"
	"github.com/devrodriguez/muevete-fitness-go-api/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomerHand struct {
	uc customers.ICustomerCrud
}

func NewCustomerHand(uc customers.ICustomerCrud) CustomerHand {
	return CustomerHand{
		uc,
	}
}

func (ch *CustomerHand) GetAllCustomers(c *gin.Context) {

	cs, err := ch.uc.GetAllCustomers(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
		Data:    cs,
	})
}

func (ch *CustomerHand) CreateCustomer(c *gin.Context) {
	var cus domain.Customer

	if err := c.BindJSON(&cus); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	if err := ch.uc.CreateCustomer(c, cus); err !=  nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: http.StatusText(http.StatusOK),
	})
}


