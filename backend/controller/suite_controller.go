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
type Suite struct {
    Service service.ISuiteService `inject:""`
}

//根据ID获取suite
func (p *Suite) GetSuiteById(c *gin.Context) {
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
    result := p.Service.GetSuiteById(id)
    var code int
    if result != nil {
        code = codes.SUCCESS
    }else {
        code = codes.NotFound
    }
    RespData(c, http.StatusOK, code, result)
}

func (p *Suite) GetSuitesByPagination(c *gin.Context) {
    valid := validation.Validation{}
    pageIndex, pageSize := GetPage(c)
    if valid.HasErrors() {
        // 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
    result := p.Service.GetSuitesByPagination(pageIndex, pageSize)
    var code int
    if *result != nil {
       code = codes.SUCCESS
    }else {
       code = codes.NotFound
    }
    RespData(c, http.StatusOK, code, result)
}

func (p *Suite) GetSuitesCounts(c *gin.Context) {
    counts := p.Service.GetSuitesCounts(bson.M{})
    result := bson.M{"total": counts}
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, result)
}

func (p *Suite) AddSuite(c *gin.Context) {
    valid := validation.Validation{}
    var suite models.Suite
    if err := c.BindJSON(&suite); err != nil{
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
    err := p.Service.AddSuite(suite)
    if err != nil {
        log.Println(err)
        code := codes.InvalidParams
        RespData(c, http.StatusOK, code, nil)
    }
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, nil)
}

func (p *Suite) EditSuite(c *gin.Context) {
    valid := validation.Validation{}
    var suite models.Suite
    if err := c.BindJSON(&suite); err != nil{
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
    err := p.Service.EditSuite(suite)
    if err != nil {
        log.Println(err)
        code := codes.InvalidParams
        RespData(c, http.StatusOK, code, nil)
    }
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, nil)
}

func (p *Suite) DeleteSuiteById(c *gin.Context) {
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
    err := p.Service.DeleteSuiteById(id)
    if err != nil {
        log.Println(err)
        code := codes.InvalidParams
        RespData(c, http.StatusOK, code, nil)
    }
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, nil)
}

