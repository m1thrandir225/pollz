package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "m1thrandir225/cicd2025/db/sqlc"
	"net/http"
)

type CreateVoteRequest struct {
	OptionID  string `json:"option_id"`
	UserID    string `json:"user_id"`
	IpAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
}

type UpdateVoteRequest struct {
	OptionID string `json:"option_id"`
}

type DeleteVoteRequest struct {
	UserID string `json:"user_id"`
}

func (server *Server) CreateVote(ctx *gin.Context) {
	var req CreateVoteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	optionId, err := uuid.Parse(req.OptionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userId, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateVoteParams{
		OptionID:  optionId,
		UserID:    userId,
		IpAddress: req.IpAddress,
		UserAgent: req.UserAgent,
	}

	newVote, err := server.store.CreateVote(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, newVote)
}

func (server *Server) UpdateVote(ctx *gin.Context) {
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

func (server *Server) DeleteVote(ctx *gin.Context) {
	var uriId UriID
	var req DeleteVoteRequest

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

	userId, err := uuid.Parse(req.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	vote, err := server.store.GetVote(ctx, voteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if vote.UserID != userId {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	_, err = server.store.DeleteVote(ctx, voteId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Status(http.StatusOK)

}
