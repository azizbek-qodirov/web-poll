package handlers

import (
	pb "auth-service/genprotos"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateResult godoc
// @Summary Create a new result
// @Description Create a new result for a poll
// @Tags results
// @Accept json
// @Produce json
// @Param CreateResultReq body pb.CreateResultReq true "Create Result Request"
// @Success 201 {object} pb.Void
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /results [post]
// @Security BearerAuth
func (h *HTTPHandler) CreateResult(c *gin.Context) {
	var req pb.CreateResultReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	_, err := h.Result.CreateResult(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server error: "+err.Error())
		return
	}

	c.JSON(http.StatusCreated, &pb.Void{})
}

// SavePollAnswer godoc
// @Summary Save a poll answer
// @Description Save the answer for a specific question in a poll
// @Tags results
// @Accept json
// @Produce json
// @Param SavePollAnswerReq body pb.SavePollAnswerReq true "Save Poll Answer Request"
// @Success 200 {object} pb.Void
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /poll_answers [post]
// @Security BearerAuth
func (h *HTTPHandler) SavePollAnswer(c *gin.Context) {
	var req pb.SavePollAnswerReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	_, err := h.Result.SavePollAnswer(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, &pb.Void{})
}

// GetPoll godoc
// @Summary Get all polls for a user
// @Description Retrieve all polls associated with a user
// @Tags results
// @Accept json
// @Produce json
// @Param GetPollReq body pb.GetPollReq true "Get Poll Request"
// @Success 200 {object} pb.GetPollRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /polls [get]
// @Security BearerAuth
func (h *HTTPHandler) GetPoll(c *gin.Context) {
	var req pb.GetPollReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}

	res, err := h.Result.GetPoll(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Server error: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
