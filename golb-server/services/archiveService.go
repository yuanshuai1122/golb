package services

import "golb/golb-server/models"

type ArchiveService struct{}

func (ArchiveService) GetArchivesList(pageNum int64, pageSize int64) interface{} {
	// 查询归档列表
	result := models.Articles{}.GetArchivesList(pageNum, pageSize)
	if result == nil {
		return nil
	}
	return result
}
