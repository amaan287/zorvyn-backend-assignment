package constants

import "errors"

var (
	ErrMissingENV          = errors.New("Missing required Environment variables")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
	ErrUserNotFound        = errors.New("user not found")
	ErrInvalidToken        = errors.New("invalid token")
	ErrNilUser             = errors.New("user is nil")
	ErrUserIdRequired      = errors.New("user id is required")
	ErrNoPassword          = errors.New("password is required")
	ErrNilTokenManager     = errors.New("token manager is nil")
)
