package routes

import (
	"gin-postgres/controllers"

	"github.com/gin-gonic/gin"
)

type ContactRoutes struct {
	contactController controllers.ContactController
}

func NewRouteContact(contactController controllers.ContactController) ContactRoutes {
	return ContactRoutes{contactController}
}

func (cr *ContactRoutes) ContactRoute(rg *gin.RouterGroup) {
	router := rg.Group("contacts")
	router.POST("/", cr.contactController.CreateContact)
	router.GET("/", cr.contactController.GetAllContacts)
	router.GET("/:id", cr.contactController.GetContactById)
	router.PUT("/:id", cr.contactController.UpdateContact)
	router.DELETE("/:id", cr.contactController.DeleteContactById)
}
