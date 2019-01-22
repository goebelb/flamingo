package store

import (
	"flamingo.me/flamingo/v3/core/auth/domain"
)

type (
	Store interface {
		DestroySessionsForUser(user domain.User) error
		SetHashAndSessionIdForUser(user domain.User, hash string, id string) error
		GetHashByUser(user domain.User) (string, error)
	}
)
