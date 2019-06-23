package actions

func (as *ActionSuite) Test_DecksHandler() {
	res := as.HTML("/decks").Get()

	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}
