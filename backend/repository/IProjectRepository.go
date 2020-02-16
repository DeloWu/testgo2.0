package repository

import (
    "testgo/models"
)

type IProjectRepository interface {
    //分页返回Projects
    GetProjectsByPagination(maps map[string]interface{}, pageIndex int, pageSize int) *[]models.Project
    //根据ID获取project
    GetProjectById(id string) models.Project
    //根据条件获取多个projects
    //GetProjects(maps map[string]interface{}) *[]models.Project
    //根据条件获取单个project
    //GetProject(maps map[string]interface{}) *models.Project
    //新增project
    AddProject(project models.Project) bool
    //编辑project
    EditProject(project models.Project) bool
    //删除project
    DeleteProjectById(id interface{}) bool
    //获取所有projects数量
    GetProjectCount(maps map[string]interface{}) int

}