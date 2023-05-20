package routers

import (
	"GIN/controllers/Upload"
	"github.com/gin-gonic/gin"
)

func UploadRoutersInit(r *gin.Engine)  {
	uploadRouters:= r.Group("/upload")
	{
		// 新建文章时缩略图上传
		uploadRouters.POST("/thumbnailUpload",Upload.UploadController{}.ThumbnailUpload)
		// 新建文章时缩略图上传
		uploadRouters.POST("/advertisingImgUpload",Upload.UploadController{}.AdvertisingImgUpload)
		//
		uploadRouters.POST("/articleImgUpload",Upload.UploadController{}.ArticleImgUpload)
		//
		uploadRouters.POST("/articleVideoUpload",Upload.UploadController{}.ArticleVideoUpload)
		//
		uploadRouters.POST("/fileUpload",Upload.UploadController{}.FileUpload)
		//
		uploadRouters.POST("/userIcon",Upload.UploadController{}.UserIcon)
	}
}
