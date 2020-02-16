package repository

import (
    "log"
    //"gopkg.in/mgo.v2/bson"
    "testgo/common/datasource"
    "testgo/models"
    "time"
)

//依赖注入
type ProjectRepository struct {
    Source datasource.IDb `inject:""`
    Base BaseRepository `inject:"inline"`
}

//分页获取projects
func (p *ProjectRepository) GetProjectsByPagination(maps map[string]interface{}, pageIndex int, pageSize int) *[]models.Project{
    var projects []models.Project
    err := p.Base.FindByPagination("project", maps, projects, pageIndex, pageSize)
    if err != nil{
        p.Base.Log.Error(err)
    }
    //debug
    log.Println("**************************************")
    log.Println(projects)
    log.Println("**************************************")
    return &projects
}

//根据id获取单个project
func (p *ProjectRepository) GetProjectById(id string) models.Project{
    var project models.Project
    err := p.Base.FindById("project", id, project)
    if err != nil{
        p.Base.Log.Error(err)
    }
    return project
}

//新增单个Project
func (p *ProjectRepository) AddProject(project models.Project) bool {
    project.CreateTime = time.Now().Unix()
    project.UpdateTime = time.Now().Unix()
    err := p.Base.Insert("project", project)
    if err != nil {
        p.Base.Log.Error(err)
        return false
    }
    return true
}

//编辑单个project
func (p *ProjectRepository) EditProject(project models.Project) bool {
    idSelector := map[string]interface{}{
        "_id": project.ID,
    }
    err := p.Base.Update("project", idSelector, project)
    if err != nil {
        p.Base.Log.Error(err)
        return false
    }
    return true
}

//删除单个project
func (p *ProjectRepository) DeleteProjectById(id interface{}) bool {
    err := p.Base.DeleteById("project", id)
    if err != nil {
        p.Base.Log.Error(err)
        return false
    }
    return true
}

//获取总数
func (p *ProjectRepository) GetProjectCount(maps map[string]interface{}) int {
    num, err := p.Base.Count("project", maps)
    if err != nil {
        p.Base.Log.Error(err)
        return 0
    }
    return num
}