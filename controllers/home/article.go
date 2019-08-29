package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"leechan.inline/models/admin"
	"leechan.inline/utils"
	"strconv"
	"time"
)

type ArticleController struct {
	BaseController
}
// 类表
func (c *ArticleController) List() {


	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量
	categoryId, _ := c.GetInt("c", 0)           // 页数

	o := orm.NewOrm()
	article := new(admin.Article)

	var articles []*admin.Article
	qs := o.QueryTable(article)
	qs = qs.Filter("status", 1)
	qs = qs.Filter("User__Name__isnull", false)
	qs = qs.Filter("Category__Name__isnull", false)

	if categoryId != 0 {

		category := new(admin.Category)
		var categorys []*admin.Category
		cqs := o.QueryTable(category)
		cqs = cqs.Filter("status", 1)
		cqs.OrderBy("-sort").All(&categorys)

		ids := utils.CategoryTreeR(categorys,categoryId,0)

		var cids []int
		cids = append(cids,categoryId)
		for _,v := range ids{
			cids = append(cids,v.Id)
		}

		/*c.Data["json"] = cids
		c.ServeJSON()
		c.StopRun()*/

		qs = qs.Filter("Category__ID__in", cids)


	}
	c.Data["CategoryID"] = &categoryId
	// 查出当前分类下的所有子分类id

	date := c.GetString("date")
	if date != "" {
		if len(date) == 7 {
			start := ""
			end := ""
			dateNumStr := beego.Substr(date, len("2018-"),2)
			yearNumStr := beego.Substr(date, len("20"),2)
			dateNum,_:=strconv.Atoi(dateNumStr)
			yearNum,_:=strconv.Atoi(yearNumStr)

			start = utils.SubString(date, len("2018-01"))+"-01 00:00:00"
			if dateNum >= 12 {
				endYearStr := strconv.Itoa(yearNum+1)
				end = utils.SubString(date, len("20"))+endYearStr+"-01-01 00:00:00"
			}

			if dateNum < 9 {
				endStr := strconv.Itoa(dateNum+1)
				end = utils.SubString(date, len("2018-0"))+endStr+"-01 00:00:00"
			}
			if dateNum >= 9 && dateNum < 12 {
				endStr := strconv.Itoa(dateNum+1)
				end = utils.SubString(date, len("2018-"))+endStr+"-01 00:00:00"
			}

			/*c.Data["json"] = []string{start,end}
			c.ServeJSON()
			c.StopRun()*/

			qs = qs.Filter("created__gte", start)
			qs = qs.Filter("created__lte", end)
			c.Data["Date"] = utils.SubString(start, len("2018-01"))

		}else {
			date = utils.SubString(date, len("2018-01-01"))
			tm, _ := time.Parse("2006-01-02", date)
			unix := tm.Unix() //1566432000

			startFormat := time.Unix(unix, 0).Format("2006-01-02 15:04:05")
			moreUnix, _ := utils.ToInt64(int64(60 * 60 * 24))
			endFormat := time.Unix(unix+moreUnix, 0).Format("2006-01-02 15:04:05")
			start := utils.SubString(startFormat, len("2018-01-01")) + " 00:00:00"
			end := utils.SubString(endFormat, len("2018-01-01")) + " 00:00:00"

			// 刷选
			qs = qs.Filter("created__gte", start)
			qs = qs.Filter("created__lte", end)
			c.Data["Date"] = utils.SubString(start, len("2018-01-01"))
		}
	}


	// 统计
	count, err := qs.Count()
	if err != nil {
		c.Abort("404")
	}
	/*c.Data["json"] = count
	c.ServeJSON()
	c.StopRun()*/


	// 获取数据
	_, err = qs.OrderBy("-id","-pv").RelatedSel().Limit(limit).Offset(offset).All(&articles)
	if err != nil {
		panic(err)
	}
	c.Data["Data"] = &articles
	c.Data["Paginator"] = utils.GenPaginator(page, limit, count)


	// Menu
	c.Menu()
	c.Layout()

	c.TplName = "home/list.html"
}

