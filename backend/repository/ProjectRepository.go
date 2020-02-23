package repository

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/common/datasource"
    "testgo/common/logger"
    "testgo/models"
    "time"
)

//依赖注入
type ProjectRepository struct {
    Source datasource.IDb `inject:""`
    Log logger.ILogger `inject:""`
    Base BaseRepository `inject:"inline"`
}

//分页获取projects
func (p *ProjectRepository) GetProjectsByPagination(maps interface{}, pageIndex int, pageSize int) (*[]models.Project, error) {
    var result []models.Project
    err := p.Base.FindByPagination("project", maps, &result, pageIndex, pageSize)
    if err != nil {
        p.Base.Log.Error(err)
        return nil, err
    }
    return &result, err
}

//根据id获取单个project
func (p *ProjectRepository) GetProjectById(id string) (*models.Project, error) {
    var result models.Project
    err := p.Base.FindById("project", id, &result)
    if err != nil{
        p.Base.Log.Error(err)
        return nil, err
    }
    return &result, err
}

//新增单个Project
func (p *ProjectRepository) AddProject(project models.Project) error {
    project.ID = bson.NewObjectId()
    project.CreateTime = time.Now().Unix()
    project.UpdateTime = time.Now().Unix()
    err := p.Base.Insert("project", project)
    if err != nil {
        p.Base.Log.Error(err)
        return err
    }
    return nil
}

//编辑单个project
func (p *ProjectRepository) EditProject(updatedProject models.Project) error {
    idSelector := bson.M{
        "_id": updatedProject.ID,
    }
    updatedProject.UpdateTime = time.Now().Unix()
    err := p.Base.Update("project", idSelector, updatedProject)
    if err != nil {
        p.Base.Log.Error(err)
        return err
    }
    return nil
}

//删除单个project
func (p *ProjectRepository) DeleteProjectById(id string) error {
    err := p.Base.DeleteById("project", id)
    if err != nil {
        p.Base.Log.Error(err)
        return err
    }
    return nil
}

//获取总数
func (p *ProjectRepository) GetProjectCounts(maps interface{}) (int, error) {
    num, err := p.Base.Count("project", maps)
    if err != nil {
        p.Base.Log.Error(err)
        return 0, err
    }
    return num, nil
}