package response

import "net/http"

const (
	StatusOK                  = http.StatusOK                  // 200
	StatusCreated             = http.StatusCreated             // 201
	StatusNoContent           = http.StatusNoContent           // 204
	StatusBadRequest          = http.StatusBadRequest          // 400
	StatusUnauthorized        = http.StatusUnauthorized        // 401
	StatusForbidden           = http.StatusForbidden           // 403
	StatusNotFound            = http.StatusNotFound            // 404
	StatusConflict            = http.StatusConflict            // 409
	StatusUnprocessableEntity = http.StatusUnprocessableEntity // 422
	StatusInternalServerError = http.StatusInternalServerError // 500
)
