package domain

import "errors"

var (
	// ErrParams will throw if one of the parameters is missing from the request
	ErrParams = errors.New("departure/destination param missing")
	// ErrFields will throw if one of the fields are missing
	ErrFields = errors.New("all fields are required")
	// ErrFind will throw if fail to find the route
	ErrFind = errors.New("could not find the route")
	// ErrCreate will throw if fail to create the route
	ErrCreate = errors.New("could not create the route")
	// ErrConflict will throw if the roude already exists
	ErrConflict = errors.New("route already exists")
	// ErrNotFound will throw if the route is not found
	ErrNotFound = errors.New("route not found")
)
