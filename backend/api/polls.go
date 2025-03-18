package api

import (
	"context"
	"errors"
	db "m1thrandir225/cicd2025/db/sqlc"
	"m1thrandir225/cicd2025/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CreatePollRequest struct {
	Description string   `json:"description" binding:"required"`
	ActiveUntil string   `json:"active_until" binding:"required"`
	Options     []string `json:"options" binding:"required"`
}

type GetPollResponse struct {
	Poll    db.Poll         `json:"poll"`
	Options []db.PollOption `json:"options"`
}

type UpdatePollRequest struct {
	Description string `json:"description" binding:"required"`
	Status      *bool  `json:"status" binding:"required"`
	ActiveUntil string `json:"active_until" binding:"required"`
}

type UpdatePollStatusRequest struct {
	ActiveUntil string `json:"active_until" binding:"required"`
}

func (server *Server) createPoll(ctx *gin.Context) {
	var req CreatePollRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	activeTime, err := time.Parse(UtcDateFormat, req.ActiveUntil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("active_until format mismatch")))
		return
	}

	/**
	* The active time for the new poll cannot be in the past
	 */

	if !util.DateBefore(time.Now().UTC(), activeTime.UTC()) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "poll active time must be in the future"})
		return
	}
	/**
	* Get the userID from the current active payload context
	 */
	payload, err := getPayloadFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	args := db.CreatePollParams{
		Description: req.Description,
		CreatedBy:   payload.UserId,
		ActiveUntil: activeTime,
	}

	result, err := server.store.CreatePoll(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	optionArgs := db.CreateMultipleOptionsParams{
		PollID:      result.ID,
		OptionTexts: req.Options,
	}

	options, err := server.store.CreateMultipleOptions(ctx, optionArgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := GetPollResponse{
		Poll:    result,
		Options: options,
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) getPolls(ctx *gin.Context) {
	payload, err := getPayloadFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	polls, err := server.store.GetPolls(ctx, payload.UserId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
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
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	options, err := server.store.GetOptionsForPoll(context.Background(), poll.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
		}
	}

	result := GetPollResponse{
		Options: options,
		Poll:    poll,
	}

	ctx.JSON(http.StatusOK, result)
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

	activeTime, err := time.Parse(UtcDateFormat, req.ActiveUntil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	/**
	* The new active time for a poll cannot be before the current time of updating,
	* i.e. has to be in the future
	 */
	if util.DateBefore(time.Now(), activeTime) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "poll active time must be in the future"})
		return
	}

	/**
	* Verify if the current poll is an active one. Only active polls can be updated
	 */
	pollData, err := server.store.IsPollActive(ctx, pollId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "poll is not active "})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	/**
	 * Only the user that has created the poll can update its contents
	 */
	err = verifyUserIdWithToken(pollData.CreatedBy, ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	args := db.UpdatePollParams{
		ID:          pollId,
		Description: req.Description,
		ActiveUntil: activeTime,
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
	activeTime, err := time.Parse(UtcDateFormat, req.ActiveUntil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	/**
	* The new active time for a poll cannot be before the current time of updating,
	* i.e. has to be in the future
	 */
	if util.DateBefore(time.Now(), activeTime) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "poll active time must be in the future"})
		return
	}

	/**
	* Verify if the current poll is an active one. Only active polls can be updated
	 */
	pollData, err := server.store.IsPollActive(ctx, pollId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "poll is not active "})
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	/**
	 * Only the user that has created the poll can update its contents
	 */
	err = verifyUserIdWithToken(pollData.CreatedBy, ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	args := db.UpdatePollStatusParams{
		ID:          pollId,
		ActiveUntil: activeTime,
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

	/**
	 * Only the user that has created the poll can update its contents
	 */
	err = verifyUserIdWithToken(poll.CreatedBy, ctx)
	if err != nil {
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
