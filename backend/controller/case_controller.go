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
type Case struct {
	Service service.ICaseService `inject:""`
}

//根据ID获取case
func (p *Case) GetCaseById(c *gin.Context) {
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
	result := p.Service.GetCaseById(id)
	var code int
	if result != nil {
		code = codes.SUCCESS
	} else {
		code = codes.NotFound
	}
	RespData(c, http.StatusOK, code, result)
	return
}

func (p *Case) GetCasesByPagination(c *gin.Context) {
	valid := validation.Validation{}
	pageIndex, pageSize := GetPage(c)
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	result := p.Service.GetCasesByPagination(pageIndex, pageSize)
	var code int
	if *result != nil {
		code = codes.SUCCESS
	} else {
		code = codes.NotFound
	}
	RespData(c, http.StatusOK, code, result)
	return
}

func (p *Case) GetCasesCounts(c *gin.Context) {
	counts := p.Service.GetCasesCounts(bson.M{})
	result := bson.M{"total": counts}
	code := codes.SUCCESS
	RespData(c, http.StatusOK, code, result)
	return
}

func (p *Case) AddCase(c *gin.Context) {
	valid := validation.Validation{}
	var ca models.Case
	if err := c.BindJSON(&ca); err != nil {
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
	err := p.Service.AddCase(ca)
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

func (p *Case) EditCase(c *gin.Context) {
	valid := validation.Validation{}
	var ca models.Case
	if err := c.BindJSON(&ca); err != nil {
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
	err := p.Service.EditCase(ca)
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

func (p *Case) DeleteCaseById(c *gin.Context) {
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
	err := p.Service.DeleteCaseById(id)
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
