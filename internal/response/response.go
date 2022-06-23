package response

import (
	"encoding/json"
	"fmt"
	"forest/internal/constant"
	"forest/pkg/bootstrap"
	"forest/pkg/log"
	"github.com/gin-gonic/gin"
	"strings"
)

type Response struct {
	ErrorCode int         `json:"code"`
	ErrorMsg  string      `json:"message"`
	Success   bool        `json:"success"`
	Data      interface{} `json:"data"`
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

	resp := &Response{ErrorCode: code, Success: false, ErrorMsg: err.Error(), Data: ""}
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

	resp := &Response{ErrorCode: constant.SuccessCode, Success: true, ErrorMsg: "", Data: data}
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

	resp := &Response{ErrorCode: constant.SuccessCode, Success: true, ErrorMsg: errorMsg, Data: data}
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

	resp := &Response{ErrorCode: code, Success: false, ErrorMsg: data, Data: data}
	c.JSON(200, resp)
	logResp := &LogResponse{
		Response: *resp,
		TraceId:  traceId,
		Stack:    stack,
	}
	logResponse, _ := json.Marshal(logResp)
	c.Set("response", logResponse)

}
