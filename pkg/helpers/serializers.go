package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/io-m/app-hyphen/internal/shared/types"
	"github.com/io-m/app-hyphen/pkg/constants"
)

// Decodes received payload from the client. Allows up to 1MB of payload size
func DecodePayload[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var payload T
	r.Body = http.MaxBytesReader(w, r.Body, constants.MAX_SIZE_PAYLOAD)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return payload, errors.New("failed to read request body")
	}
	// Unmarshal the request body into a payload struct
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return payload, errors.New("failed to unmarshal request body")
	}
	return payload, nil
}

// Encodes payload received, sets headers and ships it to the client
func EncodePayload(w http.ResponseWriter, payload map[constants.ResponseMapKey]types.JSONResponse, status int, headers ...http.Header) {
	bytes, _ := json.Marshal(payload)
	if len(headers) > 0 {
		for key, val := range headers[0] {
			w.Header()[key] = val
		}
	}
	w.Header().Set(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	w.WriteHeader(status)
	_, _ = w.Write(bytes)
}

// Returns error JSON by signifying Failure field as true and setting error message and status code
func ErrorResponse(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	var payload types.JSONResponse
	payload.Code = statusCode
	payload.Message = err.Error()
	response := make(map[constants.ResponseMapKey]types.JSONResponse)
	response[constants.ERROR] = payload
	EncodePayload(w, response, statusCode)
}

// Returns success JSON by signifying Failure field as false and setting data field. Default status code is 200
func SuccessResponse(w http.ResponseWriter, data any, message string, status ...int) {
	statusCode := http.StatusOK
	if len(status) > 0 {
		statusCode = status[0]
	}
	var payload types.JSONResponse
	payload.Code = statusCode
	payload.Message = message
	payload.Data = data
	response := make(map[constants.ResponseMapKey]types.JSONResponse)
	response[constants.SUCCESS] = payload
	EncodePayload(w, response, statusCode)
}
