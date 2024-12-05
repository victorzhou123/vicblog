package event

import (
	"encoding/json"

	"github.com/victorzhou123/vicblog/common/domain/mq"
)

func ToMessage(data any) (mq.Message, error) {

	b, err := json.Marshal(data)
	if err != nil {
		return mq.Message{}, err
	}

	msg := mq.Message{
		Body: b,
	}

	return msg, nil
}
