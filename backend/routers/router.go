package routers

import (
    "github.com/facebookgo/inject"
    "github.com/gin-gonic/gin"
    "log"
    "testgo/common/datasource"
    "testgo/common/logger"
    "testgo/common/middleware/cors"
    "testgo/common/middleware/jwt"
    "testgo/common/setting"
    "testgo/common/controller"
    "testgo/common/repository"
    "testgo/common/service"
)
