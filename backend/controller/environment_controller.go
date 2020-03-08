package controller

import (
	"github.com/astaxie/beego/validation"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	//"strconv"
	"github.com/gin-gonic/gin"
	"testgo/common/codes"
	"testgo/models"
	"testgo/service"
)

//依赖注入
type Environment struct {
	Service service.IEnvironmentService `inject:""`
}

//根据ID获取environment
func (e *Environment) GetEnvironmentById(c *gin.Context) {
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
	result := e.Service.GetEnvironmentById(id)
	var code int
	if result != nil {
		code = codes.SUCCESS
	} else {
		code = codes.NotFound
	}
	RespData(c, http.StatusOK, code, result)
	return
}

func (e *Environment) GetEnvironmentsByPagination(c *gin.Context) {
	valid := validation.Validation{}
	pageIndex, pageSize := GetPage(c)
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	result := e.Service.GetEnvironmentsByPagination(pageIndex, pageSize)
	var code int
	if *result == nil {
		code = codes.NotFound
	} else {
		code = codes.SUCCESS
	}
	RespData(c, http.StatusOK, code, result)
	return
}

func (e *Environment) GetEnvironmentsCounts(c *gin.Context) {
	counts := e.Service.GetEnvironmentsCounts(bson.M{})
	result := bson.M{"total": counts}
	code := codes.SUCCESS
	RespData(c, http.StatusOK, code, result)
	return
}

func (e *Environment) AddEnvironment(c *gin.Context) {
	valid := validation.Validation{}
	var environment models.Environment
	if err := c.BindJSON(&environment); err != nil {
		log.Println(err)
		code := codes.InvalidParams
		RespData(c, http.StatusOK, code, nil)
		return
	}
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	err := e.Service.AddEnvironment(environment)
	if err != nil {
		log.Println(err)
		code := codes.InvalidParams
		RespData(c, http.StatusOK, code, nil)
		return
	}
	code := codes.SUCCESS
	RespData(c, http.StatusOK, code, nil)
	return
}

func (e *Environment) EditEnvironment(c *gin.Context) {
	valid := validation.Validation{}
	var environment models.Environment
	if err := c.BindJSON(&environment); err != nil {
		log.Println(err)
		code := codes.InvalidParams
		RespData(c, http.StatusOK, code, nil)
		return
	}
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	err := e.Service.EditEnvironment(environment)
	if err != nil {
		log.Println(err)
		code := codes.InvalidParams
		RespData(c, http.StatusOK, code, nil)
		return
	}
	code := codes.SUCCESS
	RespData(c, http.StatusOK, code, nil)
	return
}

func (e *Environment) DeleteEnvironmentById(c *gin.Context) {
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
	err := e.Service.DeleteEnvironmentById(id)
	if err != nil {
		log.Println(err)
		code := codes.InvalidParams
		RespData(c, http.StatusOK, code, nil)
		return
	}
	code := codes.SUCCESS
	RespData(c, http.StatusOK, code, nil)
	return
}
