package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"leechan.inline/models/admin"
	"time"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Layout() {

	o := orm.NewOrm()

	// 月份排序
	articleTime := new(admin.Article)
	var articlesTime []*admin.Article
	nqs := o.QueryTable(articleTime)
	nqs = nqs.Filter("status", 1)
	nqs.RelatedSel().All(&articlesTime,"Created")
	var datetime = make(map[string]int64)
	for _,v := range articlesTime  {
		//str = append(str ,v.Created.Format("2006-01"))
		//c.Ctx.WriteString(v.Created.Format("2006-01"))
		k := v.Created.Format("2006-01")
		datetime[k] = datetime[k] + 1
	}
	c.Data["DateTime"] = datetime


	// 阅读排序
	articleReadSort := new(admin.Article)
	var articlesReadSort []*admin.Article
	nqrs := o.QueryTable(articleReadSort)
	nqrs = nqrs.Filter("status", 1)
	nqrs = nqrs.OrderBy("-Pv")
	nqrs.Limit(5).All(&articlesReadSort,"Id","Title","Pv")
	c.Data["ArticlesReadSort"] = articlesReadSort
}

func (c *BaseController) Menu()  {

	o := orm.NewOrm()

	category := new(admin.Category)
	var categorys []*admin.Category
	cqs := o.QueryTable(category)
	cqs = cqs.Filter("status", 1)
	cqs = cqs.Filter("pid", 0)
	cqs.OrderBy("-sort").All(&categorys)
	c.Data["Menu"] = categorys

}
func (c *BaseController) Prepare(){
	c.Data["bgClass"] = "bgColor"
	c.Data["T"] = time.Now()
}

