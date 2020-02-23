package controller

import (
    "github.com/astaxie/beego/validation"
    "gopkg.in/mgo.v2/bson"
    "log"

    //"log"
    "net/http"
    //"strconv"
    "testgo/common/codes"
    "testgo/models"

    //"github.com/astaxie/beego/validation"
    "github.com/gin-gonic/gin"
    //"gopkg.in/mgo.v2/bson"
    "testgo/service"
)

//依赖注入
type Project struct {
    Service service.IProjectService `inject:""`
}

//根据ID获取project
func (p *Project) GetProjectById(c *gin.Context) {
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
    result := p.Service.GetProjectById(id)
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, result)
}

func (p *Project) GetProjectsByPagination(c *gin.Context) {
    valid := validation.Validation{}
    pageIndex, pageSize := GetPage(c)
    if valid.HasErrors() {
        // 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
    result := p.Service.GetProjectsByPagination(pageIndex, pageSize)
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, result)
}

func (p *Project) GetProjectsCounts(c *gin.Context) {
    counts := p.Service.GetProjectsCounts(bson.M{})
    result := bson.M{"total": counts}
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, result)
}

func (p *Project) AddProject(c *gin.Context) {
    valid := validation.Validation{}
    var project models.Project
    if err := c.BindJSON(&project); err != nil{
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
    err := p.Service.AddProject(project)
    if err != nil {
        log.Println(err)
        code := codes.InvalidParams
        RespData(c, http.StatusOK, code, nil)
    }
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, nil)
}

func (p *Project) EditProject(c *gin.Context) {
    valid := validation.Validation{}
    var project models.Project
    if err := c.BindJSON(&project); err != nil{
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
    err := p.Service.EditProject(project)
    if err != nil {
        log.Println(err)
        code := codes.InvalidParams
        RespData(c, http.StatusOK, code, nil)
    }
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, nil)
}

func (p *Project) DeleteProjectById(c *gin.Context) {
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
    err := p.Service.DeleteProjectById(id)
    if err != nil {
        log.Println(err)
        code := codes.InvalidParams
        RespData(c, http.StatusOK, code, nil)
    }
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, nil)
}

