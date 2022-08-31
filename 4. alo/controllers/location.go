package controllers 

import (
	"github.com/gin-gonic/gin"
	
	"backend/service"
	"net/http"
)

func GetCountry(c *gin.Context) {
	country, err := service.LocationService.GetCountry(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, country)
}