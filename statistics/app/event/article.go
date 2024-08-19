package event

import (
	"encoding/json"

	cmappevent "github.com/victorzhou123/vicblog/common/app/event"
	cmevent "github.com/victorzhou123/vicblog/common/domain/event"
	"github.com/victorzhou123/vicblog/common/log"
	statsdmsvc "github.com/victorzhou123/vicblog/statistics/domain/service"
)

const (
	topicAddArticleReadTimes = cmappevent.TopicAddArticleReadTimes

	fieldArticleId = cmappevent.FieldArticleId
)

type articleVisitsSubscriber struct {
	articleVisits statsdmsvc.ArticleVisitsService
}

func NewArticleVisitsSubscriber(
	articleVisits statsdmsvc.ArticleVisitsService,
) cmevent.Subscriber {
	return &articleVisitsSubscriber{articleVisits}
}

func (s *articleVisitsSubscriber) Consume(e cmevent.Event) {

	var body map[string]string

	if err := json.Unmarshal(e.Message().Body, &body); err != nil {
		log.Errorf("article visits subscriber unmarshal message failed, err: %s", err.Error())

		return
	}

	articleId, ok := body[fieldArticleId]
	if !ok {
		log.Errorf("article visits subscriber field `article` cannot be found")

		return
	}

	if articleId == "" {
		return
	}

	if err := s.articleVisits.IncreaseOneVisitsOfToday(); err != nil {
		log.Errorf("increase visits of today failed, err: %s", err.Error())
	}
}

func (s *articleVisitsSubscriber) Topics() []string {
	return []string{topicAddArticleReadTimes}
}
