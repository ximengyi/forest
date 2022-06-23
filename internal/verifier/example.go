package verifier

import (
	"github.com/gin-gonic/gin"
)

type EnvInput struct {
	PageInput
	Id   int `json:"id" uri:"id" comment:"id" validate:"numeric"`
}

func (params *EnvInput) BindingValidUriParams(c *gin.Context) error{

	return GetValidUriParams(c, params)
}

func (params *EnvInput) BindingValidQueryParams(c *gin.Context) error{

	return DefaultGetValidParams(c, params)

}

