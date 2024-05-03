package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	survey "survey_app"
)

func (h *Handler) createSurvey(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input survey.Data
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Surveys.CreateSurvey(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllSurveys(c *gin.Context) {
}

func (h *Handler) getSurveyById(c *gin.Context) {

}

func (h *Handler) updateSurvey(c *gin.Context) {

}

func (h *Handler) deleteSurvey(c *gin.Context) {

}
