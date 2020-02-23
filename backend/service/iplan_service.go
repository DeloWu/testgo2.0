package service

import (
    "testgo/models"
)

type IPlanService interface {
    //根据ID获取plan
    GetPlanById(id string) *models.Plan
    //分页获取plans
    GetPlansByPagination(pageIndex int, pageSize int) *[]models.Plan
    //获取总数
    GetPlansCounts(maps interface{}) int
    //新增plan
    AddPlan(plan models.Plan) error
    //编辑单个plan
    EditPlan(updatedPlan models.Plan) error
    //删除单个plan
    DeletePlanById(id string) error
}
