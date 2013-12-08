package main

import (
	"html/template"
	"strings"

	"code.google.com/p/go.tools/present"
)

func init() {
	present.Register("blockquote", parseBlockquote)
}

type Blockquote struct {
	Body template.HTML
}

func (b Blockquote) TemplateName() string { return "blockquote" }

func parseBlockquote(ctx *present.Context, fileName string, lineno int, text string) (present.Elem, error) {
	body := strings.Join(strings.Fields(text)[1:], " ")
	return Blockquote{Body: template.HTML(body)}, nil
}
