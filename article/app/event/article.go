package event

import (
	"encoding/json"

	articledmsvc "github.com/victorzhou123/vicblog/article/domain/article/service"
	cmevent "github.com/victorzhou123/vicblog/common/domain/event"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/common/log"
)

const (
	topicAddArticleReadTimes = "get_article"

	fieldArticleId = "articleId"
)

type articleSubscriber struct {
	article articledmsvc.ArticleService
}

func NewArticleSubscriber(
	article articledmsvc.ArticleService,
) cmevent.Subscriber {
	return &articleSubscriber{article}
}

func (s *articleSubscriber) Consume(e cmevent.Event) {

	var body map[string]string

	if err := json.Unmarshal(e.Message().Body, &body); err != nil {
		log.Errorf("article subscriber unmarshal message failed, err: %s", err.Error())

		return
	}

	articleId, ok := body[fieldArticleId]
	if !ok {
		log.Errorf("article subscriber field `article` cannot be found")

		return
	}

	amountOne, err := cmprimitive.NewAmount(1)
	if err != nil {
		return
	}

	if err := s.article.AddArticleReadTimes(cmprimitive.NewId(articleId), amountOne); err != nil {
		log.Errorf("article %s AddArticleReadTimes failed, err: %s", articleId, err.Error())
	}
}

func (s *articleSubscriber) Topics() []string {
	return []string{topicAddArticleReadTimes}
}
