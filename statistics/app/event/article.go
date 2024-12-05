package event

import (
	"encoding/json"
	"errors"

	cmappevent "github.com/victorzhou123/vicblog/common/app/event"
	"github.com/victorzhou123/vicblog/common/domain/mq"
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
) *articleVisitsSubscriber {
	return &articleVisitsSubscriber{articleVisits}
}

func (s *articleVisitsSubscriber) Consume(e mq.Event) error {

	var body map[string]string

	if err := json.Unmarshal(e.Message().Body, &body); err != nil {
		log.Errorf("article visits subscriber unmarshal message failed, err: %s", err.Error())

		return err
	}

	articleId, ok := body[fieldArticleId]
	if !ok {
		log.Errorf("article visits subscriber field `article` cannot be found")

		return errors.New("article visits subscriber field")
	}

	if articleId == "" {
		return errors.New("article id not found")
	}

	if err := s.articleVisits.IncreaseOneVisitsOfToday(); err != nil {
		log.Errorf("increase visits of today failed, err: %s", err.Error())
		return err
	}

	return nil
}

func (s *articleVisitsSubscriber) Topics() []string {
	return []string{topicAddArticleReadTimes}
}
