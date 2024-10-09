package audit

type Audit interface {
	Check(string) bool
}
