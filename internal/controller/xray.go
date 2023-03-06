package controller

import (
	"fmt"
	"forest/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"time"
)

type XrayController struct {
}

func XrayRegister(router *gin.RouterGroup) {
	Xray := XrayController{}
	router.Any("/printResult", Xray.PrintResult)
	router.GET("/jwt", Xray.JwtToken)
}

func (Xray *XrayController) PrintResult(c *gin.Context) {
	//params := &verifier.XrayInput{}
	//if err := params.BindingValidJsonParams(c); err != nil {
	//	response.Error(c, 422, err)
	//	return
	//}
	data, _ := c.GetRawData()
	fmt.Println(string(data))
	response.Success(c, string(data))
	return
}

func (Xray *XrayController) JwtToken(c *gin.Context) {

	key := "x46SBqsNNx9kwTts"
	secret := "Wjczydjn46SBqsNNx9kwTts"
	playloads := map[string]string{"Red": "#da1337"}
	claims := make(jwt.MapClaims)
	var iat int64 = time.Now().Unix()
	claims["exp"] = iat + 8600
	claims["iat"] = iat
	claims["key"] = key
	for k, v := range playloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	data, _ := token.SignedString([]byte(secret))

	response.Success(c, data)
	return
}
