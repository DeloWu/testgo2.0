package service

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/models"
    "testgo/repository"
)

//依赖注入
type CaseService struct {
    Repository repository.ICaseRepository `inject:""`
}

//根据ID获取case
func (s *CaseService) GetCaseById(id string) *models.Case{
     result, err := s.Repository.GetCaseById(id)
     if err != nil{
         return nil
     }
     return result
}

//获取分页数据
func (s *CaseService) GetCasesByPagination(pageIndex int, pageSize int) *[]models.Case{
    maps := bson.M{}
    result, err := s.Repository.GetCasesByPagination(maps, pageIndex, pageSize)
    if err != nil{
        return nil
    }
    return result
}

//获取总数
func (s *CaseService) GetCasesCounts(maps interface{}) int{
    counts, err := s.Repository.GetCaseCounts(maps)
    if err != nil{
        return 0
    }
    return counts
}

//新增case
func (s *CaseService) AddCase(c models.Case) error {
    return s.Repository.AddCase(c)
}

//编辑单个case
func (s *CaseService) EditCase(updatedCase models.Case) error{
    return s.Repository.EditCase(updatedCase)
}

//删除单个case
func (s *CaseService) DeleteCaseById(id string) error{
    return s.Repository.DeleteCaseById(id)
}