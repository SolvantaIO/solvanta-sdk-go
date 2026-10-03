package solvanta

import "fmt"

// Error is returned when the API responds with a non-success status.
type Error struct {
	// StatusCode is the HTTP status code.
	StatusCode int
	// Code is the machine-readable error code from the response body.
	Code string
	// Message is the human-readable detail from the response body.
	Message string
}

func (e *Error) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("solvanta: %d %s: %s", e.StatusCode, e.Code, e.Message)
	}
	return fmt.Sprintf("solvanta: HTTP %d", e.StatusCode)
}

// Common error codes returned by the API.
const (
	ErrUnauthorized              = "unauthorized"
	ErrAccountSuspended          = "account_suspended"
	ErrHostNotAllowed            = "host_not_allowed"
	ErrNotFound                  = "not_found"
	ErrValidation                = "validation_error"
	ErrProxyInvalid              = "proxy_invalid"
	ErrInsufficientCredits       = "insufficient_credits"
	ErrConcurrencyLimitExceeded  = "concurrency_limit_exceeded"
	ErrRateLimited               = "rate_limited"
	ErrSolveFailed               = "solve_failed"
	ErrSolveTimeout              = "solve_timeout"
	ErrSolverUnavailable         = "solver_unavailable"
	ErrNoCapacity                = "no_capacity"
	ErrIdempotencyConflict       = "idempotency_conflict"
	ErrInternalError             = "internal_error"
)
