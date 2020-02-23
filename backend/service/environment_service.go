package service

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/models"
    "testgo/repository"
)

//依赖注入
type EnvironmentService struct {
    Repository repository.IEnvironmentRepository `inject:""`
}

func (s *EnvironmentService) GetEnvironmentById(id string) *models.Environment{
     result, err := s.Repository.GetEnvironmentById(id)
     if err != nil{
         return nil
     }
     return result
}

//获取分页数据
func (s *EnvironmentService) GetEnvironmentsByPagination(pageIndex int, pageSize int) *[]models.Environment{
    maps := bson.M{}
    result, err := s.Repository.GetEnvironmentsByPagination(maps, pageIndex, pageSize)
    if err != nil{
        return nil
    }
    return result
}

//获取总数
func (s *EnvironmentService) GetEnvironmentsCounts(maps interface{}) int{
    counts, err := s.Repository.GetEnvironmentCounts(maps)
    if err != nil{
        return 0
    }
    return counts
}

func (s *EnvironmentService) AddEnvironment(environment models.Environment) error {
    return s.Repository.AddEnvironment(environment)
}

func (s *EnvironmentService) EditEnvironment(updatedEnvironment models.Environment) error{
    return s.Repository.EditEnvironment(updatedEnvironment)
}

func (s *EnvironmentService) DeleteEnvironmentById(id string) error{
    return s.Repository.DeleteEnvironmentById(id)
}