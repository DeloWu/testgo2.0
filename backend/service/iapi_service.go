package service

import (
    "testgo/models"
)

type IApiService interface {
    //根据ID获取api
    GetApiById(id string) *models.Api
    //分页获取apis
    GetApisByPagination(pageIndex int, pageSize int) *[]models.Api
    //获取总数
    GetApisCounts(maps interface{}) int
    //新增api
    AddApi(api models.Api) error
    //编辑单个api
    EditApi(updatedApi models.Api) error
    //删除单个api
    DeleteApiById(id string) error
}
