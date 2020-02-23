package repository

import (
    "testgo/models"
)

type ISuiteRepository interface {
    //分页返回Suites
    GetSuitesByPagination(maps interface{}, pageIndex int, pageSize int) (*[]models.Suite, error)
    //根据ID获取suite
    GetSuiteById(id string) (*models.Suite, error)
    //根据条件获取多个suites
    //GetSuites(maps interface{}) *[]models.Suite
    //根据条件获取单个suite
    //GetSuite(maps interface{}) *models.Suite
    //新增suite
    AddSuite(suite models.Suite) error
    //编辑suite
    EditSuite(updatedSuite models.Suite) error
    //删除suite
    DeleteSuiteById(id string) error
    //获取所有suites数量
    GetSuiteCounts(maps interface{}) (int, error)

}