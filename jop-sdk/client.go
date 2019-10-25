package jop_sdk

import (
	"strings"
	"github.com/valyala/fasthttp"
	"time"
	"fmt"
	"sort"
	"crypto/md5"
	"github.com/sirupsen/logrus"
	"net/url"
)

// "method"       // API接口名称
// "app_key"      // 分配给应用的AppKey
// "access_token" // Oauth2颁发的动态令牌,暂不支持使用
// "timestamp"    // 时间戳，格式为yyyy-MM-dd  HH:mm:ss，时区为GMT+8
// "format"       // 响应格式。暂时只支持json
// "v"            // API协议版本，可选值：2.0
// "sign_method"  // 签名的摘要算法， md5
// "sign"         // API输入参数签名结果

func makeRequest(httpMethod string, params map[string]string) (fasthttp.ResponseHeader, []byte) {
	params = setSystemParam(params)
	checkRequiredParams(params)
	checkApiGateway(params)

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseResponse(res)
		fasthttp.ReleaseRequest(req)
	}()
	req.Header.SetUserAgent("jop-go")
	req.Header.SetMethod(strings.ToUpper(httpMethod))

	req.SetRequestURI(RootEndpoint)

	params["sign"] = getSign(params)

	queryStr := getQueryString(params)
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

func checkApiGateway(params map[string]string) {
	if strings.Contains(RootEndpoint, "router.jd.com") {
		if _, ok := params["param_json"]; !ok {
			panic("U use <union>, business params key must be <param_json>")
		}
	} else if strings.Contains(RootEndpoint, "router.jd.com") {
		if _, ok := params["360buy_param_json"]; !ok {
			panic("U use <jos>, business params key must be <360buy_param_json>")
		}
	} else {
		panic("API Gateway may be wrong")
	}
}

// 设置系统参数
func setSystemParam(params map[string]string) map[string]string {
	params["app_key"] = AppKey
	params["timestamp"] = time.Now().Local().Format("2006-01-02 15:04:05")
	if "" != AccessToken {
		params["access_token"] = AccessToken
	}
	params["format"] = "json"
	params["sign_method"] = "md5"

	return params
}

// 检查必选参数
func checkRequiredParams(params map[string]string) {
	if "" == AppKey || "" == SecretKey {
		panic("AppKey and SecretKey must be set")
	}
	required := []string{"method", "v"}
	for i := range required {
		if _, ok := params[required[i]]; !ok {
			panic(required[i] + " is required")
		}
	}
}

//  把所有请求参数按照参数名称的 ASCII 码表顺序进行排序并拼接
func getConcatParams(params map[string]string) string {
	s := make([]string, len(params))
	for k := range params {
		s = append(s, fmt.Sprintf("%s%s", k, params[k]))
	}
	sort.Strings(s)
	var concat string
	for i := range s {
		concat += s[i]
	}
	return concat
}

// 组装 HTTP 请求，将所有参数名和参数值采用 utf-8 进行 URL 编码（
func getQueryString(params map[string]string) string {
	s := make([]string, 0)
	for k := range params {
		s = append(s, fmt.Sprintf("%s=%s", k, params[k]))
	}
	joins := strings.Join(s, "&")
	return url.PathEscape(joins)
}

// 把 appSecret 的值拼接在字符串的两端，使用 MD5 进行加密，并转化成大写
func getSign(params map[string]string) string {
	concatParams := getConcatParams(params)
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(SecretKey+concatParams+SecretKey))))
}
