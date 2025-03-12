package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "m1thrandir225/cicd2025/db/sqlc"
	"net/http"
)

type CreatePollRequest struct {
	Description string `json:"description" binding:"required"`
	UserID      string `json:"user_id" binding:"required,uuid"`
}

type UpdatePollRequest struct {
	Description string `json:"description" binding:"required"`
	Status      *bool  `json:"status" binding:"required"`
}

type UpdatePollStatusRequest struct {
	Status *bool `json:"status" binding:"required"`
}

type DeletePollRequest struct {
	UserID string `json:"user_id" binding:"required,uuid"`
}

func (server *Server) createPoll(ctx *gin.Context) {
	var req CreatePollRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userId, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreatePollParams{
		Description: req.Description,
		CreatedBy:   userId,
	}

	result, err := server.store.CreatePoll(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) getPolls(ctx *gin.Context) {
	polls, err := server.store.GetPolls(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, polls)
}

func (server *Server) getPoll(ctx *gin.Context) {
	var uriId UriID
	if err := ctx.ShouldBindUri(&uriId); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pollId, err := uuid.Parse(uriId.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	poll, err := server.store.GetPoll(ctx, pollId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, poll)
}

func (server *Server) updatePoll(ctx *gin.Context) {
	var uriId UriID
	var req UpdatePollRequest

	if err := ctx.ShouldBindUri(&uriId); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pollId, err := uuid.Parse(uriId.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.UpdatePollParams{
		ID:          pollId,
		Description: req.Description,
		IsActive:    *req.Status,
	}

	updated, err := server.store.UpdatePoll(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updated)

}

func (server *Server) updatePollStatus(ctx *gin.Context) {
	var uriId UriID
	var req UpdatePollStatusRequest

	if err := ctx.ShouldBindUri(&uriId); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return

	}

	pollId, err := uuid.Parse(uriId.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.UpdatePollStatusParams{
		ID:       pollId,
		IsActive: *req.Status,
	}

	updated, err := server.store.UpdatePollStatus(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

func (server *Server) deletePoll(ctx *gin.Context) {
	var uriId UriID
	var req DeletePollRequest
	if err := ctx.ShouldBindUri(&uriId); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pollId, err := uuid.Parse(uriId.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userId, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	poll, err := server.store.GetPoll(ctx, pollId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if poll.CreatedBy != userId {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_, err = server.store.DeletePoll(ctx, pollId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
