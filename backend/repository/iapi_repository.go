package repository

import (
    "testgo/models"
)

type IApiRepository interface {
    //分页返回Apis
    GetApisByPagination(maps interface{}, pageIndex int, pageSize int) (*[]models.Api, error)
    //根据ID获取api
    GetApiById(id string) (*models.Api, error)
    //根据条件获取多个apis
    //GetApis(maps interface{}) *[]models.Api
    //根据条件获取单个api
    //GetApi(maps interface{}) *models.Api
    //新增api
    AddApi(api models.Api) error
    //编辑api
    EditApi(updatedApi models.Api) error
    //删除api
    DeleteApiById(id string) error
    //获取所有apis数量
    GetApiCounts(maps interface{}) (int, error)

}