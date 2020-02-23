package repository

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/common/datasource"
    "testgo/common/logger"
    "testgo/models"
    "time"
)

//依赖注入
type CaseRepository struct {
    Source datasource.IDb `inject:""`
    Log logger.ILogger `inject:""`
    Base BaseRepository `inject:"inline"`
}

//分页获取cases
func (r *CaseRepository) GetCasesByPagination(maps interface{}, pageIndex int, pageSize int) (*[]models.Case, error) {
    var result []models.Case
    err := r.Base.FindByPagination("case", maps, &result, pageIndex, pageSize)
    if err != nil {
        r.Base.Log.Error(err)
        return nil, err
    }
    return &result, err
}

//根据id获取单个case
func (r *CaseRepository) GetCaseById(id string) (*models.Case, error) {
    var result models.Case
    err := r.Base.FindById("case", id, &result)
    if err != nil{
        r.Base.Log.Error(err)
        return nil, err
    }
    return &result, err
}

//新增单个Case
func (r *CaseRepository) AddCase(c models.Case) error {
   c.ID = bson.NewObjectId()
   c.CreateTime = time.Now().Unix()
   c.UpdateTime = time.Now().Unix()
   err := r.Base.Insert("case", c)
   if err != nil {
       r.Base.Log.Error(err)
       return err
   }
   return nil
}

//编辑单个case
func (r *CaseRepository) EditCase(updatedCase models.Case) error {
    idSelector := bson.M{
        "_id": updatedCase.ID,
    }
    updatedCase.UpdateTime = time.Now().Unix()
    err := r.Base.Update("case", idSelector, updatedCase)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

//删除单个case
func (r *CaseRepository) DeleteCaseById(id string) error {
    err := r.Base.DeleteById("case", id)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

//获取总数
func (r *CaseRepository) GetCaseCounts(maps interface{}) (int, error) {
    num, err := r.Base.Count("case", maps)
    if err != nil {
        r.Base.Log.Error(err)
        return 0, err
    }
    return num, nil
}