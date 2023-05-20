package routers

import (
	"GIN/controllers/Article"
	"github.com/gin-gonic/gin"
)

func ArticleRoutersInit(r *gin.Engine)  {
	articleRouters:= r.Group("/article")
	{
		// 通过审核的数量
		articleRouters.GET("/showArticleCount",Article.ArticleController{}.ShowArticleCount)
		// 查询所有审核通过文章，不分页
		articleRouters.GET("/showAllArticleInfo",Article.ArticleController{}.ShowAll)
		// 下拉框筛选文章审核状态，查询出相关内容并进行分页
		articleRouters.GET("/search/pass",Article.ArticleController{}.SearchPass)
		// 查询所有通过审核并且所有可见的文章  筛选条件  分类id
		articleRouters.GET("/byClassifyId",Article.ArticleController{}.SearchByClassifyId)
		// 根据文章标题进行搜索 并且 进行分页  判断是管理员还是普通用户
		articleRouters.GET("/search/title",Article.ArticleController{}.SearchTitle)
		// 后台文章管理：分页查询+通过用户权限获取文章信息   管理员获取全部  普通用户获取自己的
		articleRouters.GET("/page/byUserType",Article.ArticleController{}.SearchbyUserType)
		// 增加文章
		articleRouters.POST("/insertArticleInfo",Article.ArticleController{}.Add)
		// 修改文章
		articleRouters.POST("/updateArticleInfo",Article.ArticleController{}.Edit)
		// 根据ID删除文章
		articleRouters.DELETE("/delete",Article.ArticleController{}.Delete)
		// 根据ID查询用户审核通过的文章
		articleRouters.GET("/showArticleCountByUserId",Article.ArticleController{}.ShowArticleCountByUserId)
		// 根据文章id获取文章信息
		articleRouters.GET("/showArticleInfo",Article.ArticleController{}.ShowArticleInfo)
		// 修改文章点击量，每点击一次+1
		articleRouters.GET("/updateArticleClick",Article.ArticleController{}.UpdateArticleClick)
		// 根据文章id修改文章审核状态
		articleRouters.GET("/updateArticlePass",Article.ArticleController{}.UpdateArticlePass)
		// 获取审核通过的文章 分页  带查询参数
		articleRouters.GET("/page/showAll",Article.ArticleController{}.ShowAllPass)
	}
}
