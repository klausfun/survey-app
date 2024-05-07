package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

type getAllSurveysResponse struct {
	Data []survey.Surveys `json:"data"`
}

func (h *Handler) getAllSurveys(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	surveys, err := h.services.Surveys.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllSurveysResponse{
		Data: surveys,
	})
}

func (h *Handler) getSurveyById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	surveyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid is param")
		return
	}

	sur, err := h.services.Surveys.GetById(userId, surveyId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, sur)
}

func (h *Handler) updateSurvey(c *gin.Context) {

}

func (h *Handler) deleteSurvey(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	surveyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid is param")
		return
	}

	err = h.services.Surveys.Delete(userId, surveyId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}
