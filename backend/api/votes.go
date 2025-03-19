package api

import (
	db "m1thrandir225/cicd2025/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateVoteRequest struct {
	OptionID  string `json:"option_id"`
	IpAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
}

type UpdateVoteRequest struct {
	OptionID string `json:"option_id"`
}

func (server *Server) createVote(ctx *gin.Context) {
	var req CreateVoteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	optionId, err := uuid.Parse(req.OptionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateVoteParams{
		OptionID:  optionId,
		IpAddress: req.IpAddress,
		UserAgent: req.UserAgent,
	}

	newVote, err := server.store.CreateVote(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, newVote)
}

func (server *Server) updateVote(ctx *gin.Context) {
	var uriId UriID
	var req UpdateVoteRequest

	if err := ctx.ShouldBindUri(&uriId); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	voteId, err := uuid.Parse(uriId.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	optionId, err := uuid.Parse(req.OptionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.UpdateVoteOptionParams{
		ID:       voteId,
		OptionID: optionId,
	}

	updated, err := server.store.UpdateVoteOption(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

func (server *Server) deleteVote(ctx *gin.Context) {
	var uriId UriID

	if err := ctx.ShouldBindUri(&uriId); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	voteId, err := uuid.Parse(uriId.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.store.DeleteVote(ctx, voteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}
