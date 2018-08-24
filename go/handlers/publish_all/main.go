package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gemcook/playground-sls/go/models"
	"github.com/gemcook/playground-sls/go/services/publish"
	"github.com/gemcook/playground-sls/go/utils"
)

// Handler handles
func Handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data models.Publish

	err := json.Unmarshal([]byte(event.Body), &data)
	if err != nil {
		return utils.NewErrorResponse(err, http.StatusBadRequest)
	}

	log.Println(data.Message)
	topicName, err := publish.PublishAll(&data)
	if err != nil {
		return utils.NewErrorResponse(err, http.StatusBadRequest)
	}

	return utils.NewResponse(`{"message":"`+*topicName+`"}`, http.StatusOK)
}

func main() {
	lambda.Start(Handler)
}
