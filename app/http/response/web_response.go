package response

import "github.com/goravel/framework/contracts/http"

func ApiResponse(writer http.Context, statusCode int, response interface{}, messages string) http.Response {
	var status string
	status = "success"
	if statusCode != 200 {
		status = "failed"
	}
	if statusCode == 200 && messages == "" {
		messages = "success"
		status = "success"
	} else if statusCode == 201 && messages == "" {
		messages = "created"
	} else if statusCode == 400 && messages == "" {
		messages = "bad request"
	} else if statusCode == 401 && messages == "" {
		messages = "unauthorize"
	} else if statusCode == 404 && messages == "" {
		messages = "not found"
	} else if statusCode == 500 && messages == "" {
		messages = "internal server error"
	}
	return writer.Response().Json(statusCode, http.Json{
		"status":   status,
		"messages": messages,
		"data":     response,
	})
}
