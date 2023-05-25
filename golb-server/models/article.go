package models

import (
	"golb/golb-server/utils"
	"time"
)

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
	// 创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	// 修改时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

// TableName 绑定表名
func (Articles) TableName() string {
	return "articles"
}

// GetArticlesList 获取文章列表（分页）
func (Articles) GetArticlesList(pageNum int64, pageSize int64, keywords string) interface{} {
	offset := (pageNum - 1) * pageSize
	var count int64
	var articleList []Articles
	// 查总数
	if err := utils.DB.Where("(title LIKE ? or abstract LIKE ?) and status = ?", keywords+"%", keywords+"%", "normal").Find(&articleList).Count(&count).Error; err != nil {
		return nil
	}
	// 分页查询
	if err := utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("(title LIKE ? or abstract LIKE ?) and status = ?", keywords+"%", keywords+"%", "normal").Find(&articleList).Error; err != nil {
		return nil
	} else {
		totalPage := utils.ReturnTotalPage(count, pageSize)
		return map[string]interface{}{
			"pageNum":   pageNum,
			"list":      articleList,
			"pageSize":  pageSize,
			"total":     count,
			"totalPage": totalPage,
		}
	}

}
