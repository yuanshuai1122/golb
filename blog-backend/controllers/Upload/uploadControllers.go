package Upload

import (
	"GIN/utils"
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
)

type UploadController struct{}

func (con UploadController) AdvertisingImgUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	fmt.Println(file.Filename)
	dst := path.Join("./static/advertisingImg", file.Filename)

	if err != nil {
		//	出错
		utils.Failed(c, "上传失败!")
	} else {
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)
		utils.SuccessData(c, dst)
	}
}

/**
 * 缩略图上传
 * @author lizibin
 */
func (con UploadController) ThumbnailUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	fmt.Println(file.Filename)
	dst := path.Join("./static/thumbnail", file.Filename)
	//dst := path.Join("/usr/local/nginx/html/upImg//thumbnail" , file.Filename)
	if err != nil {
		//	出错
		utils.Failed(c, "上传失败!")
	} else {
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)
		utils.SuccessData(c, dst)
	}
}

func (con UploadController) ArticleImgUpload(c *gin.Context) {
	file, err := c.FormFile("wangeditor-uploaded-image")
	fmt.Println(file.Filename)
	dst := path.Join("./static/articleImg", file.Filename)

	if err != nil {
		//	出错
		utils.Failed(c, "上传失败!")
	} else {
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)
		utils.BackEditorImg(c, dst)
	}
}

func (con UploadController) ArticleVideoUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	fmt.Println(file.Filename)
	dst := path.Join("./static/articleVideo", file.Filename)

	if err != nil {
		//	出错
		utils.Failed(c, "上传失败!")
	} else {
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)
		utils.SuccessData(c, dst)
	}
}

func (con UploadController) FileUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	fmt.Println(file.Filename)
	dst := path.Join("./static/upload", file.Filename)

	if err != nil {
		//	出错
		utils.Failed(c, "上传失败!")
	} else {
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)
		utils.SuccessData(c, dst)
	}
}

func (con UploadController) UserIcon(c *gin.Context) {
	file, err := c.FormFile("file")
	fmt.Println(file.Filename)
	dst := path.Join("./static/userIcon", file.Filename)

	if err != nil {
		//	出错
		utils.Failed(c, "上传失败!")
	} else {
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)
		utils.SuccessData(c, dst)
	}
}
