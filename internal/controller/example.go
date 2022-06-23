package controller

import (
	"forest/internal/dao"
	"forest/internal/response"
	"forest/internal/verifier"
	"github.com/gin-gonic/gin"
)

type ExampleController struct {
}

func ExampleRegister(router *gin.RouterGroup) {
	example := ExampleController{}
	router.GET("/example", example.GetEnvExamples)
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
