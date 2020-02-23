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
    var environment controller.Environment
    var api controller.Api
    var ca controller.Case
    var suite controller.Suite
    var plan controller.Plan
    //var tag cv1.Tag
    //var myjwt jwt.JWT
    db := datasource.Db{}
    zap := logger.Logger{}
    //Injection
    var injector inject.Graph
    err := injector.Provide(
        &inject.Object{Value: &project},
        &inject.Object{Value: &environment},
        &inject.Object{Value: &api},
        &inject.Object{Value: &ca},
        &inject.Object{Value: &suite},
        &inject.Object{Value: &plan},
        &inject.Object{Value: &db},
        &inject.Object{Value: &zap},
        //&inject.Object{Value: &myjwt},
        &inject.Object{Value: &repository.ProjectRepository{}},
        &inject.Object{Value: &service.ProjectService{}},
        &inject.Object{Value: &repository.EnvironmentRepository{}},
        &inject.Object{Value: &service.EnvironmentService{}},
        &inject.Object{Value: &repository.ApiRepository{}},
        &inject.Object{Value: &service.ApiService{}},
        &inject.Object{Value: &repository.CaseRepository{}},
        &inject.Object{Value: &service.CaseService{}},
        &inject.Object{Value: &repository.SuiteRepository{}},
        &inject.Object{Value: &service.SuiteService{}},
        &inject.Object{Value: &repository.PlanRepository{}},
        &inject.Object{Value: &service.PlanService{}},
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
        //查询总数
        apiV1Group.GET("/project-total", project.GetProjectsCounts)
        //查询单个详情
        apiV1Group.GET("/project/:id", project.GetProjectById)
        //分页查询
        apiV1Group.GET("/projects", project.GetProjectsByPagination)
        //新增
        apiV1Group.POST("/project", project.AddProject)
        //更新
        apiV1Group.PUT("/project", project.EditProject)
        //删除
        apiV1Group.DELETE("/project/:id", project.DeleteProjectById)

        //查询总数
        apiV1Group.GET("/environment-total", environment.GetEnvironmentsCounts)
        //查询单个详情
        apiV1Group.GET("/environment/:id", environment.GetEnvironmentById)
        //分页查询
        apiV1Group.GET("/environments", environment.GetEnvironmentsByPagination)
        //新增
        apiV1Group.POST("/environment", environment.AddEnvironment)
        //更新
        apiV1Group.PUT("/environment", environment.EditEnvironment)
        //删除
        apiV1Group.DELETE("/environment/:id", environment.DeleteEnvironmentById)

        //查询总数
        apiV1Group.GET("/api-total", api.GetApisCounts)
        //查询单个详情
        apiV1Group.GET("/api/:id", api.GetApiById)
        //分页查询
        apiV1Group.GET("/apis", api.GetApisByPagination)
        //新增
        apiV1Group.POST("/api", api.AddApi)
        //更新
        apiV1Group.PUT("/api", api.EditApi)
        //删除
        apiV1Group.DELETE("/api/:id", api.DeleteApiById)

        //查询总数
        apiV1Group.GET("/case-total", ca.GetCasesCounts)
        //查询单个详情
        apiV1Group.GET("/case/:id", ca.GetCaseById)
        //分页查询
        apiV1Group.GET("/cases", ca.GetCasesByPagination)
        //新增
        apiV1Group.POST("/case", ca.AddCase)
        //更新
        apiV1Group.PUT("/case", ca.EditCase)
        //删除
        apiV1Group.DELETE("/case/:id", ca.DeleteCaseById)

        //查询总数
        apiV1Group.GET("/suite-total", suite.GetSuitesCounts)
        //查询单个详情
        apiV1Group.GET("/suite/:id", suite.GetSuiteById)
        //分页查询
        apiV1Group.GET("/suites", suite.GetSuitesByPagination)
        //新增
        apiV1Group.POST("/suite", suite.AddSuite)
        //更新
        apiV1Group.PUT("/suite", suite.EditSuite)
        //删除
        apiV1Group.DELETE("/suite/:id", suite.DeleteSuiteById)

        //查询总数
        apiV1Group.GET("/plan-total", plan.GetPlansCounts)
        //查询单个详情
        apiV1Group.GET("/plan/:id", plan.GetPlanById)
        //分页查询
        apiV1Group.GET("/plans", plan.GetPlansByPagination)
        //新增
        apiV1Group.POST("/plan", plan.AddPlan)
        //更新
        apiV1Group.PUT("/plan", plan.EditPlan)
        //删除
        apiV1Group.DELETE("/plan/:id", plan.DeletePlanById)
    }
}