package actions

import "github.com/gobuffalo/buffalo"

// Combos page.
func CombosHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("combos.html"))
}
