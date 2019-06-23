package actions

func (as *ActionSuite) Test_CardsHandler() {
	res := as.HTML("/cards").Get()

	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}
