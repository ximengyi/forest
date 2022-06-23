package verifier

import (
	"forest/internal/constant"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"strings"
)

type PageInput struct {
	PageSize int `form:"size" json:"size" comment:"每页记录数" validate:"numeric" example:"10"`
	Page     int `form:"page" json:"page" comment:"页数" validate:"numeric" example:"1"`
}

func DefaultGetValidParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBind(params); err != nil {
		return err
	}
	//获取验证器
	valid, err := GetValidator(c)
	if err != nil {
		return err
	}
	//获取翻译器
	trans, err := GetTranslation(c)
	if err != nil {
		return err
	}
	err = valid.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

func GetValidator(c *gin.Context) (*validator.Validate, error) {
	val, ok := c.Get(constant.ValidatorKey)
	if !ok {
		return nil, errors.New("未设置验证器")
	}
	validator, ok := val.(*validator.Validate)
	if !ok {
		return nil, errors.New("获取验证器失败")
	}
	return validator, nil
}

func GetTranslation(c *gin.Context) (ut.Translator, error) {
	trans, ok := c.Get(constant.TranslatorKey)
	if !ok {
		return nil, errors.New("未设置翻译器")
	}
	translator, ok := trans.(ut.Translator)
	if !ok {
		return nil, errors.New("获取翻译器失败")
	}
	return translator, nil
}

func GetValidUriParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBindUri(params); err != nil {
		return err
	}
	//获取验证器
	valid, err := GetValidator(c)
	if err != nil {
		return err
	}
	//获取翻译器
	trans, err := GetTranslation(c)
	if err != nil {
		return err
	}
	err = valid.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}
