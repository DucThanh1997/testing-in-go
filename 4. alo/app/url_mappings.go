package app

import (
	"backend/controllers"
)

func mapUrls() {
	router.GET("/locations/:id", controllers.GetCountry)
}