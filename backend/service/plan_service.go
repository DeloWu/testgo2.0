package service

import (
    "gopkg.in/mgo.v2/bson"
    "testgo/models"
    "testgo/repository"
)

//依赖注入
type PlanService struct {
    Repository repository.IPlanRepository `inject:""`
}

//根据ID获取plan
func (s *PlanService) GetPlanById(id string) *models.Plan{
     result, err := s.Repository.GetPlanById(id)
     if err != nil{
         return nil
     }
     return result
}

//获取分页数据
func (s *PlanService) GetPlansByPagination(pageIndex int, pageSize int) *[]models.Plan{
    maps := bson.M{}
    result, err := s.Repository.GetPlansByPagination(maps, pageIndex, pageSize)
    if err != nil{
        return nil
    }
    return result
}

//获取总数
func (s *PlanService) GetPlansCounts(maps interface{}) int{
    counts, err := s.Repository.GetPlanCounts(maps)
    if err != nil{
        return 0
    }
    return counts
}

//新增plan
func (s *PlanService) AddPlan(plan models.Plan) error {
    return s.Repository.AddPlan(plan)
}

//编辑单个plan
func (s *PlanService) EditPlan(updatedPlan models.Plan) error{
    return s.Repository.EditPlan(updatedPlan)
}

//删除单个plan
func (s *PlanService) DeletePlanById(id string) error{
    return s.Repository.DeletePlanById(id)
}