package models

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// Error .
type Error struct {
	Err  string `json:"error"`
	Code int    `json:"status_code"`
}

// Define error
var (
	ErrorInvalidID         = errors.New("Invalid ID")
	ErrorInvalidProject    = errors.New("Invalid project")
	ErrorInvalidName       = errors.New("Invalid name")
	ErrorInvalidPhone      = errors.New("Invailid phone")
	ErrorInvalidMember     = errors.New("Invalid member")
	ErrorInvalidAssignment = errors.New("Invalid assignment")
)

// HandleError Log and ouput error
func HandleError(err error, errorType error, code int, outputWriter http.ResponseWriter) {
	log.Println(err)

	e := Error{
		Err:  errorType.Error(),
		Code: code,
	}

	b, _ := json.Marshal(e)
	outputWriter.Write(b)
}
