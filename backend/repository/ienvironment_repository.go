package repository

import (
    "testgo/models"
)

type IEnvironmentRepository interface {
    //分页返回Environments
    GetEnvironmentsByPagination(maps interface{}, pageIndex int, pageSize int) (*[]models.Environment, error)
    //根据ID获取environment
    GetEnvironmentById(id string) (*models.Environment, error)
    //根据条件获取多个environments
    //GetEnvironments(maps interface{}) *[]models.Environment
    //根据条件获取单个environment
    //GetEnvironment(maps interface{}) *models.Environment
    //新增environment
    AddEnvironment(environment models.Environment) error
    //编辑environment
    EditEnvironment(updatedEnvironment models.Environment) error
    //删除environment
    DeleteEnvironmentById(id string) error
    //获取所有environments数量
    GetEnvironmentCounts(maps interface{}) (int, error)

}