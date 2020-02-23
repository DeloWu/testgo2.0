package repository

import (
    "testgo/models"
)

type IPlanRepository interface {
    //分页返回Plans
    GetPlansByPagination(maps interface{}, pageIndex int, pageSize int) (*[]models.Plan, error)
    //根据ID获取plan
    GetPlanById(id string) (*models.Plan, error)
    //根据条件获取多个plans
    //GetPlans(maps interface{}) *[]models.Plan
    //根据条件获取单个plan
    //GetPlan(maps interface{}) *models.Plan
    //新增plan
    AddPlan(plan models.Plan) error
    //编辑plan
    EditPlan(updatedPlan models.Plan) error
    //删除plan
    DeletePlanById(id string) error
    //获取所有plans数量
    GetPlanCounts(maps interface{}) (int, error)

}