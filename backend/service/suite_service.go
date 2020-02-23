package service

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/models"
    "testgo/repository"
)

//依赖注入
type SuiteService struct {
    Repository repository.ISuiteRepository `inject:""`
}

//根据ID获取suite
func (s *SuiteService) GetSuiteById(id string) *models.Suite{
     result, err := s.Repository.GetSuiteById(id)
     if err != nil{
         return nil
     }
     return result
}

//获取分页数据
func (s *SuiteService) GetSuitesByPagination(pageIndex int, pageSize int) *[]models.Suite{
    maps := bson.M{}
    result, err := s.Repository.GetSuitesByPagination(maps, pageIndex, pageSize)
    if err != nil{
        return nil
    }
    return result
}

//获取总数
func (s *SuiteService) GetSuitesCounts(maps interface{}) int{
    counts, err := s.Repository.GetSuiteCounts(maps)
    if err != nil{
        return 0
    }
    return counts
}

//新增suite
func (s *SuiteService) AddSuite(suite models.Suite) error {
    return s.Repository.AddSuite(suite)
}

//编辑单个suite
func (s *SuiteService) EditSuite(updatedSuite models.Suite) error{
    return s.Repository.EditSuite(updatedSuite)
}

//删除单个suite
func (s *SuiteService) DeleteSuiteById(id string) error{
    return s.Repository.DeleteSuiteById(id)
}