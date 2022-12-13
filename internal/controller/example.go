package controller

import (
	"fmt"
	"forest/internal/dao"
	"forest/internal/response"
	"forest/internal/verifier"
	"forest/pkg/db"
	"github.com/gin-gonic/gin"
	"time"
)

type ExampleController struct {
}

func ExampleRegister(router *gin.RouterGroup) {
	example := ExampleController{}
	router.GET("/example", example.GetEnvExamples)
	router.GET("/redis", example.RedisDemo)
}

func (example *ExampleController) RedisDemo(c *gin.Context) {

	rdb, err := db.RedisConnFactory(0)
	if err != nil {
		fmt.Println(err)
	}
	key := "mytest"
	rdb.Set(c, key, "hello redis is ok", 3600*time.Second)
	res := rdb.Get(c, key)
	fmt.Println(key)
	response.Success(c, res.Val())

	return

}

func (example *ExampleController) GetEnvExamples(c *gin.Context) {
	params := &verifier.EnvInput{}
	if err := params.BindingValidUriParams(c); err != nil {
		response.Error(c, 422, err)
		return
	}
	if err := params.BindingValidQueryParams(c); err != nil {
		response.Error(c, 422, err)
		return
	}
	page := 1
	pageSize := 10
	if params.Page > 0 {
		page = params.Page
	}
	if params.PageSize > 0 {
		pageSize = params.PageSize
	}

	res, _, err := (&dao.Envtest{}).PageList(page, pageSize, "")
	if err != nil {
		response.Error(c, 4000, err)
		return
	}

	response.Success(c, res)
	return
}
