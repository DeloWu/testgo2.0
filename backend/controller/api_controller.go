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
type Api struct {
	Service service.IApiService `inject:""`
}

//根据ID获取api
func (p *Api) GetApiById(c *gin.Context) {
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
	result := p.Service.GetApiById(id)
	var code int
	if result != nil {
		code = codes.SUCCESS
	} else {
		code = codes.NotFound
	}
	RespData(c, http.StatusOK, code, result)
	return
}

func (p *Api) GetApisByPagination(c *gin.Context) {
	valid := validation.Validation{}
	pageIndex, pageSize := GetPage(c)
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	result := p.Service.GetApisByPagination(pageIndex, pageSize)
	var code int
	if *result != nil {
		code = codes.SUCCESS
	} else {
		code = codes.NotFound
	}
	RespData(c, http.StatusOK, code, result)
	return
}

func (p *Api) GetApisCounts(c *gin.Context) {
	counts := p.Service.GetApisCounts(bson.M{})
	result := bson.M{"total": counts}
	code := codes.SUCCESS
	RespData(c, http.StatusOK, code, result)
	return
}

func (p *Api) AddApi(c *gin.Context) {
	valid := validation.Validation{}
	var api models.Api
	if err := c.BindJSON(&api); err != nil {
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
	err := p.Service.AddApi(api)
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

func (p *Api) EditApi(c *gin.Context) {
	valid := validation.Validation{}
	var api models.Api
	if err := c.BindJSON(&api); err != nil {
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
	err := p.Service.EditApi(api)
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

func (p *Api) DeleteApiById(c *gin.Context) {
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
	err := p.Service.DeleteApiById(id)
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
