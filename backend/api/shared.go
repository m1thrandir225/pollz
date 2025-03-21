package api

import (
	"errors"
	"m1thrandir225/cicd2025/token"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const UtcDateFormat = "2006-01-02T15:04:05.000Z"

type UriID struct {
	ID string `uri:"id" binding:"required"`
}

func getPayloadFromContext(c *gin.Context) (*token.Payload, error) {
	data, exists := c.Get(authorizationPayloadKey)
	if !exists {
		return nil, errors.New("authorization header not found")
	}
	payload := data.(*token.Payload)
	return payload, nil
}

func verifyUserIdWithToken(userId uuid.UUID, ctx *gin.Context) error {
	data, exists := ctx.Get(authorizationPayloadKey)
	if !exists {
		return errors.New(authorizationPayloadKey + " was not found in the request context")
	}
	payload := data.(*token.Payload)

	if payload.UserId != userId {
		return errors.New("context payload doesn't match with header payload")
	}
	return nil
}
