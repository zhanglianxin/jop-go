package jop_sdk

import (
	"strings"
	"github.com/valyala/fasthttp"
	"fmt"
	"sort"
	"crypto/md5"
	"github.com/sirupsen/logrus"
	"net/url"
	"encoding/json"
	"reflect"
	"time"
)

type Param struct {
	Method       string                 `json:"method"`                 // API接口名称
	AppKey       string                 `json:"app_key"`                // 分配给应用的AppKey
	AccessToken  string                 `json:"access_token,omitempty"` // Oauth2颁发的动态令牌,暂不支持使用
	Timestamp    string                 `json:"timestamp"`              // 时间戳，格式为yyyy-MM-dd  HH:mm:ss，时区为GMT+8
	Format       string                 `json:"format"`                 // 响应格式。暂时只支持json
	Version      string                 `json:"v"`                      // API协议版本，可选值：2.0
	SignMethod   string                 `json:"sign_method"`            // 签名的摘要算法， md5
	Sign         string                 `json:"sign"`                   // API输入参数签名结果
	ParamJson    map[string]interface{} `json:"param_json,omitempty"`
	BuyParamJson map[string]interface{} `json:"360buy_param_json,omitempty"`
}

func NewParam(method, v string) *Param {
	return &Param{
		AppKey:      AppKey,
		Timestamp:   time.Now().Local().Format("2006-01-02 15:04:05"),
		AccessToken: AccessToken,
		Format:      "json",
		SignMethod:  "md5",
		Method:      method,
		Version:     v,
	}
}

func makeRequestNew(httpMethod string, param *Param) (fasthttp.ResponseHeader, []byte) {
	param.checkRequiredParams()
	param.checkApiGateway()

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(res)
		fasthttp.ReleaseRequest(req)
	}()
	req.Header.SetUserAgent("jop-go")
	req.Header.SetMethod(strings.ToUpper(httpMethod))

	req.SetRequestURI(RootEndpoint)

	param.getSign()

	queryStr := param.getQueryString()
	switch string(req.Header.Method()) {
	case "GET":
		req.URI().SetQueryString(queryStr)
	case "POST":
		req.SetBodyString(queryStr)
	}

	err := fasthttp.Do(req, res)

	logrus.Debugln("req:", req)
	logrus.Debugln("res:", res)

	if nil != err {
		logrus.Errorf("fasthttp: %s", err)
		panic(err)
	}
	return res.Header, res.Body()
}

func (param *Param) checkApiGateway() {
	if strings.Contains(RootEndpoint, "router.jd.com") {
		if nil == param.ParamJson {
			panic("U use <union>, business params key must be <param_json>")
		}
	} else if strings.Contains(RootEndpoint, "router.jd.com") {
		if nil == param.BuyParamJson {
			panic("U use <jos>, business params key must be <360buy_param_json>")
		}
	} else {
		panic("API Gateway may be wrong")
	}
}

// 检查必选参数
func (param *Param) checkRequiredParams() {
	if "" == AppKey || "" == SecretKey {
		panic("AppKey and SecretKey must be set")
	}
	if "" == param.Format || "" == param.SignMethod || "" == param.Timestamp {
		panic("format, sign_method and timestamp must be set")
	}
	if "" == param.Method {
		panic("method is required")
	}
	if "" == param.Version {
		panic("version is required")
	}
}

//  把所有请求参数按照参数名称的 ASCII 码表顺序进行排序并拼接
func (param *Param) getConcatParams() string {
	var params map[string]interface{}
	bs, _ := json.Marshal(param)
	json.Unmarshal(bs, &params)
	s := make([]string, len(params))
	for k := range params {
		if "sign" != k {
			v := params[k]
			if "string" != reflect.TypeOf(params[k]).String() {
				valueBs, _ := json.Marshal(params[k])
				v = string(valueBs)
			}
			s = append(s, fmt.Sprintf("%s%s", k, v))
		}
	}
	sort.Strings(s)
	var concat string
	for i := range s {
		concat += s[i]
	}
	return concat
}

// 组装 HTTP 请求，将所有参数名和参数值采用 utf-8 进行 URL 编码（
func (param *Param) getQueryString() string {
	bs, _ := json.Marshal(param)
	var params map[string]interface{}
	json.Unmarshal(bs, &params)
	s := make([]string, 0)
	for k := range params {
		v := params[k]
		if "string" != reflect.TypeOf(params[k]).String() {
			valueBs, _ := json.Marshal(params[k])
			v = string(valueBs)
		}
		s = append(s, fmt.Sprintf("%s=%s", k, v))
	}
	joins := strings.Join(s, "&")
	return url.PathEscape(joins)
}

// 把 appSecret 的值拼接在字符串的两端，使用 MD5 进行加密，并转化成大写
func (param *Param) getSign() {
	concatParams := param.getConcatParams()
	param.Sign = strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(SecretKey+concatParams+SecretKey))))
}
