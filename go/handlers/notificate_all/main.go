package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gemcook/playground-sls/go/models"
	"github.com/gemcook/playground-sls/go/utils"
)

// Handler handles
func Handler(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var data models.Notification

	err := json.Unmarshal([]byte(event.Body), &data)
	if err != nil {
		return utils.NewErrorResponse(err, http.StatusBadRequest)
	}

	sess, err := session.NewSession()
	if err != nil {
		panic(err)
	}

	svc := sns.New(sess)

	// トピックを作成する

	topicName := "tfrcm-push-notification-all"

	topicInput := &sns.CreateTopicInput{
		Name: &topicName,
	}

	topicARN, err := svc.CreateTopic(topicInput)
	if err != nil {
		panic(err)
	}

	log.Println(topicARN)
	log.Println(`{"message":"` + topicName + `"}`)

	// TODO: targetUsersのメールアドレスでApplicationのエンドポイントARNを取得する。
	// applicationInput := &sns.ListPlatformApplicationsInput{}
	// applicationLists, error := svc.ListEndpointsByPlatformApplication()

	// log.Println(applicationLists)

	// TODO: 取得したApplicationのエンドポイントARNにtopicARNをサブスクライブさせる。

	// TODO: m_messagesからsend-tommorowのメッセージを取得する。

	// TODO: メッセージをPublishする。
	message := "Hello"
	tArn := "arn:aws:sns:ap-northeast-1:598003641956:tfrcm-push-notification-all"
	publishInput := &sns.PublishInput{
		Message:  &message,
		TopicArn: &tArn,
	}

	publishResult, err := svc.Publish(publishInput)
	if err != nil {
		panic(err)
	}

	log.Println(publishResult)

	return utils.NewResponse(`{"message":"`+topicName+`"}`, http.StatusOK)
}

func main() {
	lambda.Start(Handler)
}
