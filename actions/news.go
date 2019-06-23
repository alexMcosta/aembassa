package actions

import "github.com/gobuffalo/buffalo"

// News page.
func NewsHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("news.html"))
}
