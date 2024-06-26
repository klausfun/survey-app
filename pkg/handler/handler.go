package handler

import (
	"github.com/gin-gonic/gin"
	"survey_app/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		surveys := api.Group("/surveys")
		{
			surveys.GET("/", h.getAllSurveys)
			surveys.GET("/:id", h.getSurveyById)
			surveys.PUT("/:id", h.updateSurvey)
		}

		admin := api.Group("/admin", h.adminIdentity)
		{
			survey := admin.Group("/surveys")
			{
				survey.POST("/", h.createSurvey)
				survey.GET("/", h.getAllSurveys)
				survey.GET("/:id", h.getSurveyById)
				survey.DELETE("/:id", h.deleteSurvey)
			}

			usersList := admin.Group("/usersList")
			{
				usersList.GET("/", h.getAllUsers)
				usersList.GET("/:id", h.getUserById)
			}
		}
	}

	return router
}
