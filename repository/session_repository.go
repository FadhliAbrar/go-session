package repository

import (
	"context"
	"github.com/FadhliAbrar/golang-session/model"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, session model.Session)
	FindSessionUserId(ctx context.Context, session string) model.User
}
