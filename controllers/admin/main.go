package admin

type MainController struct {
	BaseController
}

func (c *MainController) Index() {
	c.TplName = "admin/index.html"
}

func (c *MainController) Welcome() {
	c.TplName = "admin/welcome.html"
}
