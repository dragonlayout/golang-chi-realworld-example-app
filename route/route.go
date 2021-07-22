package route

import (
	"fmt"
	"github.com/dragonlayout/golang-chi-realworld-example-app/model"
	"github.com/dragonlayout/golang-chi-realworld-example-app/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func New(svc service.UserService) *gin.Engine {
	// Creates a router without any middleware by default
	route := gin.New()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with
	// GIN_MODE=release.
	// By default gin.DefaultWriter = os.stdout
	// TODO: logger 配置
	route.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one
	// TODO: 全局异常处理
	//route.Use(gin.Recovery())
	route.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		//if err, ok := err.(string); ok {
		//	c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		//}
		c.JSON(http.StatusOK, gin.H{
			"errCode": -1,
			"errMsg":  "服务异常, 请稍后再试",
			"sysTime": time.Now().UnixNano() / int64(time.Millisecond), // java: System.currentTimeMillis()
		})
		c.AbortWithStatus(http.StatusOK)
	}))

	// health check
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// TODO: 添加路由
	// 1. 用户注册
	route.POST("/api/users", func(c *gin.Context) {
		// request model bind and validate
		// post body
		var requestBody model.UserRegistrationRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err := svc.Registration(requestBody)
		if err != nil {
			fmt.Println(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user": "response",
		})
	})

	return route
}
