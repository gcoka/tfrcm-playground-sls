package publish

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/gemcook/playground-sls/go/models"
)

// PublishAll is SNS Publish Function
func PublishAll(data *models.Publish) (*string, error) {
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

	topicArnOutput, err := svc.CreateTopic(topicInput)
	if err != nil {
		return nil, err
	}

	// TODO: targetUsersのメールアドレスでApplicationのエンドポイントARNを取得する。
	// applicationInput := &sns.ListPlatformApplicationsInput{}
	// applicationLists, error := svc.ListEndpointsByPlatformApplication()

	// log.Println(applicationLists)

	// TODO: 取得したApplicationのエンドポイントARNにtopicARNをサブスクライブさせる。

	// TODO: m_messagesからsend-tommorowのメッセージを取得する。

	// TODO: メッセージをPublishする。
	message := data.Message

	publishInput := &sns.PublishInput{
		Message:  &message,
		TopicArn: topicArnOutput.TopicArn,
	}

	_, err = svc.Publish(publishInput)
	if err != nil {
		return nil, err
	}

	return &topicName, nil
}
