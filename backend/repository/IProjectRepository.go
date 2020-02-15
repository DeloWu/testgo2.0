package repository

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/models"
)

type IProjectRepository interface {
    //分页返回Projects
    GetProjectsByPagination(PageNum int, PageSize int, maps bson.M) *[]models.Project
    //根据条件获取多个projects
    GetProjects(maps bson.M) *[]models.Project
    //根据条件获取单个project
    GetProject(maps bson.M) *models.Project
    //新增project
    AddProject(project models.Project) bool
    //编辑project
    EditProject(project models.Project) bool
    //删除project
    DeleteProject(id bson.ObjectId) bool
    //获取所有projects数量
    GetProjectCount(maps bson.M) (count int)

}