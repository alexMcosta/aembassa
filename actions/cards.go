package actions

import "github.com/gobuffalo/buffalo"

// Cards Page.
func CardsHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("cards.html"))
}
