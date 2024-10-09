package auditimpl

import (
	"github.com/victorzhou123/sensitive-word-check/client"

	"github.com/victorzhou123/vicblog/common/domain/audit"
)

type auditimpl struct {
	cli client.SensitiveWordChecker
}

func NewAuditImpl(cfg *Config) (audit.Audit, error) {

	cli, err := client.NewSensitiveWordCheckClient(cfg.Addr, cfg.expireTime())
	if err != nil {
		return nil, err
	}

	return &auditimpl{
		cli: cli,
	}, nil
}

func (a *auditimpl) Check(word string) bool {
	return a.cli.Check(word)
}
