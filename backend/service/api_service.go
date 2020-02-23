package service

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/models"
    "testgo/repository"
)

//依赖注入
type ApiService struct {
    Repository repository.IApiRepository `inject:""`
}

//根据ID获取api
func (s *ApiService) GetApiById(id string) *models.Api{
     result, err := s.Repository.GetApiById(id)
     if err != nil{
         return nil
     }
     return result
}

//获取分页数据
func (s *ApiService) GetApisByPagination(pageIndex int, pageSize int) *[]models.Api{
    maps := bson.M{}
    result, err := s.Repository.GetApisByPagination(maps, pageIndex, pageSize)
    if err != nil{
        return nil
    }
    return result
}

//获取总数
func (s *ApiService) GetApisCounts(maps interface{}) int{
    counts, err := s.Repository.GetApiCounts(maps)
    if err != nil{
        return 0
    }
    return counts
}

//新增api
func (s *ApiService) AddApi(api models.Api) error {
    return s.Repository.AddApi(api)
}

//编辑单个api
func (s *ApiService) EditApi(updatedApi models.Api) error{
    return s.Repository.EditApi(updatedApi)
}

//删除单个api
func (s *ApiService) DeleteApiById(id string) error{
    return s.Repository.DeleteApiById(id)
}