package models

import "errors"

// Error .
type Error struct {
	Err  string `json:"error"`
	Code int    `json:"status_code"`
}

// Define error
var (
	ErrorInvalidID      = errors.New("Invalid ID")
	ErrorInvalidProject = errors.New("Invalid project")
)
