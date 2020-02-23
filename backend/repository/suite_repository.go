package repository

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/common/datasource"
    "testgo/common/logger"
    "testgo/models"
    "time"
)

//依赖注入
type SuiteRepository struct {
    Source datasource.IDb `inject:""`
    Log logger.ILogger `inject:""`
    Base BaseRepository `inject:"inline"`
}

//分页获取suites
func (r *SuiteRepository) GetSuitesByPagination(maps interface{}, pageIndex int, pageSize int) (*[]models.Suite, error) {
    var result []models.Suite
    err := r.Base.FindByPagination("suite", maps, &result, pageIndex, pageSize)
    if err != nil {
        r.Base.Log.Error(err)
        return nil, err
    }
    return &result, err
}

//根据id获取单个suite
func (r *SuiteRepository) GetSuiteById(id string) (*models.Suite, error) {
    var result models.Suite
    err := r.Base.FindById("suite", id, &result)
    if err != nil{
        r.Base.Log.Error(err)
        return nil, err
    }
    return &result, err
}

//新增单个Suite
func (r *SuiteRepository) AddSuite(suite models.Suite) error {
    suite.ID = bson.NewObjectId()
    suite.CreateTime = time.Now().Unix()
    suite.UpdateTime = time.Now().Unix()
    err := r.Base.Insert("suite", suite)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

//编辑单个suite
func (r *SuiteRepository) EditSuite(updatedSuite models.Suite) error {
    idSelector := bson.M{
        "_id": updatedSuite.ID,
    }
    updatedSuite.UpdateTime = time.Now().Unix()
    err := r.Base.Update("suite", idSelector, updatedSuite)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

//删除单个suite
func (r *SuiteRepository) DeleteSuiteById(id string) error {
    err := r.Base.DeleteById("suite", id)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

//获取总数
func (r *SuiteRepository) GetSuiteCounts(maps interface{}) (int, error) {
    num, err := r.Base.Count("suite", maps)
    if err != nil {
        r.Base.Log.Error(err)
        return 0, err
    }
    return num, nil
}