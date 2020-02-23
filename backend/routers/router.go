package routers

import (
    "github.com/facebookgo/inject"
    "github.com/gin-gonic/gin"
    "log"
    "testgo/common/datasource"
    "testgo/common/logger"
    "testgo/common/middleware/cors"
    //"testgo/common/middleware/jwt"
    "testgo/common/setting"
    "testgo/controller"
    "testgo/repository"
    "testgo/service"
)

//InitRouter 初始化Router
func InitRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(cors.CorsHandler())
    r.Use(gin.Recovery())
    gin.SetMode(setting.Config.APP.RunMode)
    Configure(r)
    return r
}

//Configure 配置router
func Configure(r *gin.Engine) {
    //controller declare
    var project controller.Project
    //var tag cv1.Tag
    //var myjwt jwt.JWT
    db := datasource.Db{}
    zap := logger.Logger{}
    //Injection
    var injector inject.Graph
    err := injector.Provide(
        &inject.Object{Value: &project},
        &inject.Object{Value: &db},
        &inject.Object{Value: &zap},
        //&inject.Object{Value: &myjwt},
        &inject.Object{Value: &repository.ProjectRepository{}},
        &inject.Object{Value: &service.ProjectService{}},
        &inject.Object{Value: &repository.BaseRepository{}},

    )
    if err != nil {
        log.Fatal("inject fatal: ", err)
    }
    //zap log init
    zap.Init()
    //database connect
    err = db.Connect()
    if err != nil {
        log.Fatal("db fatal:", err)
    }

    if err := injector.Populate(); err != nil {
        log.Fatal("injector fatal: ", err)
    }
    //var authMiddleware = myjwt.GinJWTMiddlewareInit(jwt.AllUserAuthorizator)
    //r.NoRoute(authMiddleware.MiddlewareFunc(), jwt.NoRouteHandler)
    //r.POST("/login", authMiddleware.LoginHandler)
    apiV1Group := r.Group("/api/v1")
    {
        //查询projects总数
        apiV1Group.GET("/project-total", project.GetProjectsCounts)
        //查询单个project详情
        apiV1Group.GET("/project/:id", project.GetProjectById)
        //分页查询
        apiV1Group.GET("/projects", project.GetProjectsByPagination)
        //新增
        apiV1Group.POST("/project", project.AddProject)
        //更新
        apiV1Group.PUT("/project", project.EditProject)
        //删除
        apiV1Group.DELETE("/project/:id", project.DeleteProjectById)
    }
}