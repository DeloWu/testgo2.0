package repository

import (
    "testgo/models"
)

type IProjectRepository interface {
    //分页返回Projects
    GetProjectsByPagination(maps interface{}, pageIndex int, pageSize int) (*[]models.Project, error)
    //根据ID获取project
    GetProjectById(id string) (*models.Project, error)
    //根据条件获取多个projects
    //GetProjects(maps interface{}) *[]models.Project
    //根据条件获取单个project
    //GetProject(maps interface{}) *models.Project
    //新增project
    AddProject(project models.Project) error
    //编辑project
    EditProject(updatedProject models.Project) error
    //删除project
    DeleteProjectById(id string) error
    //获取所有projects数量
    GetProjectCounts(maps interface{}) (int, error)

}