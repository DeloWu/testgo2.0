package service

import (
    "testgo/models"
    "testgo/repository"
)

//依赖注入
type ProjectService struct {
    Repository repository.IProjectRepository `inject:""`
}

//根据ID获取project
func (p *ProjectService) GetProjectById(id string) models.Project{
    return p.Repository.GetProjectById(id)
}

//获取分页数据
func (p *ProjectService) GetProjectsByPagination(pageIndex int, pageSize int) *[]models.Project{
    maps := map[string]interface{}{}
    return p.Repository.GetProjectsByPagination(maps, pageIndex, pageSize)
}

//获取总数
func (p *ProjectService) GetProjectsCount(maps map[string]interface{}) int{
    return p.Repository.GetProjectCount(maps)
}

//新增project
func (p *ProjectService) AddProject(project models.Project) bool {
    return p.Repository.AddProject(project)
}

//编辑单个project
func (p *ProjectService) EditProject(project models.Project) bool{
    return p.Repository.EditProject(project)
}

//删除单个project
func (p *ProjectService) DeleteProjectById(id interface{}) bool{
    return p.DeleteProjectById(id)
}