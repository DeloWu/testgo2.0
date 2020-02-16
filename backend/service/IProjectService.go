package service

import (
    "testgo/models"
)

type IProjectService interface {
    //根据ID获取project
    GetProjectById(id string) models.Project
    //分页获取projects
    GetProjectsByPagination(pageIndex int, pageSize int) *[]models.Project
    //获取总数
    GetProjectsCount(maps map[string]interface{}) int
    //新增project
    AddProject(project models.Project) bool
    //编辑单个project
    EditProject(project models.Project) bool
    //删除单个project
    DeleteProjectById(id interface{}) bool
}
