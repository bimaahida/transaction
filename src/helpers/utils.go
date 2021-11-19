package helpers

import (
	"log"
	"net/http"

	"transaction/src/definitions"
)

type ResponseError struct {
	Message string `json:"message"`
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Fatal(err)
	switch err {
	case definitions.ErrInternalServerError:
		return http.StatusInternalServerError
	case definitions.ErrNotFound:
		return http.StatusNotFound
	case definitions.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
