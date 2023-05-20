package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS          int = 0
	FAILED           int = 1001 //操作失败
	MISSSID          int = 1002 //ID为空
	NOID             int = 1003 //ID在数据库不存在
	DELETEFAILED     int = 1004 //执行数据库删除失败
	ADDFAILED        int = 1005 //执行数据库删除失败
	SEARCHPAGEFAILED int = 1006 //执行数据库删除失败
)

// 返回成功不带数据
func Success(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "OK",
	})
}

// 返回成功带数据
func SuccessData(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": SUCCESS,
		"msg":  "OK",
		"data": v,
	})
}

// 返回失败 + 提示信息
func Failed(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": FAILED,
		"msg":  v,
	})
}

// 返回失败 + 提示信息 + 状态码
func FailedMsgAndCode(c *gin.Context, v interface{}, x interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": x,
		"msg":  v,
	})
}

// 执行数据库删除失败
func DeleteFailed(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": DELETEFAILED,
		"msg":  "执行数据库删除失败",
	})
}

// 执行数据库新增失败
func AddFailed(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": ADDFAILED,
		"msg":  "执行数据库新增失败",
	})
}

// ID为空
func MissId(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": MISSSID,
		"msg":  "ID为空!",
	})
}

// ID在数据库不存在
func NoId(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": NOID,
		"msg":  "ID在数据库不存在!",
	})
}

// 分页查询失败
func SearchPageFaild(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code": SEARCHPAGEFAILED,
		"msg":  "分页查询失败!",
	})
}

// Editer 上传图片专用返回格式
func BackEditorImg(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"error": 0,
		"data": map[string]interface{}{
			"url":  v,
			"alt":  "",
			"href": "",
		},
	})
}
