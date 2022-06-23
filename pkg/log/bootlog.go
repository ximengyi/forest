package log

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

// 通用DLTag常量定义
const (
	DLTagUndefind      = "_undef"
	DLTagMySqlFailed   = "_com_mysql_failure"
	DLTagRedisFailed   = "_com_redis_failure"
	DLTagMySqlSuccess  = "_com_mysql_success"
	DLTagRedisSuccess  = "_com_redis_success"
	DLTagThriftFailed  = "_com_thrift_failure"
	DLTagThriftSuccess = "_com_thrift_success"
	DLTagHTTPSuccess   = "_com_http_success"
	DLTagHTTPFailed    = "_com_http_failure"
	DLTagTCPFailed     = "_com_tcp_failure"
	DLTagRequestIn     = "_com_request_in"
	DLTagRequestOut    = "_com_request_out"
)

const (
	_dlTag          = "dltag"
	_traceId        = "traceid"
	_spanId         = "spanid"
	_childSpanId    = "cspanid"
	_dlTagBizPrefix = "_com_"
	_dlTagBizUndef  = "_com_undef"
)

//var Log *BootLogger
var LocalIP = net.ParseIP("127.0.0.1")

type BootTrace struct {
	TraceId     string
	SpanId      string
	Caller      string
	SrcMethod   string
	HintCode    int64
	HintContent string
}

type TraceContext struct {
	BootTrace
	CSpanId string
}

type BootLogger struct {
}

func init()  {
	ips := GetLocalIPs()
	if len(ips) > 0 {
		LocalIP = ips[0]
	}
}

func (l *BootLogger) TagInfo( trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	Info(parseParams(m))
}

func (l *BootLogger) TagWarn(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	Warn(parseParams(m))
}

func (l *BootLogger) TagError(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	Error(parseParams(m))
}

func (l *BootLogger) TagTrace(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	Info(parseParams(m))
}

func (l *BootLogger) TagDebug(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	Debug(parseParams(m))
}

//func (l *BootLogger) Close() {
//	Close()
//}

// 生成业务dltag
func CreateBizDLTag(tagName string) string {
	if tagName == "" {
		return _dlTagBizUndef
	}

	return _dlTagBizPrefix + tagName
}

// 校验dltag合法性
func checkDLTag(dltag string) string {
	if strings.HasPrefix(dltag, _dlTagBizPrefix) {
		return dltag
	}

	if strings.HasPrefix(dltag, "_com_") {
		return dltag
	}

	if dltag == DLTagUndefind {
		return dltag
	}
	return dltag
}

//map格式化为string
func parseParams(m map[string]interface{}) string {
	var dltag string = "_undef"
	if _dltag, _have := m["dltag"]; _have {
		if __val, __ok := _dltag.(string); __ok {
			dltag = __val
		}
	}
	for _key, _val := range m {
		if _key == "dltag" {
			continue
		}
		dltag = dltag + "||" + fmt.Sprintf("%v=%+v", _key, _val)
	}
	dltag = strings.Trim(fmt.Sprintf("%q", dltag), "\"")
	return dltag
}


var bootins *BootLogger
var bootonce sync.Once

func GetBootLogInstance() *BootLogger {
	bootonce.Do(func() {
		bootins = &BootLogger{}
	})
	return bootins
}


func NewTrace() *TraceContext {
	trace := &TraceContext{}
	trace.TraceId = GetTraceId()
	trace.SpanId  = NewSpanId()
	return trace
}

func NewSpanId() string {
	timestamp := uint32(time.Now().Unix())
	ipToLong := binary.BigEndian.Uint32(LocalIP.To4())
	b := bytes.Buffer{}
	b.WriteString(fmt.Sprintf("%08x", ipToLong^timestamp))
	b.WriteString(fmt.Sprintf("%08x", rand.Int31()))
	return b.String()
}

func GetTraceId() (traceId string) {
	return calcTraceId(LocalIP.String())
}

func calcTraceId(ip string) (traceId string) {
	now := time.Now()
	timestamp := uint32(now.Unix())
	timeNano := now.UnixNano()
	pid := os.Getpid()

	b := bytes.Buffer{}
	netIP := net.ParseIP(ip)
	if netIP == nil {
		b.WriteString("00000000")
	} else {
		b.WriteString(hex.EncodeToString(netIP.To4()))
	}
	b.WriteString(fmt.Sprintf("%08x", timestamp&0xffffffff))
	b.WriteString(fmt.Sprintf("%04x", timeNano&0xffff))
	b.WriteString(fmt.Sprintf("%04x", pid&0xffff))
	b.WriteString(fmt.Sprintf("%06x", rand.Int31n(1<<24)))
	b.WriteString("b0") // 末两位标记来源,b0为go

	return b.String()
}

func GetLocalIPs() (ips []net.IP) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP)
			}
		}
	}
	return ips
}
