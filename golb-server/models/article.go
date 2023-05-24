package models

import "golb/golb-server/utils"

// Articles 文章实体
type Articles struct {
	// 主键id
	Id int `gorm:"column:id" json:"id"`
	// 用户id
	UserId int `gorm:"column:user_id" json:"userId"`
	// 文章标题
	Title string `gorm:"column:title" json:"title"`
	// 文章封面图
	CoverImg string `gorm:"column:cover_img" json:"coverImg"`
	// 文章摘要
	Abstract string `gorm:"column:abstract" json:"abstract"`
	// 状态 normal：正常 deleted：已删除
	Status string `gorm:"column:status" json:"status"`
}

// TableName 绑定表名
func (Articles) TableName() string {
	return "articles"
}

// GetArticlesList 获取文章列表（不分页）
func (Articles) GetArticlesList() interface{} {
	var articleList []Articles
	if err := utils.DB.Where("status = 'normal'").Find(&articleList).Error; err != nil {
		return nil
	} else {
		return articleList
	}
}
