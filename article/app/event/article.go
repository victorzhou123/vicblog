package event

import (
	"encoding/json"
	"errors"

	articledmsvc "github.com/victorzhou123/vicblog/article/domain/article/service"
	cmappevent "github.com/victorzhou123/vicblog/common/app/event"
	"github.com/victorzhou123/vicblog/common/domain/mq"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/common/log"
)

const (
	topicAddArticleReadTimes = cmappevent.TopicAddArticleReadTimes

	fieldArticleId = cmappevent.FieldArticleId
)

type articleSubscriber struct {
	article articledmsvc.ArticleService
}

func NewArticleSubscriber(
	article articledmsvc.ArticleService,
) *articleSubscriber {
	return &articleSubscriber{article}
}

func (s *articleSubscriber) Consume(e mq.Event) error {

	var body map[string]string

	if err := json.Unmarshal(e.Message().Body, &body); err != nil {
		log.Errorf("article subscriber unmarshal message failed, err: %s", err.Error())

		return err
	}

	articleId, ok := body[fieldArticleId]
	if !ok {
		log.Errorf("article subscriber field `article` cannot be found")

		return errors.New("article subscriber field")
	}

	amountOne, err := cmprimitive.NewAmount(1)
	if err != nil {
		return err
	}

	if err := s.article.AddArticleReadTimes(cmprimitive.NewId(articleId), amountOne); err != nil {
		log.Errorf("article %s AddArticleReadTimes failed, err: %s", articleId, err.Error())
	}

	return nil
}

func (s *articleSubscriber) Topics() []string {
	return []string{topicAddArticleReadTimes}
}
