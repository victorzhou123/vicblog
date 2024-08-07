package util

import "text/template"

func XssEscape(v string) string {
	return template.HTMLEscapeString(v)
}
