package actions

func (as *ActionSuite) Test_NewsHandler() {
	res := as.HTML("/news").Get()

	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Welcome to Buffalo")
}
