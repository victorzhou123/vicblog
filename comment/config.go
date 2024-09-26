package comment

import "github.com/victorzhou123/vicblog/comment/infrastructure/qqinfoimpl"

type Config struct {
	QQInfo qqinfoimpl.Config `json:"qq_info"`
}
