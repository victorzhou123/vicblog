package infrastructure

import (
	"github.com/victorzhou123/vicblog/common/infrastructure/auditimpl"
	"github.com/victorzhou123/vicblog/common/infrastructure/authimpl"
	"github.com/victorzhou123/vicblog/common/infrastructure/kafkaimpl"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
	"github.com/victorzhou123/vicblog/common/infrastructure/oss"
)

type Config struct {
	Auth  authimpl.Config  `json:"auth"`
	Mysql mysql.Config     `json:"mysql"`
	Kafka kafkaimpl.Config `json:"kafka"`
	Oss   oss.Config       `json:"oss"`
	Audit auditimpl.Config `json:"audit"`
}
