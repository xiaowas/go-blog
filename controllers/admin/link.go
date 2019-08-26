package admin

type LinkController struct {
	BaseController
}

func (c *LinkController) List() {
	c.TplName = "admin/link-list.html"
}
