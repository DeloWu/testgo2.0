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
func (p *ProjectService) GetProjectById(id string) *models.Project{
     result, err := p.Repository.GetProjectById(id)
     if err != nil{
         return nil
     }
     return result
}

//获取分页数据
func (p *ProjectService) GetProjectsByPagination(pageIndex int, pageSize int) *[]models.Project{
    maps := bson.M{}
    result, err := p.Repository.GetProjectsByPagination(maps, pageIndex, pageSize)
    if err != nil{
        return nil
    }
    return result
}

//获取总数
func (p *ProjectService) GetProjectsCounts(maps interface{}) int{
    counts, err := p.Repository.GetProjectCounts(maps)
    if err != nil{
        return 0
    }
    return counts
}

//新增project
func (p *ProjectService) AddProject(project models.Project) error {
    return p.Repository.AddProject(project)
}

//编辑单个project
func (p *ProjectService) EditProject(updatedProject models.Project) error{
    return p.Repository.EditProject(updatedProject)
}

//删除单个project
func (p *ProjectService) DeleteProjectById(id string) error{
    return p.Repository.DeleteProjectById(id)
}