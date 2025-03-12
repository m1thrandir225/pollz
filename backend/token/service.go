package token

import (
	"github.com/google/uuid"
	"time"
)

type TokenMaker interface {
	CreateToken(userId uuid.UUID, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
