package routers

import (
	"GIN/controllers/Comment"
	"github.com/gin-gonic/gin"
)

func CommentRouterInit(r *gin.Engine)  {
	commentRouters:= r.Group("/comment")
	{
		// 根据评论id删除分类
		commentRouters.DELETE("/delete",Comment.CommentController{}.Delete)
		// 根据文章id删除分类
		commentRouters.DELETE("/deleteByArticleId",Comment.CommentController{}.DeleteByArticleId)
		// 添加评论
		commentRouters.POST("/insertCommentInfo",Comment.CommentController{}.InsertCommentInfo)
		// 查询所有评论
		commentRouters.GET("/showAllComment",Comment.CommentController{}.ShowAllComment)
		// 根据文章id查询文章对应的评论 太难了，目前只支持两级评论
		commentRouters.GET("/showCommentByArticleId",Comment.CommentController{}.ShowCommentByArticleId)
		// 查询评论总数
		commentRouters.GET("/count",Comment.CommentController{}.GetCount)
		// 后台评论管理：分页查询  + 搜索框内容  管理员获取全部  普通用户获取自己的
		commentRouters.GET("/page/byUserType",Comment.CommentController{}.GetCountByUserType)
	}
}
