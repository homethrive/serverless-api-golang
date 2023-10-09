package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

// GetUser intermediates dataflow between datasource (dynamodb) and api response
func Authorize(req events.APIGatewayProxyRequest) (
	*events.APIGatewayProxyResponse, error,
) {
	// token := req.event.AuthorizationToken
	
	return apiResponse(http.StatusOK, nil)
}