package service

import (
    "testgo/models"
)

type IEnvironmentService interface {
    //根据ID获取environment
    GetEnvironmentById(id string) *models.Environment
    //分页获取environments
    GetEnvironmentsByPagination(pageIndex int, pageSize int) *[]models.Environment
    //获取总数
    GetEnvironmentsCounts(maps interface{}) int
    //新增environment
    AddEnvironment(environment models.Environment) error
    //编辑单个environment
    EditEnvironment(updatedEnvironment models.Environment) error
    //删除单个environment
    DeleteEnvironmentById(id string) error
}
