package service

import (
    "testgo/models"
)

type ICaseService interface {
    //根据ID获取case
    GetCaseById(id string) *models.Case
    //分页获取cases
    GetCasesByPagination(pageIndex int, pageSize int) *[]models.Case
    //获取总数
    GetCasesCounts(maps interface{}) int
    //新增case
    AddCase(c models.Case) error
    //编辑单个case
    EditCase(updatedCase models.Case) error
    //删除单个case
    DeleteCaseById(id string) error
}
