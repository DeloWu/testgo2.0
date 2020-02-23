package service

import (
    "testgo/models"
)

type ISuiteService interface {
    //根据ID获取suite
    GetSuiteById(id string) *models.Suite
    //分页获取suites
    GetSuitesByPagination(pageIndex int, pageSize int) *[]models.Suite
    //获取总数
    GetSuitesCounts(maps interface{}) int
    //新增suite
    AddSuite(suite models.Suite) error
    //编辑单个suite
    EditSuite(updatedSuite models.Suite) error
    //删除单个suite
    DeleteSuiteById(id string) error
}
