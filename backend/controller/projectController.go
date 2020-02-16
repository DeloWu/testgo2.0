package controller

import (
    "github.com/astaxie/beego/validation"
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
    var viewProject models.Project
    valid := validation.Validation{}
    id := c.Query("id")
    valid.Required(id, "id")
    if valid.HasErrors() {
        // 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
    //id = bson.ObjectIdHex(id)
    viewProject = p.Service.GetProjectById(id)
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, &viewProject)
}

func (p *Project) GetProjectsByPagination(c *gin.Context) {
    var viewProjects *[]models.Project
    viewProjects = new([]models.Project)
    valid := validation.Validation{}
    pageIndex, pageSize := GetPage(c)
    if valid.HasErrors() {
        // 如果有错误信息，证明验证没通过
        // 打印错误信息
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
    log.Println("===============================")
    log.Println(viewProjects)
    log.Println("===============================")
    log.Println(*viewProjects)
    log.Println("===============================")

    log.Println("===============================")
    log.Println("pageIndex:", pageIndex, "  pageSize:", pageSize)
    log.Println("===============================")

    viewProjects = p.Service.GetProjectsByPagination(pageIndex, pageSize)
    //viewProjects := map[string]string{"foo": "bar"}
    log.Println("===============================")
    log.Println(viewProjects)
    log.Println("===============================")
    log.Println(*viewProjects)
    log.Println("===============================")
    code := codes.SUCCESS
    RespData(c, http.StatusOK, code, viewProjects)

}

//调试使用
func MockData(c *gin.Context) {
    c.JSON(200, gin.H{
        "BlogL": "www.hello.com",
        "wechat": "myWechat",
    })
}