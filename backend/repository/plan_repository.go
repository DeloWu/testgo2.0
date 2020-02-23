package repository

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/common/datasource"
    "testgo/common/logger"
    "testgo/models"
    "time"
)

//依赖注入
type PlanRepository struct {
    Source datasource.IDb `inject:""`
    Log logger.ILogger `inject:""`
    Base BaseRepository `inject:"inline"`
}

//分页获取plans
func (r *PlanRepository) GetPlansByPagination(maps interface{}, pageIndex int, pageSize int) (*[]models.Plan, error) {
    var result []models.Plan
    err := r.Base.FindByPagination("plan", maps, &result, pageIndex, pageSize)
    if err != nil {
        r.Base.Log.Error(err)
        return nil, err
    }
    return &result, err
}

//根据id获取单个plan
func (r *PlanRepository) GetPlanById(id string) (*models.Plan, error) {
    var result models.Plan
    err := r.Base.FindById("plan", id, &result)
    if err != nil{
        r.Base.Log.Error(err)
        return nil, err
    }
    return &result, err
}

//新增单个Plan
func (r *PlanRepository) AddPlan(plan models.Plan) error {
    plan.ID = bson.NewObjectId()
    plan.CreateTime = time.Now().Unix()
    plan.UpdateTime = time.Now().Unix()
    err := r.Base.Insert("plan", plan)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

//编辑单个plan
func (r *PlanRepository) EditPlan(updatedPlan models.Plan) error {
    idSelector := bson.M{
        "_id": updatedPlan.ID,
    }
    updatedPlan.UpdateTime = time.Now().Unix()
    err := r.Base.Update("plan", idSelector, updatedPlan)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

//删除单个plan
func (r *PlanRepository) DeletePlanById(id string) error {
    err := r.Base.DeleteById("plan", id)
    if err != nil {
        r.Base.Log.Error(err)
        return err
    }
    return nil
}

//获取总数
func (r *PlanRepository) GetPlanCounts(maps interface{}) (int, error) {
    num, err := r.Base.Count("plan", maps)
    if err != nil {
        r.Base.Log.Error(err)
        return 0, err
    }
    return num, nil
}