package routers

import (
	"github.com/astaxie/beego"
	"leechan.inline/controllers/admin"
	"leechan.inline/controllers/home"
)

func init() {

	// 后台文章模块
	beego.Router("/", &home.MainController{})
	beego.Router("/admin", &admin.MainController{}, "get:Index")
	beego.Router("/admin/welcome", &admin.MainController{}, "get:Welcome")
	beego.Router("/admin/article", &admin.ArticleController{}, "get:List;post:Save")
	beego.Router("/admin/article/edit", &admin.ArticleController{}, "get:Put")
	beego.Router("/admin/article/delete", &admin.ArticleController{}, "Post:Delete")
	beego.Router("/admin/article/update", &admin.ArticleController{}, "Post:Update")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "get:Add")
	// 后台分类模块
	beego.Router("/admin/cate", &admin.CateController{}, "get:List;post:Save")
	beego.Router("/admin/cate/add", &admin.CateController{}, "get:Add")
	beego.Router("/admin/cate/edit", &admin.CateController{}, "get:Put")
	beego.Router("/admin/cate/delete", &admin.CateController{}, "Post:Delete")
	beego.Router("/admin/cate/update", &admin.CateController{}, "Post:Update")
	// 后台登录
	beego.Router("/admin/login", &admin.LoginController{}, "Get:Sign;Post:Login")

	// 后台评论
	beego.Router("/admin/review", &admin.ReviewController{}, "get:List")
	beego.Router("/admin/review/edit", &admin.ReviewController{}, "Get:Put")
	beego.Router("/admin/review/delete", &admin.ReviewController{}, "Post:Delete")
	beego.Router("/admin/review/update", &admin.ReviewController{}, "Post:Update")


	// 后台友情链接
	beego.Router("/admin/link", &admin.LinkController{}, "get:List")
	// 后台留言
	beego.Router("/admin/message", &admin.MessageController{}, "get:List")

	// 后台留言模块
	beego.Router("/admin/message/update", &admin.MessageController{}, "Post:Update")
	beego.Router("/admin/message/edit", &admin.MessageController{}, "Get:Put")
	beego.Router("/admin/message/delete", &admin.MessageController{}, "Post:Delete")


	// 前台列表
	beego.Router("/list.html", &home.ArticleController{}, "get:List")
	// 前台详情
	beego.Router("/detail/:id([0-9]+).html", &home.ArticleController{}, "get:Detail")
	// 前台统计文章PV
	beego.Router("/pv/:id([0-9]+).html", &home.ArticleController{}, "get:Pv")
	// 前台留言列表
	beego.Router("/message.html", &home.MessageController{}, "get:Get")
	// 前台保存
	beego.Router("/message/save", &home.MessageController{}, "Post:Save")

	// 评论保存
	beego.Router("article/review", &home.ArticleController{}, "Post:Review")
	beego.Router("article/review/:id([0-9]+).html", &home.ArticleController{}, "Get:ReviewList")



}
