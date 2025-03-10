package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "m1thrandir225/cicd2025/db/sqlc"
	"net/http"
)

type CreatePollOptionRequest struct {
	PollID     string `json:"poll_id"`
	OptionText string `json:"option_text"`
}

type UpdatePollOptionRequest struct {
	OptionText string `json:"option_text"`
}

type DeletePollOptionRequest struct {
	UserID string `json:"user_id"`
}

func (server *Server) createOption(ctx *gin.Context) {
	var req CreatePollOptionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	pollId, err := uuid.Parse(req.PollID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreatePollOptionParams{
		PollID:     pollId,
		OptionText: req.OptionText,
	}

	option, err := server.store.CreatePollOption(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, option)
}

func (server *Server) updateOption(ctx *gin.Context) {
	var uriId UriID
	var req UpdatePollOptionRequest

	if err := ctx.ShouldBindUri(&uriId); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	optionId, err := uuid.Parse(uriId.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.UpdatePollOptionParams{
		ID:         optionId,
		OptionText: req.OptionText,
	}

	updated, err := server.store.UpdatePollOption(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

func (server *Server) deleteOption(ctx *gin.Context) {
	var uriId UriID
	var req DeletePollOptionRequest
	if err := ctx.ShouldBindUri(&uriId); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	optionId, err := uuid.Parse(uriId.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	userId, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	option, err := server.store.GetOption(ctx, optionId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	poll, err := server.store.GetPoll(ctx, option.PollID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if poll.CreatedBy != userId {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_, err = server.store.DeletePollOption(ctx, optionId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.Status(http.StatusOK)
}
