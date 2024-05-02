package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createSurvey(c *gin.Context) {

}

func (h *Handler) getAllSurveys(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getSurveyById(c *gin.Context) {

}

func (h *Handler) updateSurvey(c *gin.Context) {

}

func (h *Handler) deleteSurvey(c *gin.Context) {

}
