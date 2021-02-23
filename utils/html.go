package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

func FormatHtml(html string, data interface{}) string {
	messageHTML := ""
	funcMap := template.FuncMap{
		"attr": func(s string) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
		"safeURL": func(u string) template.URL {
			return template.URL(u)
		},
		"dateSlice": func(s time.Time, i, j int) string {
			return s.String()[i:j]
		},
	}
	t := template.Must(template.New("email").Funcs(funcMap).Parse(html))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		fmt.Println(err)

	} else {
		messageHTML = buf.String()
	}

	return messageHTML
}
