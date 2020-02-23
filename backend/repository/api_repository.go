package repository

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/common/datasource"
    "testgo/common/logger"
    "testgo/models"
    "time"
)

//依赖注入
type ApiRepository struct {
    Source datasource.IDb `inject:""`
    Log logger.ILogger `inject:""`
    Base BaseRepository `inject:"inline"`
}

//分页获取apis
func (r *ApiRepository) GetApisByPagination(maps interface{}, pageIndex int, pageSize int) (*[]models.Api, error) {
    var result []models.Api
    err := r.Base.FindByPagination("api", maps, &result, pageIndex, pageSize)
    if err != nil {
        r.Base.Log.Error(err)
        return nil, err
    }
    return &result, err
}

//根据id获取单个api
func (r *ApiRepository) GetApiById(id string) (*models.Api, error) {
    var result models.Api
    err := r.Base.FindById("api", id, &result)
    if err != nil{
        r.Base.Log.Error(err)
        return nil, err
    }
    return &result, err
}

//新增单个Api
func (r *ApiRepository) AddApi(api models.Api) error {
    api.ID = bson.NewObjectId()
    api.CreateTime = time.Now().Unix()
    api.UpdateTime = time.Now().Unix()
    err := r.Base.Insert("api", api)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

//编辑单个api
func (r *ApiRepository) EditApi(updatedApi models.Api) error {
    idSelector := bson.M{
        "_id": updatedApi.ID,
    }
    updatedApi.UpdateTime = time.Now().Unix()
    err := r.Base.Update("api", idSelector, updatedApi)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

//删除单个api
func (r *ApiRepository) DeleteApiById(id string) error {
    err := r.Base.DeleteById("api", id)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

//获取总数
func (r *ApiRepository) GetApiCounts(maps interface{}) (int, error) {
    num, err := r.Base.Count("api", maps)
    if err != nil {
        r.Base.Log.Error(err)
        return 0, err
    }
    return num, nil
}