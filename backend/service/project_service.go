package service

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/models"
    "testgo/repository"
)

//依赖注入
type ProjectService struct {
    Repository repository.IProjectRepository `inject:""`
}

//根据ID获取project
func (s *ProjectService) GetProjectById(id string) *models.Project{
     result, err := s.Repository.GetProjectById(id)
     if err != nil{
         return nil
     }
     return result
}

//获取分页数据
func (s *ProjectService) GetProjectsByPagination(pageIndex int, pageSize int) *[]models.Project{
    maps := bson.M{}
    result, err := s.Repository.GetProjectsByPagination(maps, pageIndex, pageSize)
    if err != nil{
        return nil
    }
    return result
}

//获取总数
func (s *ProjectService) GetProjectsCounts(maps interface{}) int{
    counts, err := s.Repository.GetProjectCounts(maps)
    if err != nil{
        return 0
    }
    return counts
}

//新增project
func (s *ProjectService) AddProject(project models.Project) error {
    return s.Repository.AddProject(project)
}

//编辑单个project
func (s *ProjectService) EditProject(updatedProject models.Project) error{
    return s.Repository.EditProject(updatedProject)
}

//删除单个project
func (s *ProjectService) DeleteProjectById(id string) error{
    return s.Repository.DeleteProjectById(id)
}