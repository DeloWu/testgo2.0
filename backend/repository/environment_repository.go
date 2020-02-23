package repository

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/common/datasource"
    "testgo/common/logger"
    "testgo/models"
    "time"
)

//依赖注入
type EnvironmentRepository struct {
    Source datasource.IDb `inject:""`
    Log logger.ILogger `inject:""`
    Base BaseRepository `inject:"inline"`
}

func (r *EnvironmentRepository) GetEnvironmentsByPagination(maps interface{}, pageIndex int, pageSize int) (*[]models.Environment, error) {
    var result []models.Environment
    err := r.Base.FindByPagination("environment", maps, &result, pageIndex, pageSize)
    if err != nil {
        r.Base.Log.Error(err)
        return nil, err
    }
    return &result, err
}

func (r *EnvironmentRepository) GetEnvironmentById(id string) (*models.Environment, error) {
    var result models.Environment
    err := r.Base.FindById("environment", id, &result)
    if err != nil{
        r.Base.Log.Error(err)
        return nil, err
    }
    return &result, err
}

func (r *EnvironmentRepository) AddEnvironment(environment models.Environment) error {
    environment.ID = bson.NewObjectId()
    environment.CreateTime = time.Now().Unix()
    environment.UpdateTime = time.Now().Unix()
    err := r.Base.Insert("environment", environment)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

func (r *EnvironmentRepository) EditEnvironment(updatedEnvironment models.Environment) error {
    idSelector := bson.M{
        "_id": updatedEnvironment.ID,
    }
    updatedEnvironment.UpdateTime = time.Now().Unix()
    err := r.Base.Update("environment", idSelector, updatedEnvironment)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

func (r *EnvironmentRepository) DeleteEnvironmentById(id string) error {
    err := r.Base.DeleteById("environment", id)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

func (r *EnvironmentRepository) GetEnvironmentCounts(maps interface{}) (int, error) {
    num, err := r.Base.Count("environment", maps)
    if err != nil {
        r.Base.Log.Error(err)
        return 0, err
    }
    return num, nil
}