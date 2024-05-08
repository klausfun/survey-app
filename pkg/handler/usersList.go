package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	survey "survey_app"
)

type getAllUsersResponse struct {
	Data []survey.User `json:"data"`
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.Users.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllUsersResponse{
		Data: users,
	})
}

func (h *Handler) getUserById(c *gin.Context) {

}
