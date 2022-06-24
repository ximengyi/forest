package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"forest/internal/constant"
	"forest/pkg/bootstrap"
	"forest/pkg/log"
	"github.com/gin-gonic/gin"
	"strings"
)



type Response struct {
	Code      int         `json:"Code"`
	Message   string      `json:"Message"`
	RequestId string      `json:"RequestId"`
	Data      interface{} `json:"Data"`
}


type LogResponse struct {
	Response
	TraceId interface{} `json:"trace_id"`
	Stack   interface{} `json:"stack"`
}

func Error(c *gin.Context, code int, err error) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*log.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	stack := ""
	if c.Query("is_debug") == "1" || bootstrap.GetAppEnv() == "dev" {
		stack = strings.Replace(fmt.Sprintf("%+v", err), err.Error()+"\n", "", -1)
	}

	resp := &Response{Code: code, RequestId: traceId, Message: err.Error(), Data: ""}
	c.JSON(200, resp)
	logResp := &LogResponse{
		Response: *resp,
		TraceId:  traceId,
		Stack:    stack,
	}
	logResponse, _ := json.Marshal(logResp)
	c.Set("response", string(logResponse))
	c.AbortWithError(200, err)
}

func Success(c *gin.Context, data interface{}) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*log.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	resp := &Response{Code: constant.SuccessCode, RequestId: traceId, Message: "", Data: data}
	c.JSON(200, resp)
	logResp := &LogResponse{
		Response: *resp,
		TraceId:  traceId,
	}
	logResponse, _ := json.Marshal(logResp)
	c.Set("response", string(logResponse))
}

func MessageSuccess(c *gin.Context, errorMsg string, data interface{}) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*log.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	resp := &Response{Code: constant.SuccessCode,  RequestId: traceId, Message: errorMsg,  Data: data}
	c.JSON(200, resp)
	logResp := &LogResponse{
		Response: *resp,
		TraceId:  traceId,
	}
	logResponse, _ := json.Marshal(logResp)
	c.Set("response", string(logResponse))
}

func ErrorMsg(c *gin.Context, code int, data string) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*log.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	stack := ""
	if c.Query("is_debug") == "1" || bootstrap.GetAppEnv() == "dev" {
		stack = data
	}

	resp := &Response{Code: code, RequestId: traceId, Message: data, Data: data}
	c.JSON(200, resp)
	logResp := &LogResponse{
		Response: *resp,
		TraceId:  traceId,
		Stack:    stack,
	}
	logResponse, _ := json.Marshal(logResp)
	c.Set("response", logResponse)

}


func Error404(c *gin.Context ) {
	err := errors.New("404 not found")
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*log.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	stack := ""
	if c.Query("is_debug") == "1" || bootstrap.GetAppEnv() == "dev" {
		stack = strings.Replace(fmt.Sprintf("%+v", err), err.Error()+"\n", "", -1)
	}

	resp := &Response{Code: 404, RequestId: traceId, Message: err.Error(), Data: ""}
	c.JSON(200, resp)
	logResp := &LogResponse{
		Response: *resp,
		TraceId:  traceId,
		Stack:    stack,
	}
	logResponse, _ := json.Marshal(logResp)
	c.Set("response", string(logResponse))
	c.AbortWithError(404, err)
}