package grifts

import (
	"github.com/alexmcosta/aembassador/aembassa/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
