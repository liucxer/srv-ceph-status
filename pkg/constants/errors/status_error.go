package errors

import "net/http"

//go:generate codegen gen error StatusError
type StatusError int

func (StatusError) ServiceCode() int {
	return 801 * 1e3
}

// 400
const (
	StatusBadRequestError StatusError = http.StatusBadRequest*1e6 + iota + 1
)

// 403
const (
	// Forbidden
	ForbiddenError StatusError = http.StatusForbidden*1e6 + iota + 1
)

// 404
const (
	// NotFound
	NotFoundError StatusError = http.StatusNotFound*1e6 + iota + 1
)

// 409
const (
	// Conflict
	ConflictError StatusError = http.StatusConflict*1e6 + iota + 1
)

// 500
const (
	// InternalServerError
	InternalServerError StatusError = http.StatusInternalServerError*1e6 + iota + 1
)
