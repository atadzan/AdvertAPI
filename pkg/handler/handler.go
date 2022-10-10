package handler

import (
	"github.com/atadzan/AdvertAPI/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()

	api := router.Group("/api", h.userIdentity)
	{
		advert := api.Group("/advert")
		{
			advert.POST("/", h.addAdvert)
			advert.DELETE("/:id", h.deleteAdvert)
			advert.PUT("/edit/:id", h.updateAdvert)
			advert.PUT("/:id", h.addFavList)
			advert.GET("fav", h.getFavList)
			advert.DELETE("fav/:id", h.deleteFav)
		}
	}
	advert := router.Group("api/advert")
	{
		advert.GET("/", h.getAdverts)
		advert.GET("/:id", h.getAdvertById)
		advert.GET("/image/:id", h.getImage)
	}
	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	return router
}