// 详情
func (c *ArticleController) Detail() {

	id := c.Ctx.Input.Param(":id")
	// 基础数据
	o := orm.NewOrm()
	article := new(admin.Article)
	var articles []*admin.Article
	qs := o.QueryTable(article)
	err := qs.Filter("id", id).RelatedSel().One(&articles)
	if err != nil {
		c.Abort("404")
	}

	/*c.Data["json"]= &articles
	c.ServeJSON()
	c.StopRun()*/

	c.Data["Data"] = &articles[0]
	
	c.Menu()
	c.Layout()
	c.TplName = "home/detail.html"
}

// 统计访问量
func (c *ArticleController) Pv()  {

	ids := c.Ctx.Input.Param(":id")
	id,_:=strconv.Atoi(ids)
	/*c.Data["json"] = c.Input()
	c.ServeJSON()
	c.StopRun()*/

	response := make(map[string]interface{})

	o := orm.NewOrm()

	article := admin.Article{Id: id}
	if o.Read(&article) == nil {
		article.Pv = article.Pv + 1

		valid := validation.Validation{}
		valid.Required(article.Id, "Id")
		if valid.HasErrors() {
			// 如果有错误信息，证明验证没通过
			// 打印错误信息
			for _, err := range valid.Errors {
				//log.Println(err.Key, err.Message)
				response["msg"] = "Error."
				response["code"] = 500
				response["err"] = err.Key + " " + err.Message
				c.Data["json"] = response
				c.ServeJSON()
				c.StopRun()
			}
		}

		if _, err := o.Update(&article); err == nil {
			response["msg"] = "Success."
			response["code"] = 200
			response["id"] = id
		} else {
			response["msg"] = "Error."
			response["code"] = 500
			response["err"] = err.Error()
		}
	} else {
		response["msg"] = "Error."
		response["code"] = 500
		response["err"] = "ID 不能为空！"
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

// 评论
func (c *ArticleController) Review() {

	aid, _ := c.GetInt("aid")
	name := c.GetString("name")
	review := c.GetString("review")
	site := c.GetString("site","")

	o := orm.NewOrm()
	reviewsMd := admin.Review{
		Name:    	name,
		Review:     review,
		Site:    	site,
		ArticleId:	aid,
		Status:   	1,
	}

	response := make(map[string]interface{})

	valid := validation.Validation{}
	valid.Required(reviewsMd.Name, "Name")
	valid.Required(reviewsMd.Review, "Review")
	valid.Required(reviewsMd.ArticleId, "ArticleId")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			//log.Println(err.Key, err.Message)
			response["msg"] = "新增失败！"
			response["code"] = 500
			response["err"] = err.Key + " " + err.Message
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}
	}

	// 更新评论数量
	article:= admin.Article{Id: aid}
	o.Read(&article)
	article.Review = article.Review + 1
	o.Update(&article)


	if id, err := o.Insert(&reviewsMd); err == nil {
		response["msg"] = "新增成功！"
		response["code"] = 200
		response["id"] = id
	} else {
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = err.Error()
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

func (c *ArticleController) ReviewList()  {

	id := c.Ctx.Input.Param(":id")
	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量
	response := make(map[string]interface{})

	o := orm.NewOrm()
	review := new(admin.Review)

	var reviews []*admin.Review
	qs := o.QueryTable(review)
	qs = qs.Filter("status", 1)
	qs = qs.Filter("article_id", id)

	// 获取数据
	_, err := qs.OrderBy("-id").Limit(limit).Offset(offset).All(&reviews)
	if err != nil {
		response["msg"] = "Error."
		response["code"] = 500
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	// 统计
	count, err := qs.Count()
	if err != nil {
		response["msg"] = "Error."
		response["code"] = 500
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	response["Data"] = &reviews
	response["Paginator"] = utils.GenPaginator(page, limit, count)

	response["msg"] = "Success."
	response["code"] = 200
	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()

}
