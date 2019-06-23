package actions

import "github.com/gobuffalo/buffalo"

// decks page.
func DecksHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("decks.html"))
}
