package verifier

import (
	"github.com/gin-gonic/gin"
)

type XrayInput struct {

	ParamVal  string `form:"" json:"Landmarks" binding:"required"  uri:"" comment:"" `
}

func (params *XrayInput) BindingValidUriParams(c *gin.Context) error{

	return GetValidUriParams(c, params)
}

func (params *XrayInput) BindingValidQueryParams(c *gin.Context) error{

	return DefaultGetValidParams(c, params)

}

func (params *XrayInput) BindingValidJsonParams(c *gin.Context) error{

	return GetValidJsonParams(c, params)

}
