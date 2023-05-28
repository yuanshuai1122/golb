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

type ArticlesInfo struct {
	// 主键id
	Id int `gorm:"column:id" json:"id"`
	// 文章id
	ArticleId int `gorm:"column:article_id" json:"articleId"`
	// 文章内容
	Content string `gorm:"column:content" json:"content"`
	// 浏览次数
	Views int64 `gorm:"column:views" json:"views"`
}

// TableName 绑定表名
func (Articles) TableName() string {
	return "articles"
}

// Archives 归档
type Archives struct {
	// 文章标题
	Title string `gorm:"column:title" json:"title"`
	// 创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
}

// ArticleInfo 文章详情
type ArticleDetail struct {
	// 文章标题
	Title string `gorm:"column:title" json:"title"`
	// 文章封面图
	CoverImg string `gorm:"column:cover_img" json:"coverImg"`
	// 文章摘要
	Abstract string `gorm:"column:abstract" json:"abstract"`
	// 文章内容
	Content string `gorm:"column:content" json:"content"`
	// 浏览次数
	Views int64 `gorm:"column:views" json:"views"`
	// 创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	// 修改时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
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

// GetArchivesList 获取归档列表
func (Articles) GetArchivesList(pageNum int64, pageSize int64) interface{} {
	offset := (pageNum - 1) * pageSize
	var count int64
	var archivesList []Archives
	// 查总数
	if err := utils.DB.Table("articles").Select([]string{"title", "create_time"}).Where("status = ?", "normal").Count(&count).Error; err != nil {
		return nil
	}
	// 分页查询
	if err := utils.DB.Table("articles").Select([]string{"title", "create_time"}).Where("status = ?", "normal").Order("create_time desc").Limit(int(pageSize)).Offset(int(offset)).Scan(&archivesList).Error; err != nil {
		return nil
	}
	totalPage := utils.ReturnTotalPage(count, pageSize)
	return map[string]interface{}{
		"pageNum":   pageNum,
		"list":      archivesList,
		"pageSize":  pageSize,
		"total":     count,
		"totalPage": totalPage,
	}
}

// GetArticleInfoById 根据id查询文章详情
func (Articles) GetArticleInfoById(id int64) interface{} {
	var articleDetail ArticleDetail
	if err := utils.DB.Table("articles").Select("articles.title, articles.cover_img, articles.abstract, articles.create_time, articles.update_time, articles_info.content, articles_info.views").Where("articles.id = ?", id).Joins("inner join articles_info on articles.id = articles_info.article_id").Scan(&articleDetail).Error; err != nil {
		return nil
	}

	return articleDetail

}
