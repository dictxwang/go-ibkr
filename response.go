package ibkr

import "errors"

var (
	// ErrBadRequest :
	// Need to send the request with GET / POST (must be capitalized)
	ErrBadRequest = errors.New("bad request")
	// ErrInvalidRequest :
	// 1. Need to use the correct key to access;
	// 2. Need to put authentication params in the request header
	ErrInvalidRequest = errors.New("authentication failed")
	// ErrForbiddenRequest :
	// Possible causes:
	// 1. IP rate limit breached;
	// 2. You send GET request with an empty json body;
	// 3. You are using U.S IP
	ErrForbiddenRequest = errors.New("access denied")
	// ErrPathNotFound :
	// Possible causes:
	// 1. Wrong path;
	// 2. Category value does not match account mode
	ErrPathNotFound = errors.New("path not found")
)
