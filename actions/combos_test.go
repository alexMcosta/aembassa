package actions

func (as *ActionSuite) Test_CombosHandler() {
	res := as.HTML("/combos").Get()

	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}
