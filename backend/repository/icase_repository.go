package repository

import (
    "testgo/models"
)

type ICaseRepository interface {
    //分页返回Cases
    GetCasesByPagination(maps interface{}, pageIndex int, pageSize int) (*[]models.Case, error)
    //根据ID获取case
    GetCaseById(id string) (*models.Case, error)
    //根据条件获取多个cases
    //GetCases(maps interface{}) *[]models.Case
    //根据条件获取单个case
    //GetCase(maps interface{}) *models.Case
    //新增case
    AddCase(c models.Case) error
    //编辑case
    EditCase(updatedCase models.Case) error
    //删除case
    DeleteCaseById(id string) error
    //获取所有cases数量
    GetCaseCounts(maps interface{}) (int, error)

}