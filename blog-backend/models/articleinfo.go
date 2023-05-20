package models

import (
	"GIN/utils"
	"time"
)
// 返回列表使用
type  ArticleInfo struct {
	EditArticleInfo
	UserId int			`gorm:"column:userId" json:"userId"`
	UserName string 	`gorm:"column:userName" json:"userName"`
	PublishTime time.Time `gorm:"column:publishTime" json:"publishTime"`
	Click int `json:"click"`
	Review int `json:"review"`
}

func (ArticleInfo) TableName() string  {
	return "articleinfo"
}

type  AddArticleInfo struct {
	EditArticleInfo
	UserId int			`gorm:"column:userId" json:"userId"`
	UserName string 	`gorm:"column:userName" json:"userName"`
}
// add时接受用
func (AddArticleInfo) TableName() string  {
	return "articleinfo"
}

type  EditArticleInfo struct {
	ArticleId int 		`gorm:"column:articleId" json:"articleId"`
	ArticleTitle string `gorm:"column:articleTitle" json:"articleTitle"`
	ArticleClassifyId int `gorm:"column:articleClassifyId" json:"articleClassifyId"`
	ArticleClassifyName string `gorm:"column:articleClassifyName" json:"articleClassifyName"`
	ArticleDase string `gorm:"column:articleDase" json:"articleDase"`
	ArticleImgLitimg string `gorm:"column:articleImgLitimg" json:"articleImgLitimg"`
	ArticleContent string `gorm:"column:articleContent" json:"articleContent"`
	ArticleState int `gorm:"column:articleState" json:"articleState"`
	ArticlePass int `gorm:"column:articlePass" json:"articlePass"`
	CommentState int `gorm:"column:commentState" json:"commentState"`
}
// add时接受用
func (EditArticleInfo) TableName() string  {
	return "articleinfo"
}

// 根据用户ID删除文章
func (ArticleInfo) DeleteUserArticByUserID(userID string) interface{}  {
	if err :=  utils.DB.Where("userID = ?",userID).Delete(&ArticleInfo{}).Error; err != nil {
		return 1
	}else {
		return 0
	}
}


// 查询审核通过的文章 分页
func (ArticleInfo) ShowAllPass(currentPage int64,pageSize int64, keyword string) interface{}  {
	offset := (currentPage - 1) * pageSize
	var count int64
	articleInfo := []ArticleInfo{}
	// 查总数
	if err := utils.DB.Where("(articleTitle LIKE ? or articleDase LIKE ? or userName LIKE ? or articleClassifyName LIKE ?) and articlePass = ?",keyword+"%",keyword+"%",keyword+"%",keyword+"%",2).Find(&articleInfo).Count(&count).Error; err != nil {
		return 1
	}
	if err :=  utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("(articleTitle LIKE ? or articleDase LIKE ? or userName LIKE ? or articleClassifyName LIKE ?) and articlePass = ?",keyword+"%",keyword+"%",keyword+"%",keyword+"%",2).Find(&articleInfo).Error; err != nil {
		return 1
	}else {
		totalPage := utils.ReturnTotalPage(count , pageSize)
		return map[string]interface{}{
			"currentPage": currentPage,
			"list": articleInfo,
			"pageSize": pageSize,
			"total": count,
			"totalPage": totalPage,
		}
	}
}


// 根据文章id返回文章名
func (ArticleInfo) SearchNameByArticleID(articleID string) interface{}  {
	articleInfo := ArticleInfo{}
	if err :=  utils.DB.Where("articleID = ?",articleID).First(&articleInfo).Error; err != nil {
		return 1
	}else {
		return articleInfo.ArticleTitle
	}
}

// ShowArticleCount 查询审核通过的文章数量
func (ArticleInfo) ShowArticleCount() interface{}  {
	var count int64
	if err := utils.DB.Where("articlePass = 2").Find(&ArticleInfo{}).Count(&count).Error; err != nil {
		return 1
	}else {
		return count
	}
}

//ShowAll 查询审核通过的文章
func (ArticleInfo) ShowAll() interface{}  {
	articleInfo := []ArticleInfo{}
	if err := utils.DB.Where("articlePass = 2").Find(&articleInfo).Error; err != nil {
		return 1
	}else {
		return articleInfo
	}
}

// 删除文章
func (ArticleInfo) Delete(articleId string) int  {
	if err :=  utils.DB.Where("articleId = ?",articleId).Delete(&ArticleInfo{}).Error; err != nil {
		return 1
	}else {
		return 0
	}
}