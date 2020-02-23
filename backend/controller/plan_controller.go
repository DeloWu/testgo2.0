package controller

import (
    "github.com/astaxie/beego/validation"
    "gopkg.in/mgo.v2/bson"
    "log"
    "net/http"
    //"strconv"
    "testgo/common/codes"
    "testgo/models"
    "github.com/gin-gonic/gin"
    "testgo/service"
)

//依赖注入
type Plan struct {
    Service service.IPlanService `inject:""`
}

//根据ID获取Plan
func (p *Plan) GetPlanById(c *gin.Context) {
    valid := validation.Validation{}
    id := c.Param("id")
    valid.Required(id, "id")
    if valid.HasErrors() {
        // 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
    result := p.Service.GetPlanById(id)
    var code int
    if result != nil {
        code = codes.SUCCESS
    }else {
        code = codes.NotFound
    }
    RespData(c, http.StatusOK, code, result)
}

func (p *Plan) GetPlansByPagination(c *gin.Context) {
    valid := validation.Validation{}
    pageIndex, pageSize := GetPage(c)
    if valid.HasErrors() {
        // 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
    result := p.Service.GetPlansByPagination(pageIndex, pageSize)
    var code int
    if *result != nil {
       code = codes.SUCCESS
    }else {
       code = codes.NotFound
    }
    RespData(c, http.StatusOK, code, result)
}

func (p *Plan) GetPlansCounts(c *gin.Context) {
    counts := p.Service.GetPlansCounts(bson.M{})
    result := bson.M{"total": counts}
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, result)
}

func (p *Plan) AddPlan(c *gin.Context) {
    valid := validation.Validation{}
    var Plan models.Plan
    if err := c.BindJSON(&Plan); err != nil{
        log.Println(err)
        code := codes.InvalidParams
        RespData(c, http.StatusOK, code, nil)
    }
    if valid.HasErrors() {
        // 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
    err := p.Service.AddPlan(Plan)
    if err != nil {
        log.Println(err)
        code := codes.InvalidParams
        RespData(c, http.StatusOK, code, nil)
    }
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, nil)
}

func (p *Plan) EditPlan(c *gin.Context) {
    valid := validation.Validation{}
    var Plan models.Plan
    if err := c.BindJSON(&Plan); err != nil{
        log.Println(err)
        code := codes.InvalidParams
        RespData(c, http.StatusOK, code, nil)
    }
    if valid.HasErrors() {
        // 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
    err := p.Service.EditPlan(Plan)
    if err != nil {
        log.Println(err)
        code := codes.InvalidParams
        RespData(c, http.StatusOK, code, nil)
    }
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, nil)
}

func (p *Plan) DeletePlanById(c *gin.Context) {
    valid := validation.Validation{}
    id := c.Param("id")
    valid.Required(id, "id")
    if valid.HasErrors() {
        // 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
    err := p.Service.DeletePlanById(id)
    if err != nil {
        log.Println(err)
        code := codes.InvalidParams
        RespData(c, http.StatusOK, code, nil)
    }
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, nil)
}

