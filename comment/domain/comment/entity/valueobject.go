package entity

import "errors"

const (
	numAuditWaiting = iota + 1
	numAuditPassed
	numAuditReject

	auditWaiting = "audit_waiting"
	auditPassed  = "audit_passed"
	auditReject  = "audit_reject"
)

var commentStatusMap map[int]string = map[int]string{
	numAuditWaiting: auditWaiting,
	numAuditPassed:  auditPassed,
	numAuditReject:  auditReject,
}

// comment status
type CommentStatus interface {
	CommentStatus() int
	CommentStatusString() string
	IsAuditWaiting() bool
	IsAuditPassed() bool
	IsAuditReject() bool
}

type commentStatus int

func NewCommentStatus(v int) (CommentStatus, error) {

	if v != numAuditWaiting && v != numAuditPassed && v != numAuditReject {
		return nil, errors.New("invalid comment status input")
	}

	return commentStatus(v), nil
}

func NewCommentStatusWaiting() CommentStatus {
	commentStatus, _ := NewCommentStatus(numAuditWaiting)

	return commentStatus
}

func (r commentStatus) CommentStatus() int {
	return int(r)
}

func (r commentStatus) CommentStatusString() string {
	if v, ok := commentStatusMap[int(r)]; !ok {
		return ""
	} else {
		return v
	}
}

func (r commentStatus) IsAuditWaiting() bool {
	return r.CommentStatus() == numAuditWaiting
}

func (r commentStatus) IsAuditPassed() bool {
	return r.CommentStatus() == numAuditPassed
}

func (r commentStatus) IsAuditReject() bool {
	return r.CommentStatus() == numAuditReject
}
